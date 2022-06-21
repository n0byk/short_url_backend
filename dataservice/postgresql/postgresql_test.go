package postgresql

import (
	"context"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

func Test_dbRepository_AddURL(t *testing.T) {
	type fields struct {
		db *pgx.Conn
	}
	type args struct {
		ctx  context.Context
		url  string
		user string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &dbRepository{
				db: tt.fields.db,
			}
			got, got1, err := db.AddURL(tt.args.ctx, tt.args.url, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("dbRepository.AddURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dbRepository.AddURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("dbRepository.AddURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dbRepository_SetUserData(t *testing.T) {
	type fields struct {
		db *pgx.Conn
	}
	type args struct {
		ctx  context.Context
		key  string
		url  string
		user string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &dbRepository{
				db: tt.fields.db,
			}
			if err := db.SetUserData(tt.args.ctx, tt.args.key, tt.args.url, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("dbRepository.SetUserData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dbRepository_DBPing(t *testing.T) {
	type fields struct {
		db *pgx.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &dbRepository{
				db: tt.fields.db,
			}
			if err := db.DBPing(); (err != nil) != tt.wantErr {
				t.Errorf("dbRepository.DBPing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dbRepository_GetUserData(t *testing.T) {
	type fields struct {
		db *pgx.Conn
	}
	type args struct {
		ctx  context.Context
		user string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.URLCatalog
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &dbRepository{
				db: tt.fields.db,
			}
			got, err := db.GetUserData(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("dbRepository.GetUserData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbRepository.GetUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbRepository_GetURL(t *testing.T) {
	type fields struct {
		db *pgx.Conn
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &dbRepository{
				db: tt.fields.db,
			}
			got, err := db.GetURL(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("dbRepository.GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dbRepository.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbRepository_BulkDelete(t *testing.T) {
	type fields struct {
		db *pgx.Conn
	}
	type args struct {
		ctx    context.Context
		urls   []string
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &dbRepository{
				db: tt.fields.db,
			}
			if err := db.BulkDelete(tt.args.ctx, tt.args.urls, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("dbRepository.BulkDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDBRepository(t *testing.T) {
	type args struct {
		db *pgx.Conn
	}
	tests := []struct {
		name string
		args args
		want dataservice.Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDBRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDBRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
