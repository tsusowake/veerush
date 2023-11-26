package entity

import "github.com/tsusowake/veerush/internal/database/generated"

type EavAttributeOption struct {
	ID               uint64
	EavAttributeCode string
	Value            string
	Ordering         uint16
	IsVisible        bool
}

func (e *EavAttributeOption) FromModel(m *generated.EavAttributeOption) *EavAttributeOption {
	return &EavAttributeOption{
		ID:               m.ID,
		EavAttributeCode: m.EavAttributeCode,
		Value:            m.Value,
		Ordering:         m.Ordering,
		IsVisible:        m.IsVisible,
	}
}

func (e *EavAttributeOption) FromModels(m []generated.EavAttributeOption) []*EavAttributeOption {
	ret := make([]*EavAttributeOption, len(m))
	for i, mm := range m {
		ret[i] = e.FromModel(&mm)
	}
	return ret
}
