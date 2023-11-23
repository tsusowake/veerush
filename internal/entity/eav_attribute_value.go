package entity

type EavAttributeValue struct {
	UserID           uint64 `db:"user_id"`
	EavAttributeCode string `db:"eav_attribute_code"`
	Value            string `db:"value"`
}
