package repository

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
)

type CreateEavAttributeValueParams struct {
	UserID           uint64
	EavAttributeCode string
	Value            string
}

type CreateEavAttributeValueRow struct {
	UserID           uint64
	EavAttributeCode string
}

type GetEavAttributeValueByUserIDAndCodeParams struct {
	UserID           uint64
	EavAttributeCode string
}

type EavAttributeValue interface {
	CreateEavAttributeValue(ctx context.Context, arg *CreateEavAttributeValueParams) (*CreateEavAttributeValueRow, error)
	GetEavAttributeValueByUserIDAndCode(ctx context.Context, arg GetEavAttributeValueByUserIDAndCodeParams) (*entity.EavAttributeValue, error)
	ListEavAttributeValueByUserID(ctx context.Context, userID uint64) ([]*entity.EavAttributeValue, error)
}
