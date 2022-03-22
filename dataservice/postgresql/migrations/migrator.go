package migrations

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/n0byk/short_url_backend/helpers"
)

func MigrateFuncs() helpers.Funcs {
	m := helpers.Funcs{}

	m[1647785956] = func(ctx context.Context, tx pgx.Tx) error {
		return helpers.Run(ctx, tx, []string{
			`CREATE TABLE url_catalog (
        id SERIAL PRIMARY KEY,
		user_id  TEXT NOT NULL,
        full_url TEXT NOT NULL,
        short_url TEXT NOT NULL UNIQUE,
		UNIQUE(full_url)
      )`,
		})
	}
	return m
}
