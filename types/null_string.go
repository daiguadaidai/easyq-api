package types

import (
	"database/sql"
	"encoding/json"
	"strings"
)

type NullString struct {
	sql.NullString
}

func NewNullString(data string, emptyNull bool) NullString {
	// 空字符串设置为null
	if emptyNull && data == "" {
		return NullString{sql.NullString{data, false}}
	}
	return NullString{sql.NullString{data, true}}
}

func (v *NullString) IsEmpty() bool {
	if !v.Valid || strings.TrimSpace(v.String) == "" {
		return true
	}
	return false
}

func (v *NullString) SetEmptyNoNull() {
	v.Valid = true
}

func (v *NullString) SetEmptyNull() {
	v.Valid = false
}

func (v *NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.String = *s
	} else {
		v.Valid = false
	}
	return nil
}

func JoinStr(strs []NullString, sep string) string {
	newStrs := make([]string, 0, len(strs))
	for _, str := range strs {
		newStrs = append(newStrs, str.String)
	}

	return strings.Join(newStrs, sep)
}
