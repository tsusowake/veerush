package db

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
	"github.com/tsusowake/veerush/internal/database/generated"
	"github.com/tsusowake/veerush/internal/database/repository"
)

type EavAttributeValue struct {
	query *generated.Queries
}

func NewEavAttributeValue(q *generated.Queries) *EavAttributeValue {
	return &EavAttributeValue{
		query: q,
	}
}

func (e *EavAttributeValue) CreateEavAttributeValue(
	ctx context.Context,
	arg *repository.CreateEavAttributeValueParams,
) (*repository.CreateEavAttributeValueRow, error) {
	ret, err := e.query.CreateEavAttributeValue(ctx, generated.CreateEavAttributeValueParams{
		UserID:           arg.UserID,
		EavAttributeCode: arg.EavAttributeCode,
		Value:            arg.Value,
	})
	if err != nil {
		return nil, err
	}
	return &repository.CreateEavAttributeValueRow{
		UserID:           ret.UserID,
		EavAttributeCode: ret.EavAttributeCode,
	}, nil
}

func (e *EavAttributeValue) GetEavAttributeValueByUserIDAndCode(
	ctx context.Context,
	arg repository.GetEavAttributeValueByUserIDAndCodeParams,
) (*entity.EavAttributeValue, error) {
	m, err := e.query.GetEavAttributeValueByUserIDAndCode(ctx, generated.GetEavAttributeValueByUserIDAndCodeParams{
		UserID:           arg.UserID,
		EavAttributeCode: arg.EavAttributeCode,
	})
	if err != nil {
		return nil, err
	}
	var vv entity.EavAttributeValue
	return vv.FromModel(&m), nil
}

func (e *EavAttributeValue) ListEavAttributeValueByUserID(
	ctx context.Context,
	userID uint64,
) ([]*entity.EavAttributeValue, error) {
	m, err := e.query.ListEavAttributeValueByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var vv entity.EavAttributeValue
	return vv.FromModels(m), nil
}
