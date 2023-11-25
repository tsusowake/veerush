package entity

type EavAttributeOption struct {
	ID               uint64 `db:"id"`
	EavAttributeCode string `db:"eav_attribute_code"`
	Value            string `db:"value"`
	Ordering         uint16 `db:"ordering"`
	IsVisible        bool   `db:"is_visible"`
}
