package filestorage

import (
	"context"
	"os"
	"reflect"
	"testing"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

func Test_fileRepository_AddURL(t *testing.T) {
	type fields struct {
		f        *os.File
		urlsDB   map[string]entities.URLdb
		userData map[string][]entities.URLCatalog
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
			f := &fileRepository{
				f:        tt.fields.f,
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			got, got1, err := f.AddURL(tt.args.ctx, tt.args.url, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.AddURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fileRepository.AddURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fileRepository.AddURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fileRepository_DBPing(t *testing.T) {
	type fields struct {
		f        *os.File
		urlsDB   map[string]entities.URLdb
		userData map[string][]entities.URLCatalog
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
			f := &fileRepository{
				f:        tt.fields.f,
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			if err := f.DBPing(); (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.DBPing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileRepository_SetUserData(t *testing.T) {
	type fields struct {
		f        *os.File
		urlsDB   map[string]entities.URLdb
		userData map[string][]entities.URLCatalog
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
			f := &fileRepository{
				f:        tt.fields.f,
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			if err := f.SetUserData(tt.args.ctx, tt.args.key, tt.args.url, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.SetUserData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileRepository_GetUserData(t *testing.T) {
	type fields struct {
		f        *os.File
		urlsDB   map[string]entities.URLdb
		userData map[string][]entities.URLCatalog
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
			f := &fileRepository{
				f:        tt.fields.f,
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			got, err := f.GetUserData(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.GetUserData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileRepository.GetUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileRepository_GetURL(t *testing.T) {
	type fields struct {
		f        *os.File
		urlsDB   map[string]entities.URLdb
		userData map[string][]entities.URLCatalog
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
			f := &fileRepository{
				f:        tt.fields.f,
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			got, err := f.GetURL(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fileRepository.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileRepository_BulkDelete(t *testing.T) {
	type fields struct {
		f        *os.File
		urlsDB   map[string]entities.URLdb
		userData map[string][]entities.URLCatalog
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
			f := &fileRepository{
				f:        tt.fields.f,
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			if err := f.BulkDelete(tt.args.ctx, tt.args.urls, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("fileRepository.BulkDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewFileRepository(t *testing.T) {
	type args struct {
		f *os.File
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
			if got := NewFileRepository(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
