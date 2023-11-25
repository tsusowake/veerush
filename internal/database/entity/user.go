package entity

import (
	"github.com/tsusowake/veerush/internal/database/generated"
)

type User struct {
	ID uint64 `db:"id"`
}

func (u *User) FromModel(m *generated.User) *User {
	return &User{
		ID: m.ID,
	}
}
