-- name: GetEavAttributeByCode :one
select *
from eav_attributes
where code = $1
limit 1;

-- name: ListEavAttributes :many
select *
from eav_attributes;