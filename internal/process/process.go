package process

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tsusowake/veerush/internal/database/generated"
	"github.com/tsusowake/veerush/internal/database/repository"
	"github.com/tsusowake/veerush/internal/database/repository/db"
)

type Process struct {
	repository.EavAttribute
	repository.EavAttributeValue
	repository.EavAttributeOption
	repository.User
}

func NewProcess(
	dbpool *pgxpool.Pool,
) *Process {
	query := generated.New(dbpool)
	return &Process{
		EavAttribute:       db.NewEavAttribute(query),
		EavAttributeValue:  db.NewEavAttributeValue(query),
		EavAttributeOption: db.NewEavAttributeOption(query),
		User:               db.NewUser(query),
	}
}
