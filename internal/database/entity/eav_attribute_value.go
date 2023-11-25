package entity

type EavAttributeValue struct {
	UserID           uint64 `db:"user_id"`
	EavAttributeCode string `db:"eav_attribute_code"`
	Value            string `db:"value"`
}

// NEED attribute
func (e EavAttributeValue) FormatValue() string {
	return "eav_attribute_values"
}
