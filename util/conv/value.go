package conv

import "github.com/jackc/pgx/v5/pgtype"

func ToNullUint16(v pgtype.Int2) *uint16 {
	if v.Valid {
		vv := uint16(v.Int16)
		return &vv
	}
	return nil
}

func ToPgTypeInt2(v *uint16) pgtype.Int2 {
	if v != nil {
		return pgtype.Int2{Int16: int16(*v), Valid: true}
	}
	return pgtype.Int2{Int16: 0, Valid: false}
}

func ToNullString(v pgtype.Text) *string {
	if v.Valid {
		vv := v.String
		return &vv
	}
	return nil
}

func ToPgTypeText(v *string) pgtype.Text {
	if v != nil {
		return pgtype.Text{String: *v, Valid: true}
	}
	if *v != "" {
		return pgtype.Text{String: *v, Valid: true}
	}
	return pgtype.Text{String: "", Valid: false}
}
