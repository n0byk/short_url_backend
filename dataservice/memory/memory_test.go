package memory

import (
	"context"
	"reflect"
	"testing"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

func Test_memoryRepository_AddURL(t *testing.T) {
	type fields struct {
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
			m := &memoryRepository{
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			got, got1, err := m.AddURL(tt.args.ctx, tt.args.url, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryRepository.AddURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("memoryRepository.AddURL() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("memoryRepository.AddURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_memoryRepository_DBPing(t *testing.T) {
	type fields struct {
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
			m := &memoryRepository{
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			if err := m.DBPing(); (err != nil) != tt.wantErr {
				t.Errorf("memoryRepository.DBPing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memoryRepository_SetUserData(t *testing.T) {
	type fields struct {
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
			m := &memoryRepository{
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			if err := m.SetUserData(tt.args.ctx, tt.args.key, tt.args.url, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("memoryRepository.SetUserData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memoryRepository_GetUserData(t *testing.T) {
	type fields struct {
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
			m := &memoryRepository{
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			got, err := m.GetUserData(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryRepository.GetUserData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryRepository.GetUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryRepository_GetURL(t *testing.T) {
	type fields struct {
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
			m := &memoryRepository{
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			got, err := m.GetURL(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryRepository.GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("memoryRepository.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryRepository_BulkDelete(t *testing.T) {
	type fields struct {
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
			m := &memoryRepository{
				urlsDB:   tt.fields.urlsDB,
				userData: tt.fields.userData,
			}
			if err := m.BulkDelete(tt.args.ctx, tt.args.urls, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("memoryRepository.BulkDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMemoryRepository(t *testing.T) {
	tests := []struct {
		name string
		want dataservice.Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
