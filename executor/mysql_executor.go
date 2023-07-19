package executor

import (
	"context"
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/gdbc"
	"github.com/daiguadaidai/easyq-api/utils"
	"sync"
	"time"
)

type MysqlExcutor struct {
	mysqlCfg    *config.MysqlConfig
	execConfig  *config.ExecConfig
	query       string
	successChan chan struct{}
	ctx         context.Context
	cancel      context.CancelFunc
	err         error
}

func NewMysqlExcutor(execConfig *config.ExecConfig, mysqlCfg *config.MysqlConfig, query string) *MysqlExcutor {
	ctx, cancel := context.WithCancel(context.Background())

	return &MysqlExcutor{
		execConfig:  execConfig,
		mysqlCfg:    mysqlCfg,
		query:       query,
		successChan: make(chan struct{}),
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (this *MysqlExcutor) Execute() ([]map[string]interface{}, []string, error) {
	defer func() {
		this.cancel()
	}()

	// 获取数据库链接
	db, err := gdbc.GetMySQLDB(this.mysqlCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("获取执行数据库链接出错. %v", err.Error())
	}

	// 获取数据库链接id
	dbOpDao := dao.NewDBOperationDao(db)
	defer dbOpDao.Close()

	threadId, err := dbOpDao.GetThreadId()
	if err != nil {
		return nil, nil, fmt.Errorf("获取链接 threadId 出错. %v", err.Error())
	}

	wg := new(sync.WaitGroup)
	// 获取 kill 链接
	wg.Add(1)
	go this.timeoutAndKill(wg, this.mysqlCfg.MysqlHost, this.mysqlCfg.MysqlPort, threadId)

	// 执行查询获取结果
	rows, columns, err := dbOpDao.QueryRows(this.query)
	if err != nil {
		err = fmt.Errorf("查询失败. %s %s", this.query, err.Error())
	}


	// 通知执行完成
	close(this.successChan)
	wg.Wait()

	return rows, columns, utils.ErrorsToError(err, this.err)
}

// 监听超时 并执行kill操作
func (this *MysqlExcutor) timeoutAndKill(wg *sync.WaitGroup, host string, port, threadId int64) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(this.execConfig.ExecMysqlExecTimeout)*time.Second)
	defer func() {
		cancel()
	}()

	for {
		select {
		case <-this.ctx.Done(): // 被主动强制退出
			// kill 链接
			if err := KillThread(this.mysqlCfg, threadId); err != nil {
				this.err = fmt.Errorf("强制退出查询 kill 失败. %s 请及时联系DBA查看该实例的校验语句是否还在执行. 如果还在执行请及时kill掉, 防止出现长时间慢查询", err.Error())
			}
			return
		case <-ctx.Done(): // 超时强制退出
			// kill 链接
			if err := KillThread(this.mysqlCfg, threadId); err != nil {
				this.err = fmt.Errorf("执行时间超过: %ds 自动 kill 失败. %s 请及时联系DBA查看该实例的校验语句是否还在执行. 如果还在执行请及时kill掉, 防止出现长时间慢查询", this.timeoutAndKill, err.Error())
			}
			return
		case <-this.successChan:
			return
		}
	}
}
