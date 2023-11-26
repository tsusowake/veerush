package repository

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
)

type CreateEavAttributeOptionParams struct {
	EavAttributeCode string
	Value            string
	Ordering         uint16
	IsVisible        bool
}

type EavAttributeOption interface {
	CreateEavAttributeOption(ctx context.Context, arg *CreateEavAttributeOptionParams) (string, error)
	ListEavAttributeOptionByEavAttributeCode(ctx context.Context, eavAttributeCode string) ([]*entity.EavAttributeOption, error)
}
