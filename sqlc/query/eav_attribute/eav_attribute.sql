-- name: GetEavAttributeByCode :one
select *
from eav_attributes
where code = $1
limit 1;

-- name: ListEavAttributes :many
select *
from eav_attributes;

-- name: CreateEavAttribute :one
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
returning code;
