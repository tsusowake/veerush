package db

import "github.com/tsusowake/veerush/internal/database/generated"

type EavAttributeValue struct {
	query *generated.Queries
}

func NewEavAttributeValue(q *generated.Queries) *EavAttributeValue {
	return &EavAttributeValue{
		query: q,
	}
}
