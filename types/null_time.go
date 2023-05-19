package types

import (
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"time"
)

type NullTime struct {
	mysql.NullTime
}

func NewNullTime(data time.Time) NullTime {
	return NullTime{mysql.NullTime{data, true}}
}

func (v *NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		b := make([]byte, 0, len(timeFormat)+2)
		b = append(b, '"')
		b = time.Time(v.Time).AppendFormat(b, timeFormat)
		b = append(b, '"')
		return b, nil
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*v = NewNullTime(now)
	v.Valid = true
	return
}
