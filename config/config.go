package config

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/daiguadaidai/easyq-api/utils"
	"strings"
)

type ServerConfig struct {
	ApiConfig         *ApiConfig   `json:"api_config" toml:"api_config"`
	LogConfig         *LogConfig   `json:"log_config" toml:"log_config"`
	EasyQMysqlConfig  *MysqlConfig `json:"easyq_mysql_config" toml:"easyq_mysql_config"`
	EasyDBMysqlConfig *MysqlConfig `json:"easydb_mysql_config" toml:"easydb_mysql_config"`
	ExecConfig        *ExecConfig  `json:"exec_config" toml:"exec_config"`
}

func NewServerConfig(
	apiConfig *ApiConfig,
	logConfig *LogConfig,
	easyQMysqlConfig *MysqlConfig,
	easyDBMysqlConfig *MysqlConfig,
	execConfig *ExecConfig,
) *ServerConfig {
	return &ServerConfig{
		ApiConfig:         apiConfig,
		LogConfig:         logConfig,
		EasyQMysqlConfig:  easyQMysqlConfig,
		EasyDBMysqlConfig: easyDBMysqlConfig,
		ExecConfig:        execConfig,
	}
}

func (this *ServerConfig) DeepClone() (*ServerConfig, error) {
	raw, err := json.Marshal(this)
	if err != nil {
		return nil, fmt.Errorf("启动配置DeepClone出错, ServerConfig -> Json: %v", err.Error())
	}

	var serverConfig ServerConfig
	err = json.Unmarshal(raw, &serverConfig)
	if err != nil {
		return nil, fmt.Errorf("启动配置DeepClone出错, Json -> ServerConfig: %v", err.Error())
	}

	return &serverConfig, nil
}

// 配置生效优先级: 命令行 > toml > 默认
func GetConfigs(serverConfigCmd *ServerConfig, changeCmdKeys []string) (*ServerConfig, error) {
	configFile := strings.TrimSpace(serverConfigCmd.ApiConfig.ConfigFile)
	// 没有指定配置文件, 直接使用命令行参数
	if configFile == "" {
		return serverConfigCmd, nil
	}

	// 指定了配置文件, 配置信息冲配置文件中读取
	serverConfigToml, err := GetConfigsFromToml(configFile)
	if err != nil {
		return nil, err
	}

	// 生成最终的配置信息, 配置生效优先级: 命令行 > toml > 默认
	serverConfigNew, err := newConfigFromConsoleAndToml(serverConfigCmd, serverConfigToml, changeCmdKeys)
	if err != nil {
		return nil, err
	}

	return serverConfigNew, nil
}

func GetConfigsFromToml(filename string) (*ServerConfig, error) {
	var serverConfig ServerConfig
	if _, err := toml.DecodeFile(filename, &serverConfig); err != nil {
		return nil, fmt.Errorf("读取启动配置失败(ServerConfig), %v \n", err.Error())
	}

	return &serverConfig, nil
}

func newConfigFromConsoleAndToml(
	serverConfigCmd, serverConfigToml *ServerConfig,
	changeCmdKeys []string, // 在命令行中指定的参数字符串
) (*ServerConfig, error) {
	serverConfigNew, err := serverConfigToml.DeepClone()
	if err != nil {
		return nil, err
	}

	serverConfigNew.ApiConfig.ConfigFile = serverConfigCmd.ApiConfig.ConfigFile

	// 将命令变成蛇形
	snakeCmdKeys := utils.ReplaceAllStrs(changeCmdKeys, "-", "_")
	// 将蛇形变为驼峰
	camelCmdKeys := utils.CamelStrs(snakeCmdKeys)

	// 获取startConfig需要修改的字段名称
	apiConfigFieldNames, err := utils.GetStructFieldNames(serverConfigCmd.ApiConfig)
	if err != nil {
		return nil, fmt.Errorf("获取ApiConfig所有字段失败. %v", err.Error())
	}
	// 获取是ApiConfig的命令行输入参数
	apiConfigCmdKeys := utils.GetExistsStrs(camelCmdKeys, apiConfigFieldNames)
	// 设置命令行手动指定开始参数
	if err := utils.SetObjFieldsFromOtherObj(apiConfigCmdKeys, serverConfigCmd.ApiConfig, serverConfigNew.ApiConfig); err != nil {
		return nil, fmt.Errorf("设置ApiConfig命令行设置值. %v", err.Error())
	}

	// 获取logConfig需要修改的字段名称
	logConfigFieldNames, err := utils.GetStructFieldNames(serverConfigCmd.LogConfig)
	if err != nil {
		return nil, fmt.Errorf("获取LogConfig所有字段失败. %v", err.Error())
	}
	// 获取是LogConfig的命令行输入参数
	logConfigCmdKeys := utils.GetExistsStrs(camelCmdKeys, logConfigFieldNames)
	// 设置命令行手动指定日志参数
	if err := utils.SetObjFieldsFromOtherObj(logConfigCmdKeys, serverConfigCmd.LogConfig, serverConfigNew.LogConfig); err != nil {
		return nil, fmt.Errorf("设置LogConfig命令行设置值. %v", err.Error())
	}

	// 获取EasyQMysqlConfig需要修改的字段名称
	easyqMysqlConfigFieldNames, err := utils.GetStructFieldNames(serverConfigCmd.EasyQMysqlConfig)
	if err != nil {
		return nil, fmt.Errorf("获取EasyQMysqlConfig所有字段失败. %v", err.Error())
	}
	// 获取是EasyQMysqlConfig的命令行输入参数
	easyqMysqlConfigCmdKeys := utils.GetExistsStrsPrefix(camelCmdKeys, easyqMysqlConfigFieldNames, "Easyq")
	// 设置命令行手动指定日志参数
	if err := utils.SetObjFieldsFromOtherObj(easyqMysqlConfigCmdKeys, serverConfigCmd.EasyQMysqlConfig, serverConfigNew.EasyQMysqlConfig); err != nil {
		return nil, fmt.Errorf("设置EasyQMysqlConfig命令行设置值. %v", err.Error())
	}

	// 获取EasyDBMysqlConfig需要修改的字段名称
	easydbMysqlConfigFieldNames, err := utils.GetStructFieldNames(serverConfigCmd.EasyDBMysqlConfig)
	if err != nil {
		return nil, fmt.Errorf("获取EasyDBMysqlConfig所有字段失败. %v", err.Error())
	}
	// 获取是EasyDBMysqlConfig的命令行输入参数
	easydbMysqlConfigCmdKeys := utils.GetExistsStrsPrefix(camelCmdKeys, easydbMysqlConfigFieldNames, "Easydb")
	// 设置命令行手动指定日志参数
	if err := utils.SetObjFieldsFromOtherObj(easydbMysqlConfigCmdKeys, serverConfigCmd.EasyDBMysqlConfig, serverConfigNew.EasyDBMysqlConfig); err != nil {
		return nil, fmt.Errorf("设置EasyDBMysqlConfig命令行设置值. %v", err.Error())
	}

	// 获取ExecConfig需要修改的字段名称
	execConfigFieldNames, err := utils.GetStructFieldNames(serverConfigCmd.ExecConfig)
	if err != nil {
		return nil, fmt.Errorf("获取ExecConfig所有字段失败. %v", err.Error())
	}
	// 获取是ExecConfig的命令行输入参数
	execConfigCmdKeys := utils.GetExistsStrs(camelCmdKeys, execConfigFieldNames)
	// 设置命令行手动指定日志参数
	if err := utils.SetObjFieldsFromOtherObj(execConfigCmdKeys, serverConfigCmd.ExecConfig, serverConfigNew.ExecConfig); err != nil {
		return nil, fmt.Errorf("设置ExecConfig命令行设置值. %v", err.Error())
	}

	return serverConfigNew, nil
}
