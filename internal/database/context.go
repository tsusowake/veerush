package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ctxKey struct{}

func SetDBPool(
	ctx context.Context,
	dbpool *pgxpool.Pool,
) context.Context {
	return context.WithValue(ctx, ctxKey{}, dbpool)
}

func GetDBPool(
	ctx context.Context,
) *pgxpool.Pool {
	return ctx.Value(ctxKey{}).(*pgxpool.Pool)
}
