package repository

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
)

type CreateEavAttributeParams struct {
	Code        string
	Name        string
	ValueType   int16
	Description string
	FieldFormat *string
	Regexp      *string
	MinLength   *uint16
	MaxLength   *uint16
	IsSelection bool
	IsRequired  bool
}

type EavAttribute interface {
	GetEavAttributeByCode(ctx context.Context, code string) (*entity.EavAttribute, error)
	ListEavAttributes(ctx context.Context) ([]*entity.EavAttribute, error)
	CreateEavAttribute(ctx context.Context, arg CreateEavAttributeParams) (string, error)
}
