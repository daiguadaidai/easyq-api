package helper

import "fmt"

func ConvertRowMapToRows(rowMap []map[string]interface{}, columnNames []string) ([][]interface{}, error) {
	rows := make([][]interface{}, 0, len(rowMap))

	for _, m := range rowMap {
		row := make([]interface{}, 0, len(columnNames))
		for _, columnName := range columnNames {
			v, ok := m[columnName]
			if !ok {
				return nil, fmt.Errorf("数据中没有匹配到字段名. %v", columnName)
			}
			switch floatValue := v.(type) {
			case float64:
				if floatValue == float64(int64(floatValue)) {
					// 是整数
					v = int(int64(floatValue))
				}
			case float32:
				if floatValue == float32(int64(floatValue)) {
					// 是整数
					v = int(int64(floatValue))
				}
			}
			row = append(row, v)
		}
		rows = append(rows, row)
	}

	return rows, nil
}
