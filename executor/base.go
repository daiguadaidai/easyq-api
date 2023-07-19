package executor

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/gdbc"
)

func KillThread(mysqlCfg *config.MysqlConfig, threadId int64) error {
	// 打开一个数据库链接
	db, err := gdbc.GetMySQLDB(mysqlCfg)
	if err != nil {
		return fmt.Errorf("kill %d 出错(打开数据库链接). %s:%d. %s", threadId, mysqlCfg.MysqlHost, mysqlCfg.MysqlPort, err.Error())
	}
	defer db.Close()

	if err := dao.NewDBOperationDao(db).Kill(threadId); err != nil {
		return fmt.Errorf("执行 kill %d; 出错. %s:%d. %s", threadId, mysqlCfg.MysqlHost, mysqlCfg.MysqlPort, err.Error())
	}

	return nil
}
