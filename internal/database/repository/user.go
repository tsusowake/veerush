package repository

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
)

type User interface {
	GetUserByID(ctx context.Context, id uint64) (*entity.User, error)
	CreateUser(ctx context.Context) (*uint64, error)
	DeleteUserByID(ctx context.Context, id uint64) error
}
