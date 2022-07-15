package dataservice

import (
	"context"

	"github.com/n0byk/short_url_backend/dataservice/entities"
)

// Data Base repository methods
type Repository interface {
	AddURL(ctx context.Context, url, user string) (string, bool, error)
	GetURL(ctx context.Context, key string) (string, error)
	SetUserData(ctx context.Context, key, url, user string) error
	GetUserData(ctx context.Context, user string) ([]entities.URLCatalog, error)
	DBPing() error
	BulkDelete(ctx context.Context, urls []string, userID string) error
	StatInfo(ctx context.Context) (entities.StatInfo, error)
}
