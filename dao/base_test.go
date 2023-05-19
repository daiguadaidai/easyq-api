package dao

import (
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/gdbc"
	"gorm.io/gorm"
)

var mysqlconfig *config.MysqlConfig = &config.MysqlConfig{
	MysqlHost:              "127.0.0.1",
	MysqlPort:              3306,
	MysqlUsername:          "easydb",
	MysqlPassword:          "easydb",
	MysqlDatabase:          "easydb",
	MysqlConnTimeout:       5,
	MysqlCharset:           "utf8mb4",
	MysqlMaxOpenConns:      8,
	MysqlMaxIdleConns:      7,
	MysqlAllowOldPasswords: 1,
	MysqlAutoCommit:        true,
}

func getGormDB() (*gorm.DB, error) {
	return gdbc.GetOrmDB(mysqlconfig)
}
