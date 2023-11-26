package db

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
	"github.com/tsusowake/veerush/internal/database/generated"
	"github.com/tsusowake/veerush/internal/database/repository"
)

type EavAttributeOption struct {
	query *generated.Queries
}

func NewEavAttributeOption(q *generated.Queries) *EavAttributeOption {
	return &EavAttributeOption{
		query: q,
	}
}

func (e *EavAttributeOption) CreateEavAttributeOption(
	ctx context.Context,
	arg *repository.CreateEavAttributeOptionParams,
) (string, error) {
	return e.query.CreateEavAttributeOption(ctx, generated.CreateEavAttributeOptionParams{
		EavAttributeCode: arg.EavAttributeCode,
		Value:            arg.Value,
		Ordering:         arg.Ordering,
		IsVisible:        arg.IsVisible,
	})
}

func (e *EavAttributeOption) ListEavAttributeOptionByEavAttributeCode(
	ctx context.Context,
	eavAttributeCode string,
) ([]*entity.EavAttributeOption, error) {
	m, err := e.query.ListEavAttributeOptionByEavAttributeCode(ctx, eavAttributeCode)
	if err != nil {
		return nil, err
	}
	var vv *entity.EavAttributeOption
	return vv.FromModels(m), nil
}
