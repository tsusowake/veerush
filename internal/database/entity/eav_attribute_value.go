package entity

import "github.com/tsusowake/veerush/internal/database/generated"

type EavAttributeValue struct {
	UserID           uint64
	EavAttributeCode string
	Value            string
}

func (e *EavAttributeValue) FromModel(m *generated.EavAttributeValue) *EavAttributeValue {
	return &EavAttributeValue{
		UserID:           m.UserID,
		EavAttributeCode: m.EavAttributeCode,
		Value:            m.Value,
	}
}

func (e *EavAttributeValue) FromModels(m []generated.EavAttributeValue) []*EavAttributeValue {
	ret := make([]*EavAttributeValue, len(m))
	for i, mm := range m {
		ret[i] = e.FromModel(&mm)
	}
	return ret
}
