package entity

type EavAttribute struct {
	Code        string `db:"code"`
	Name        string `db:"name"`
	ValueType   string `db:"value_type"`
	Description string `db:"description"`
	FieldFormat string `db:"field_format"`
	Regexp      string `db:"regexp"`
	MinLength   uint16 `db:"min_length"`
	MaxLength   uint16 `db:"max_length"`
	IsSelection bool   `db:"is_selection"`
	IsRequired  bool   `db:"is_required"`
	IsVisible   bool   `db:"is_visible"`
}
