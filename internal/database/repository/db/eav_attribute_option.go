package db

import "github.com/tsusowake/veerush/internal/database/generated"

type EavAttributeOption struct {
	query *generated.Queries
}

func NewEavAttributeOption(q *generated.Queries) *EavAttributeOption {
	return &EavAttributeOption{
		query: q,
	}
}
