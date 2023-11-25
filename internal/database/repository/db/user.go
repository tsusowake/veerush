package db

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
	"github.com/tsusowake/veerush/internal/database/generated"
)

type User struct {
	query *generated.Queries
}

func NewUser(q *generated.Queries) *User {
	return &User{
		query: q,
	}
}

func (u *User) GetUserByID(
	ctx context.Context,
	id uint64,
) (*entity.User, error) {
	v, err := u.query.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var uu *entity.User
	return uu.FromModel(&v), nil
}

func (u *User) CreateUser(ctx context.Context) (*uint64, error) {
	id, err := u.query.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (u *User) DeleteUserByID(ctx context.Context, id uint64) error {
	return u.query.DeleteUserByID(ctx, id)
}
