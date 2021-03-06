package helpers

import (
	"context"
	"fmt"
	"sort"

	"github.com/jackc/pgx/v4"
)

type Funcs map[int64]func(context.Context, pgx.Tx) error

func Commands(qs ...string) func(context.Context, pgx.Tx) error {
	return func(ctx context.Context, tx pgx.Tx) error {
		return Run(ctx, tx, qs)
	}
}

// Helper for running multiple commands in sequence.
func Run(ctx context.Context, tx pgx.Tx, qs []string) error {
	for _, q := range qs {
		_, err := tx.Exec(ctx, q)
		if err != nil {
			return err
		}
	}
	return nil
}

// Migrate logic
func Migrate(ctx context.Context, conn *pgx.Conn, funcs Funcs) error {
	var ks []int64
	for k := range funcs {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {
		return ks[i] < ks[j]
	})
	_, err := conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS migrations (mts BIGINT PRIMARY KEY)`)
	if err != nil {
		return err
	}
	for _, k := range ks {
		var exists int
		err := conn.QueryRow(ctx, `SELECT 1 FROM migrations WHERE mts = $1`, k).Scan(&exists)
		if err != nil && err != pgx.ErrNoRows {
			return err
		}
		if exists == 1 {
			continue
		}
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}
		err = funcs[k](ctx, tx)
		if err != nil {
			tx.Rollback(ctx)
			return &Error{Err: err, Key: k}
		}
		_, err = tx.Exec(ctx, `INSERT INTO migrations (mts) VALUES ($1)`, k)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
		err = tx.Commit(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

type Error struct {
	Err error
	Key int64
}

func (e *Error) Error() string {
	return fmt.Sprintf("migrate %d: %v", e.Key, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func Log(printer Printer) *Logger {
	return &Logger{Printer: printer}
}

type Logger struct {
	Printer Printer
}

func (l *Logger) Commands(qs ...string) func(context.Context, pgx.Tx) error {
	return func(ctx context.Context, tx pgx.Tx) error {
		l.print("BEGIN")
		err := l.Run(ctx, tx, qs)
		if err != nil {
			l.print("ROLLBACK")
			return err
		}
		l.print("COMMIT")
		return nil
	}
}

func (l *Logger) Run(ctx context.Context, tx pgx.Tx, qs []string) error {
	for _, q := range qs {
		l.print(q)
		_, err := tx.Exec(ctx, q)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Logger) print(v string) {
	if l.Printer != nil {
		l.Printer.Printf("%s\n", v)
	}
}

type Printer interface {
	Printf(string, ...interface{})
}
