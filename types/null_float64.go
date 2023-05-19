package types

import (
	"database/sql"
	"encoding/json"
)

type NullFloat64 struct {
	sql.NullFloat64
}

func NewNullFloat64(data float64, zeroNull bool) NullFloat64 {
	// 0是否设置为null
	if zeroNull {
		if data == 0 {
			return NullFloat64{sql.NullFloat64{data, false}}
		}
	}

	return NullFloat64{sql.NullFloat64{data, true}}
}

func (v *NullFloat64) IsZero() bool {
	if !v.Valid || v.Float64 <= 0 {
		return true
	}

	return false
}

func (v *NullFloat64) SetZeroNoNull() {
	v.Valid = true
}

func (v *NullFloat64) SetZeroNull() {
	v.Valid = false
}

func (v *NullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullFloat64) UnmarshalJSON(data []byte) error {
	var s *float64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Float64 = *s
	} else {
		v.Valid = false
	}
	return nil
}
