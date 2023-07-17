package executor

import (
	"context"
	"github.com/daiguadaidai/easyq-api/config"
)

type MysqlExcutor struct {
	mysqlCfg    *config.MysqlConfig
	query       string
	successChan chan struct{}
	ctx         context.Context
	cancel      context.CancelFunc
}

func NewMysqlExcutor(mysqlCfg *config.MysqlConfig, query string) *MysqlExcutor {
	ctx, cancel := context.WithCancel(context.Background())

	return &MysqlExcutor{
		mysqlCfg:    mysqlCfg,
		query:       query,
		successChan: make(chan struct{}),
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (this *MysqlExcutor) Execute() ([]string, []map[string]interface{}, error) {
	defer func() {
		this.cancel()
	}()

	return nil, nil, nil
}
