package utils

import (
	"time"
)

const (
	// 2023-07-19T14:44:45.439291+08:00
	DatetimeFormat  = "2006-01-02 15:04:05"
	DatetimeFormat1 = "2006-01-02 15:04:05.0"
	DatetimeFormat2 = "2006-01-02 15:04:05.00"
	DatetimeFormat3 = "2006-01-02 15:04:05.000"
	DatetimeFormat4 = "2006-01-02 15:04:05.0000"
	DatetimeFormat5 = "2006-01-02 15:04:05.00000"
	DatetimeFormat6 = "2006-01-02 15:04:05.000000"

	NationalDatetimeFormat  = "2006-01-02T15:04:05Z07:00"
	NationalDatetimeFormat1 = "2006-01-02T15:04:05.0Z07:00"
	NationalDatetimeFormat2 = "2006-01-02T15:04:05.00Z07:00"
	NationalDatetimeFormat3 = "2006-01-02T15:04:05.000Z07:00"
	NationalDatetimeFormat4 = "2006-01-02T15:04:05.0000Z07:00"
	NationalDatetimeFormat5 = "2006-01-02T15:04:05.00000Z07:00"
	NationalDatetimeFormat6 = "2006-01-02T15:04:05.000000Z07:00"

	DateFormat = "2006-01-02"

	TimeFormat  = "15:04:05"
	TimeFormat1 = "15:04:05.0"
	TimeFormat2 = "15:04:05.00"
	TimeFormat3 = "15:04:05.000"
	TimeFormat4 = "15:04:05.0000"
	TimeFormat5 = "15:04:05.00000"
	TimeFormat6 = "15:04:05.000000"
)

func StrToTime(formatStr, timeStr string, zone *time.Location) (time.Time, error) {
	return time.ParseInLocation(formatStr, timeStr, zone)
}

func SqlConvertStrToNormalDatetimeStr(ori string, scale int64) (string, error) {
	datetimeFormat := ""
	nationalDatetimeFormat := ""

	switch scale {
	case 0:
		datetimeFormat = DatetimeFormat
		nationalDatetimeFormat = NationalDatetimeFormat
	case 1:
		datetimeFormat = DatetimeFormat1
		nationalDatetimeFormat = NationalDatetimeFormat1
	case 2:
		datetimeFormat = DatetimeFormat2
		nationalDatetimeFormat = NationalDatetimeFormat2
	case 3:
		datetimeFormat = DatetimeFormat3
		nationalDatetimeFormat = NationalDatetimeFormat3
	case 4:
		datetimeFormat = DatetimeFormat4
		nationalDatetimeFormat = NationalDatetimeFormat4
	case 5:
		datetimeFormat = DatetimeFormat5
		nationalDatetimeFormat = NationalDatetimeFormat5
	case 6:
		datetimeFormat = DatetimeFormat6
		nationalDatetimeFormat = NationalDatetimeFormat6
	default:
		datetimeFormat = DatetimeFormat6
		nationalDatetimeFormat = NationalDatetimeFormat6
	}

	datetime, err := StrToTime(nationalDatetimeFormat, ori, time.Local)
	if err != nil {
		return "", err
	}

	return datetime.Format(datetimeFormat), nil
}

func SqlConvertStrToNormalDateStr(ori string, scale int64) (string, error) {
	nationalDatetimeFormat := ""

	switch scale {
	case 0:
		nationalDatetimeFormat = NationalDatetimeFormat
	case 1:
		nationalDatetimeFormat = NationalDatetimeFormat1
	case 2:
		nationalDatetimeFormat = NationalDatetimeFormat2
	case 3:
		nationalDatetimeFormat = NationalDatetimeFormat3
	case 4:
		nationalDatetimeFormat = NationalDatetimeFormat4
	case 5:
		nationalDatetimeFormat = NationalDatetimeFormat5
	case 6:
		nationalDatetimeFormat = NationalDatetimeFormat6
	default:
		nationalDatetimeFormat = NationalDatetimeFormat6
	}

	datetime, err := StrToTime(nationalDatetimeFormat, ori, time.Local)
	if err != nil {
		return "", err
	}

	return datetime.Format(DateFormat), nil
}

func SqlConvertStrToNormalTimeStr(ori string, scale int64) (string, error) {
	timeFormat := ""
	nationalTimeFormat := ""

	switch scale {
	case 0:
		timeFormat = TimeFormat
		nationalTimeFormat = NationalDatetimeFormat
	case 1:
		timeFormat = TimeFormat1
		nationalTimeFormat = NationalDatetimeFormat1
	case 2:
		timeFormat = TimeFormat2
		nationalTimeFormat = NationalDatetimeFormat2
	case 3:
		timeFormat = TimeFormat3
		nationalTimeFormat = NationalDatetimeFormat3
	case 4:
		timeFormat = TimeFormat4
		nationalTimeFormat = NationalDatetimeFormat4
	case 5:
		timeFormat = TimeFormat5
		nationalTimeFormat = NationalDatetimeFormat5
	case 6:
		timeFormat = TimeFormat6
		nationalTimeFormat = NationalDatetimeFormat6
	default:
		timeFormat = TimeFormat6
		nationalTimeFormat = NationalDatetimeFormat6
	}

	datetime, err := StrToTime(nationalTimeFormat, ori, time.Local)
	if err != nil {
		return "", err
	}

	return datetime.Format(timeFormat), nil
}
