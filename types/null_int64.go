package types

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 struct {
	sql.NullInt64
}

func NewNullInt64(data int64, zeroNull bool) NullInt64 {
	// 0 是否设置为Null
	if zeroNull {
		if data == 0 {
			return NullInt64{sql.NullInt64{data, false}}
		}
	}

	return NullInt64{sql.NullInt64{data, true}}
}

func (v *NullInt64) IsZero() bool {
	if !v.Valid || v.Int64 <= 0 {
		return true
	}

	return false
}

func (v *NullInt64) SetZeroNoNull() {
	v.Valid = true
}

func (v *NullInt64) SetZeroNull() {
	v.Valid = false
}

func (v *NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	var s *int64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Int64 = *s
	} else {
		v.Valid = false
	}
	return nil
}
