package entity

import (
	"github.com/tsusowake/veerush/internal/database/generated"
	"github.com/tsusowake/veerush/util/conv"
)

type EavAttribute struct {
	Code        string
	Name        string
	ValueType   ValueType
	Description string
	FieldFormat *string
	Regexp      *string
	MinLength   *uint16
	MaxLength   *uint16
	IsSelection bool
	IsRequired  bool
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
	CodeValueTypeString          = "code.valueType.string"
	CodeValueTypeSelectionString = "code.valueType.selection.string"
	CodeValueTypeInt             = "code.valueType.int"
	CodeValueTypeSelectionInt    = "code.valueType.selection.int"
	CodeValueTypeFloat           = "code.valueType.float"
	CodeValueTypeBool            = "code.valueType.bool"
	CodeValueTypeDate            = "code.valueType.date"
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

func (v ValueType) Value() uint8 {
	return uint8(v)
}
