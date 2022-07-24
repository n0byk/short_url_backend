package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "net/http/pprof"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4"

	grpcmethod "github.com/n0byk/short_url_backend/adapters/grpcMethod"
	httpMethod "github.com/n0byk/short_url_backend/adapters/httpMethod"
	httpMethodhelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	filestorage "github.com/n0byk/short_url_backend/dataservice/filestorage"
	memory "github.com/n0byk/short_url_backend/dataservice/memory"
	postgresql "github.com/n0byk/short_url_backend/dataservice/postgresql"
	migrations "github.com/n0byk/short_url_backend/dataservice/postgresql/migrations"
	helpers "github.com/n0byk/short_url_backend/helpers"

	pb "github.com/n0byk/short_url_backend/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var appEnv config.AppConfig

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

type server struct{}

func init() {
	flag.StringVar(&appEnv.ServerAddress, "a", "localhost:8080", "SERVER_ADDRESS")
	flag.StringVar(&appEnv.BaseURL, "b", "http://localhost:8080", "BASE_URL")
	flag.StringVar(&appEnv.FileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.StringVar(&appEnv.DB, "d", "", "DATABASE_DSN")
	flag.StringVar(&appEnv.CONFIG, "c", "", "CONFIG")
	flag.StringVar(&appEnv.CONFIG, "config", "", "CONFIG")
	flag.BoolVar(&appEnv.TLS, "s", false, "ENABLE_HTTPS")
	flag.StringVar(&appEnv.TLScrt, "crt", "./certs/https-server.crt", "TLScrt")
	flag.StringVar(&appEnv.TLSkey, "key", "./certs/https-server.key", "TLSkey")
	flag.StringVar(&appEnv.TrustedSubnet, "t", "", "TRUSTED_SUBNET")

	if err := env.Parse(&appEnv); err != nil {
		log.Fatalf("Unset vars: %v", err)
	}
	//go run -ldflags "-X main.buildVersion=v1.0.1 -X main.buildCommit=Reolve_19_inc -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')'" main.go
	config.ShowBuildInfo(buildVersion, buildDate, buildCommit)
}

// The main function
func main() {
	flag.Parse()

	if len(appEnv.CONFIG) > 1 {
		config.ReadJSONConfig(appEnv)
	}
	var storage dataservice.Repository

	if appEnv.FileStoragePath != "" {
		log.Println("File storage init.")
		f, err := os.OpenFile(appEnv.FileStoragePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		storage = filestorage.NewFileRepository(f)
	}

	if appEnv.TrustedSubnet != "" {
		_, appEnv.TrustedCIDR, _ = net.ParseCIDR(appEnv.TrustedSubnet)

	}

	if appEnv.DB != "" {
		log.Println("Postgres storage init.")
		conn, err := pgx.Connect(context.Background(), appEnv.DB)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}

		err = helpers.Migrate(context.Background(), conn, migrations.MigrateFuncs())
		if err != nil {
			log.Fatal(err)
		}
		storage = postgresql.NewDBRepository(conn)
		defer conn.Close(context.Background())
	}

	if appEnv.DB == "" && appEnv.FileStoragePath == "" {
		log.Println("Memory storage init.")
		storage = memory.NewMemoryRepository()

	}

	config.AppService = config.Service{ShortLinkLen: 7, BaseURL: appEnv.BaseURL, Storage: storage, Env: appEnv}

	srv := &http.Server{
		Addr:    appEnv.ServerAddress,
		Handler: httpMethodhelpers.Gzip(httpMethodhelpers.Cookie(httpMethod.Routers())),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {

		if appEnv.TLS {
			if err := srv.ListenAndServeTLS(appEnv.TLScrt, appEnv.TLSkey); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}

		}

		listener, err := net.Listen("tcp", appEnv.GRPCServerAddress)

		if err != nil {
			grpclog.Fatalf("failed to listen: %v", err)
		}

		opts := []grpc.ServerOption{}
		grpcServer := grpc.NewServer(opts...)

		pb.RegisterServiceLogicServer(grpcServer, &grpcmethod.GRPCLogic{})
		grpcServer.Serve(listener)

	}()

	log.Print("Started at " + appEnv.ServerAddress)
	log.Print("GRPC Started at " + appEnv.GRPCServerAddress)

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")

}
