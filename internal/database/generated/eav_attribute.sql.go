// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: eav_attribute.sql

package generated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEavAttribute = `-- name: CreateEavAttribute :one
insert into eav_attributes (code,
                            name,
                            value_type,
                            description,
                            field_format,
                            regexp,
                            min_length,
                            max_length,
                            is_selection,
                            is_required)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
returning code
`

type CreateEavAttributeParams struct {
	Code        string
	Name        string
	ValueType   uint8
	Description string
	FieldFormat pgtype.Text
	Regexp      pgtype.Text
	MinLength   pgtype.Int2
	MaxLength   pgtype.Int2
	IsSelection bool
	IsRequired  bool
}

func (q *Queries) CreateEavAttribute(ctx context.Context, arg CreateEavAttributeParams) (string, error) {
	row := q.db.QueryRow(ctx, createEavAttribute,
		arg.Code,
		arg.Name,
		arg.ValueType,
		arg.Description,
		arg.FieldFormat,
		arg.Regexp,
		arg.MinLength,
		arg.MaxLength,
		arg.IsSelection,
		arg.IsRequired,
	)
	var code string
	err := row.Scan(&code)
	return code, err
}

const getEavAttributeByCode = `-- name: GetEavAttributeByCode :one
select code, name, value_type, description, field_format, regexp, min_length, max_length, is_selection, is_required, created_at
from eav_attributes
where code = $1
limit 1
`

func (q *Queries) GetEavAttributeByCode(ctx context.Context, code string) (EavAttribute, error) {
	row := q.db.QueryRow(ctx, getEavAttributeByCode, code)
	var i EavAttribute
	err := row.Scan(
		&i.Code,
		&i.Name,
		&i.ValueType,
		&i.Description,
		&i.FieldFormat,
		&i.Regexp,
		&i.MinLength,
		&i.MaxLength,
		&i.IsSelection,
		&i.IsRequired,
		&i.CreatedAt,
	)
	return i, err
}

const listEavAttributes = `-- name: ListEavAttributes :many
select code, name, value_type, description, field_format, regexp, min_length, max_length, is_selection, is_required, created_at
from eav_attributes
`

func (q *Queries) ListEavAttributes(ctx context.Context) ([]EavAttribute, error) {
	rows, err := q.db.Query(ctx, listEavAttributes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EavAttribute
	for rows.Next() {
		var i EavAttribute
		if err := rows.Scan(
			&i.Code,
			&i.Name,
			&i.ValueType,
			&i.Description,
			&i.FieldFormat,
			&i.Regexp,
			&i.MinLength,
			&i.MaxLength,
			&i.IsSelection,
			&i.IsRequired,
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
