package entity

import (
	"github.com/tsusowake/veerush/internal/database/generated"
	"github.com/tsusowake/veerush/util/conv"
)

type EavAttribute struct {
	Code        string    `db:"code"`
	Name        string    `db:"name"`
	ValueType   ValueType `db:"value_type"`
	Description string    `db:"description"`
	FieldFormat *string   `db:"field_format"`
	Regexp      *string   `db:"regexp"`
	MinLength   *uint16   `db:"min_length"`
	MaxLength   *uint16   `db:"max_length"`
	IsSelection bool      `db:"is_selection"`
	IsRequired  bool      `db:"is_required"`
}

func (e *EavAttribute) FromModel(m *generated.EavAttribute) *EavAttribute {
	return &EavAttribute{
		Code:        m.Code,
		Name:        m.Name,
		ValueType:   ValueType(m.ValueType),
		Description: m.Description,
		FieldFormat: conv.ToNullString(m.FieldFormat),
		Regexp:      conv.ToNullString(m.Regexp),
		MinLength:   conv.ToNullUint16(m.MinLength),
		MaxLength:   conv.ToNullUint16(m.MaxLength),
		IsSelection: m.IsSelection,
		IsRequired:  m.IsRequired,
	}
}

func (e *EavAttribute) FromModels(m []generated.EavAttribute) []*EavAttribute {
	ret := make([]*EavAttribute, len(m))
	for i, mm := range m {
		ret[i] = e.FromModel(&mm)
	}
	return ret
}

type Code string

const (
	CodeJobType     = "jobType"
	CodePosition    = "position"
	CodeCompanyName = "companyName"
)

type ValueType uint8

const (
	ValueTypeNone   ValueType = 0
	ValueTypeString ValueType = 1
	ValueTypeInt    ValueType = 2
	ValueTypeFloat  ValueType = 3
	ValueTypeBool   ValueType = 4
	ValueTypeDate   ValueType = 5
)
