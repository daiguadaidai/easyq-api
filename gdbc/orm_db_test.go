package gdbc

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"sync"
	"testing"
)

var mysqlconfig *config.MysqlConfig = &config.MysqlConfig{
	MysqlHost:              "127.0.0.1",
	MysqlPort:              3306,
	MysqlUsername:          "easydb",
	MysqlPassword:          "easydb",
	MysqlDatabase:          "employees",
	MysqlConnTimeout:       5,
	MysqlCharset:           "utf8mb4",
	MysqlMaxOpenConns:      8,
	MysqlMaxIdleConns:      7,
	MysqlAllowOldPasswords: 1,
	MysqlAutoCommit:        true,
}

// 测试并发获取数据库链接(使用了单例模式)
func TestGetOrmInstance(t *testing.T) {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(_wg *sync.WaitGroup) {
			defer _wg.Done()

			db, err := GetOrmDB(mysqlconfig)
			if err != nil {
				t.Fatalf("创建数据库连接池失败: %v", err.Error())
			}

			// 查询数据库
			tmpDB, err := db.DB()
			if err != nil {
				t.Fatal(err.Error())
			}
			rows, err := tmpDB.Query("SELECT * FROM employees limit 1")
			if err != nil {
				t.Fatal(err.Error())
			}
			defer rows.Close()
			fmt.Println(rows)
		}(wg)
	}

	wg.Wait()
}
