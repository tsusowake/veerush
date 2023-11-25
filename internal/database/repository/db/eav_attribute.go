package db

import (
	"context"

	"github.com/tsusowake/veerush/internal/database/entity"
	"github.com/tsusowake/veerush/internal/database/generated"
	"github.com/tsusowake/veerush/internal/database/repository"
	"github.com/tsusowake/veerush/util/conv"
)

type EavAttribute struct {
	query *generated.Queries
}

func NewEavAttribute(q *generated.Queries) *EavAttribute {
	return &EavAttribute{
		query: q,
	}
}

func (e *EavAttribute) GetEavAttributeByCode(
	ctx context.Context,
	code string,
) (*entity.EavAttribute, error) {
	v, err := e.query.GetEavAttributeByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	var ee entity.EavAttribute
	return ee.FromModel(&v), nil
}

func (e *EavAttribute) ListEavAttributes(
	ctx context.Context,
) ([]*entity.EavAttribute, error) {
	v, err := e.query.ListEavAttributes(ctx)
	if err != nil {
		return nil, err
	}
	var ee entity.EavAttribute
	return ee.FromModels(v), nil
}

func (e *EavAttribute) CreateEavAttribute(
	ctx context.Context,
	arg repository.CreateEavAttributeParams,
) (string, error) {
	v, err := e.query.CreateEavAttribute(ctx, generated.CreateEavAttributeParams{
		Code:        arg.Code,
		Name:        arg.Name,
		ValueType:   arg.ValueType,
		Description: arg.Description,
		FieldFormat: conv.ToPgTypeText(arg.FieldFormat),
		Regexp:      conv.ToPgTypeText(arg.Regexp),
		MinLength:   conv.ToPgTypeInt2(arg.MinLength),
		MaxLength:   conv.ToPgTypeInt2(arg.MaxLength),
		IsSelection: arg.IsSelection,
		IsRequired:  arg.IsRequired,
	})
	if err != nil {
		return "", err
	}
	return v, nil
}
