-- name: GetEavAttributeValueByUserIDAndCode :one
select *
from eav_attribute_values
where user_id = $1
  and eav_attribute_code = $2
limit 1;

-- name: ListEavAttributeValueByUserID :many
select *
from eav_attribute_values
where user_id = $1;

-- name: CreateEavAttributeValue :one
insert
into eav_attribute_values (user_id,
                           eav_attribute_code,
                           value)
values ($1,
        $2,
        $3)
returning user_id, eav_attribute_code;