package process

import (
	"context"
	"time"

	"github.com/tsusowake/veerush/internal/database/entity"
	"github.com/tsusowake/veerush/internal/database/repository"
)

var eavAttributes = []repository.CreateEavAttributeParams{
	{
		Code:        entity.CodeValueTypeString,
		Name:        "value is string",
		ValueType:   entity.ValueTypeString,
		Description: "string value description",
		FieldFormat: "",
		Regexp:      "",
		MinLength:   nil,
		MaxLength:   nil,
		IsSelection: false,
		IsRequired:  true,
	},
	{
		Code:        entity.CodeValueTypeSelectionString,
		Name:        "value is selection string",
		ValueType:   entity.ValueTypeString,
		Description: "selection string value description",
		FieldFormat: "",
		Regexp:      "",
		MinLength:   nil,
		MaxLength:   nil,
		IsSelection: true,
		IsRequired:  true,
	},
	{
		Code:        entity.CodeValueTypeDate,
		Name:        "value is date",
		ValueType:   entity.ValueTypeDate,
		Description: "date value description",
		FieldFormat: time.DateOnly,
		Regexp:      "^(19|20)\\d\\d-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$",
		MinLength:   nil,
		MaxLength:   nil,
		IsSelection: false,
		IsRequired:  true,
	},
}

var eavAttributeValues = []repository.CreateEavAttributeValueParams{
	{
		EavAttributeCode: entity.CodeValueTypeString,
		Value:            "eav attirbute value string",
	},
	{
		EavAttributeCode: entity.CodeValueTypeSelectionString,
		Value:            "string option 2",
	},
	{
		EavAttributeCode: entity.CodeValueTypeDate,
		Value:            "2023-11-25",
	},
}

var eavAttributeOptions = []repository.CreateEavAttributeOptionParams{
	{
		EavAttributeCode: entity.CodeValueTypeSelectionString,
		Value:            "string option 1",
		Ordering:         1,
		IsVisible:        true,
	},
	{
		EavAttributeCode: entity.CodeValueTypeSelectionString,
		Value:            "string option 2",
		Ordering:         2,
		IsVisible:        true,
	},
	{
		EavAttributeCode: entity.CodeValueTypeSelectionString,
		Value:            "string option 3",
		Ordering:         3,
		IsVisible:        true,
	},
}

func (p *Process) DoVeerush(
	ctx context.Context,
) (*uint64, error) {
	userID, err := p.User.CreateUser(ctx)
	if err != nil {
		return nil, err
	}

	// attributes
	for _, attr := range eavAttributes {
		_, err := p.EavAttribute.CreateEavAttribute(ctx, &attr)
		if err != nil {
			return nil, err
		}
	}
	// options
	for _, opt := range eavAttributeOptions {
		_, err := p.EavAttributeOption.CreateEavAttributeOption(ctx, &opt)
		if err != nil {
			return nil, err
		}
	}
	// values
	for _, val := range eavAttributeValues {
		val.UserID = *userID
		_, err := p.EavAttributeValue.CreateEavAttributeValue(ctx, &val)
		if err != nil {
			return nil, err
		}
	}

	// TODO select & format
	return userID, nil
}
