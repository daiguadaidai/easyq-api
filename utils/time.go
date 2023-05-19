package utils

import (
	"github.com/daiguadaidai/easyq-api/types"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func StrToTime(formatStr, timeStr string, zone *time.Location) (types.Time, error) {
	datetime, err := time.ParseInLocation(formatStr, timeStr, zone)
	if err != nil {
		return types.Time{}, err
	}

	return types.Time(datetime), nil
}
