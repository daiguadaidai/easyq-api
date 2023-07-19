package dao

import (
	"database/sql"
	"fmt"
	"github.com/daiguadaidai/easyq-api/utils"
)

type DBOperationDao struct {
	DB *sql.DB
}

func NewDBOperationDao(db *sql.DB) *DBOperationDao {
	return &DBOperationDao{
		DB: db,
	}
}

func (this *DBOperationDao) Close() {
	if this.DB != nil {
		this.DB.Close()
	}
}

func (this *DBOperationDao) ShowDatabases() ([]string, error) {
	query := "SHOW DATABASES"

	rows, err := this.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("SHOW DATABASES 查询出错. %v", err)
	}
	defer rows.Close()

	var dbName string
	dbNames := make([]string, 0, 10)
	for rows.Next() {
		err = rows.Scan(&dbName) //不scan会导致连接不释放
		if err != nil {
			return nil, fmt.Errorf("Scan failed, err: %v", err)
		}

		dbNames = append(dbNames, dbName)
	}

	return dbNames, nil
}

func (this *DBOperationDao) ShowTables() ([]string, error) {
	query := "SHOW TABLES"

	rows, err := this.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("SHOW TABLES 查询出错. %v", err)
	}
	defer rows.Close()

	var tableName string
	tableNames := make([]string, 0, 10)
	for rows.Next() {
		err = rows.Scan(&tableName) //不scan会导致连接不释放
		if err != nil {
			return nil, fmt.Errorf("Scan failed, err: %v", err)
		}

		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}

/*
	查询

Return

	[]map[string]interface{}: 返回数据行 rows
	[]string: 返回字段名 columns
	error: 错
*/
func (this *DBOperationDao) QueryRows(query string) ([]map[string]interface{}, []string, error) {
	rows, err := this.DB.Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("执行查询语句出错. %s. %s", query, err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, fmt.Errorf("获取查询结果字段出错. %s. %s", query, err.Error())
	}

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, columns, fmt.Errorf("获取查询结果字段类型出错. %s. %s", query, err.Error())
	}

	values := make([]*sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	resultsz := make([]map[string]interface{}, 0, 20)
	for rows.Next() {
		results := make(map[string]interface{})
		if err = rows.Scan(scans...); err != nil {
			return nil, columns, fmt.Errorf("scan 结果出错. %s. %s", query, err.Error())
		}

		for i, value := range values {
			columnName := columns[i]

			if value == nil {
				results[columnName] = nil
				continue
			}
			// 将 rowbyte转化称想要的类型
			results[columnName], err = utils.ConvertAssign(*value, colTypes[i])
			if err != nil {
				return nil, columns, fmt.Errorf("将 rowbytes 结果转化为需要的类型的值出错. value: %v. type: %s. %s", *value, colTypes[i], err.Error())
			}
		}
		resultsz = append(resultsz, results)
	}

	return resultsz, columns, nil
}

// 获取链接ID
func (this *DBOperationDao) GetThreadId() (int64, error) {
	query := `SELECT CONNECTION_ID()`

	var threadId int64
	if err := this.DB.QueryRow(query).Scan(&threadId); err != nil {
		return 0, err
	}

	return threadId, nil
}

// Kill 链接
func (this *DBOperationDao) Kill(threadId int64) error {
	query := fmt.Sprintf(`kill %d`, threadId)

	if _, err := this.DB.Exec(query); err != nil {
		return err
	}

	return nil
}
