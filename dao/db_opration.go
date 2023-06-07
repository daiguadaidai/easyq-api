package dao

import (
	"database/sql"
	"fmt"
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
