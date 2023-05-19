package gdbc

import (
	"database/sql"
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/logger"
	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB(mysqlConfig *config.MysqlConfig) (*sql.DB, error) {
	dataSource, err, warn := mysqlConfig.GetDataSource()
	if err != nil {
		return nil, fmt.Errorf("获取数据源出错: %s", err.Error())
	}
	if warn != nil {
		logger.M.Warnf("获取数据源警告: %s", warn.Error())
	}

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("获取打开数据库失败. %s", err.Error())
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping数据库失败. %s", err.Error())
	}

	return db, nil
}
