// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: eav_attribute_value.sql

package generated

import (
	"context"
)

const createEavAttributeValue = `-- name: CreateEavAttributeValue :one
insert
into eav_attribute_values (user_id,
                           eav_attribute_code,
                           value)
values ($1,
        $2,
        $3)
returning user_id, eav_attribute_code
`

type CreateEavAttributeValueParams struct {
	UserID           uint64
	EavAttributeCode string
	Value            string
}

type CreateEavAttributeValueRow struct {
	UserID           uint64
	EavAttributeCode string
}

func (q *Queries) CreateEavAttributeValue(ctx context.Context, arg CreateEavAttributeValueParams) (CreateEavAttributeValueRow, error) {
	row := q.db.QueryRow(ctx, createEavAttributeValue, arg.UserID, arg.EavAttributeCode, arg.Value)
	var i CreateEavAttributeValueRow
	err := row.Scan(&i.UserID, &i.EavAttributeCode)
	return i, err
}

const getEavAttributeValueByUserIDAndCode = `-- name: GetEavAttributeValueByUserIDAndCode :one
select user_id, eav_attribute_code, value, created_at
from eav_attribute_values
where user_id = $1
  and eav_attribute_code = $2
limit 1
`

type GetEavAttributeValueByUserIDAndCodeParams struct {
	UserID           uint64
	EavAttributeCode string
}

func (q *Queries) GetEavAttributeValueByUserIDAndCode(ctx context.Context, arg GetEavAttributeValueByUserIDAndCodeParams) (EavAttributeValue, error) {
	row := q.db.QueryRow(ctx, getEavAttributeValueByUserIDAndCode, arg.UserID, arg.EavAttributeCode)
	var i EavAttributeValue
	err := row.Scan(
		&i.UserID,
		&i.EavAttributeCode,
		&i.Value,
		&i.CreatedAt,
	)
	return i, err
}

const listEavAttributeValueByUserID = `-- name: ListEavAttributeValueByUserID :many
select user_id, eav_attribute_code, value, created_at
from eav_attribute_values
where user_id = $1
`

func (q *Queries) ListEavAttributeValueByUserID(ctx context.Context, userID uint64) ([]EavAttributeValue, error) {
	rows, err := q.db.Query(ctx, listEavAttributeValueByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EavAttributeValue
	for rows.Next() {
		var i EavAttributeValue
		if err := rows.Scan(
			&i.UserID,
			&i.EavAttributeCode,
			&i.Value,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
