package gdbc

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetOrmDB(mysqlConfig *config.MysqlConfig) (*gorm.DB, error) {
	dataSource, err, warn := mysqlConfig.GetDataSource()
	if err != nil {
		return nil, fmt.Errorf("获取数据源出错: %s", err.Error())
	}
	if warn != nil {
		logger.M.Warnf("获取数据源警告: %s", warn.Error())
	}

	// 链接数据库
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("打开ORM数据库实例错误. %v", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("生成gorm DB失败. %v", err.Error())
	}
	sqlDB.SetMaxOpenConns(mysqlConfig.MysqlMaxOpenConns)
	sqlDB.SetMaxIdleConns(mysqlConfig.MysqlMaxIdleConns)
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("初始化ORM数据库 ping 失败. %v", err.Error())
	}

	return db, nil
}
