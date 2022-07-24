package grpcMethod

import (
	"context"
	"log"

	"github.com/n0byk/short_url_backend/config"
	pb "github.com/n0byk/short_url_backend/proto"
)

type GRPCLogic struct {
	s pb.ServiceLogicServer
}

func (s *GRPCLogic) BulkAddURL(ctx context.Context, in *pb.BulkAddURLRequest) (*pb.BulkAddURLResponse, error) {

	var response = pb.BulkAddURLResponse{}
	for _, item := range in.OriginalUrls {

		token, _, err := config.AppService.Storage.AddURL(context.Background(), string(item.ShortUrl), in.UserID)
		if err != nil {
			log.Println("Unable to get token")
			return &pb.BulkAddURLResponse{}, err
		}

		config.AppService.Storage.SetUserData(context.Background(), string(item.ShortUrl), config.AppService.BaseURL+"/"+token, in.UserID)
		response.ShortUrls = append(response.ShortUrls, &pb.BulkAddURLResponse_BulkAResponse{CorrelationId: string(item.CorrelationId), ShortUrl: config.AppService.BaseURL + "/" + token})

	}
	return &response, nil

}

func (s *GRPCLogic) BulkDelete(ctx context.Context, in *pb.BulkDeleteRequest) (*pb.Empty, error) {

	err := config.AppService.Storage.BulkDelete(context.Background(), in.Values, in.UserID)

	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (s *GRPCLogic) Stats(ctx context.Context, in *pb.Empty) (*pb.StatsResponse, error) {

	data, err := config.AppService.Storage.StatInfo(context.Background())

	if err != nil {
		return &pb.StatsResponse{}, err
	}

	return &pb.StatsResponse{UrlsCount: int32(data.URLs), UsersCount: int32(data.Users)}, nil
}

func (s *GRPCLogic) DBPing(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {

	err := config.AppService.Storage.DBPing()
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (s *GRPCLogic) AddURLHandler(ctx context.Context, in *pb.AddURLHandlerRequest) (*pb.AddURLHandlerResponse, error) {

	token, _, err := config.AppService.Storage.AddURL(context.Background(), string(in.Value), in.UserID)
	if err != nil {
		log.Println("Some error while adding new URL")
		return &pb.AddURLHandlerResponse{}, err
	}

	config.AppService.Storage.SetUserData(ctx, string(in.Value), config.AppService.BaseURL+"/"+token, in.Value)
	return &pb.AddURLHandlerResponse{Value: config.AppService.BaseURL + "/" + token}, nil
}

func (s *GRPCLogic) GetURL(ctx context.Context, in *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	fullURL, err := config.AppService.Storage.GetURL(context.Background(), in.Value)
	if err != nil {
		return &pb.GetURLResponse{}, err
	}

	return &pb.GetURLResponse{Value: fullURL}, nil
}
