-- name: ListEavAttributeOptionByEavAttributeCode :many
select *
from eav_attribute_options
where eav_attribute_code = $1
order by ordering
;

-- name: CreateEavAttributeOption :one
insert
into eav_attribute_options (eav_attribute_code,
                            value,
                            ordering,
                            is_visible)
values ($1,
        $2,
        $3,
        $4)
returning eav_attribute_code;