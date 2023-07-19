package config

import (
	"encoding/json"
	"fmt"
)

const (
	DefaultExecMysqlExecTimeout = 30   // mysql默认执行超时时间
	DefaultExecMysqlSelectLimit = 2000 // 默认执行select limit 值
)

type ExecConfig struct {
	ExecMysqlExecTimeout int64 `json:"exec_mysql_exec_timeout" toml:"exec_mysql_exec_timeout"`
	ExecMysqlSelectLimit int64 `json:"exec_mysql_select_limit" toml:"exec_mysql_select_limit"`
}

func (this *ExecConfig) DeepClone() (*ExecConfig, error) {
	raw, err := json.Marshal(this)
	if err != nil {
		return nil, fmt.Errorf("启动配置DeepClone出错, ExecConfig -> Json: %v", err.Error())
	}

	var execConfig ExecConfig
	err = json.Unmarshal(raw, &execConfig)
	if err != nil {
		return nil, fmt.Errorf("启动配置DeepClone出错, Json -> ExecConfig: %v", err.Error())
	}

	return &execConfig, nil
}
