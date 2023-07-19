/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/services"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/spf13/cobra"
	"log"
)

var apiConfig *config.ApiConfig
var easyqMysqlConfig *config.MysqlConfig
var easydbMysqlConfig *config.MysqlConfig
var logConfig *config.LogConfig
var execConfig *config.ExecConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "easyq-api",
	Short: "数据库查询平台API接口",
	Long: `数据库查询平台API接口
Example:
# 通过配置文件启动
./easyq-api --config=./easyq_api.toml

# 通过命令行启动
./easydb-api \
    --listen-host="127.0.0.1" \
    --listen-port=9101 \
    --env="prod" \
    --token-expire=86400 \
    --query-mysql-user="root" \
    --query-mysql-password="root" \
    --admin-mysql-user="root" \
    --admin-mysql-password="root" \
	--forward-request-dial-timeout=10 \
	--forward-request-response-header-timeout=5 \
    --easyq-mysql-host=127.0.0.1 \
    --easyq-mysql-port=3306 \
    --easyq-mysql-username="easydb" \
    --easyq-mysql-password="easydb" \
    --easyq-mysql-database="easydb" \
    --easyq-mysql-conn-timeout=5 \
    --easyq-mysql-charset="utf8mb4" \
    --easyq-mysql-max-open-conns=8 \
    --easyq-mysql-max-idle-conns=7 \
    --easyq-mysql-allow-old-password=1 \
    --easyq-mysql-auto-commit=true \
    --easydb-mysql-host=127.0.0.1 \
    --easydb-mysql-port=3306 \
    --easydb-mysql-username="easydb" \
    --easydb-mysql-password="easydb" \
    --easydb-mysql-database="easydb" \
    --easydb-mysql-conn-timeout=5 \
    --easydb-mysql-charset="utf8mb4" \
    --easydb-mysql-max-open-conns=8 \
    --easydb-mysql-max-idle-conns=7 \
    --easydb-mysql-allow-old-password=1 \
    --easydb-mysql-auto-commit=true \
    --exec-mysql-exec-timeout=30 \
    --exec-mysql-select-limit=2000 \
    --log-filename="logs/easyq_api.log" \
    --log-level="info" \
    --log-max-size=1024 \
    --log-max-backups=10 \
    --log-max-age=30 \
    --log-compress=false \
    --log-console=false \
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// 获取所有的 命令行key
		allCmdKeys, err := getAllCmdKeys(apiConfig, logConfig, easyqMysqlConfig, easydbMysqlConfig, execConfig)
		if err != nil {
			log.Fatalln(err.Error())
		}

		// 获取手动指定的命令
		changeCmdKeys := getChangeCmdStrs(cmd, allCmdKeys)

		if err := services.RunApi(changeCmdKeys, apiConfig, logConfig, easyqMysqlConfig, easydbMysqlConfig, execConfig); err != nil {
			log.Fatalln(err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	initStartConfig()
	initLogConfig()
	initEasyQMysqlConfig()
	initEasyDBMysqlConfig()
	initExecConfig()
}

func initStartConfig() {
	apiConfig = new(config.ApiConfig)

	rootCmd.PersistentFlags().StringVar(&apiConfig.ConfigFile, "config", config.DefaultConfigFile, "默认配置文件路径, 文件格式toml")

	rootCmd.PersistentFlags().StringVar(&apiConfig.ListenHost, "listener-host", config.DefaultListenHost, "Api监听host")
	rootCmd.PersistentFlags().Int64Var(&apiConfig.ListenPort, "listener-port", config.DefaultListenPort, "Api监听port")
	rootCmd.PersistentFlags().StringVar(&apiConfig.Env, "env", config.DefaultEnv, "环境: dev, prod")
	rootCmd.PersistentFlags().Int64Var(&apiConfig.TokenExpire, "token-expire", config.DefaultTokenExpire, "Api Token 过期时间")
	rootCmd.PersistentFlags().StringVar(&apiConfig.QueryMysqlUser, "query-mysql-user", config.DefaultQueryMySQLUser, "查询数据库用户名")
	rootCmd.PersistentFlags().StringVar(&apiConfig.QueryMysqlPassword, "query-mysql-password", config.DefaultQueryMySQLPassword, "查询数据库密码")
	rootCmd.PersistentFlags().StringVar(&apiConfig.AdminMysqlUser, "admin-mysql-user", config.DefaultAdminMySQLUser, "管理数据库用户名")
	rootCmd.PersistentFlags().StringVar(&apiConfig.AdminMysqlPassword, "admin-mysql-password", config.DefaultAdminMySQLPassword, "管理数据库密码")
	rootCmd.PersistentFlags().Int64Var(&apiConfig.ForwardRequestDialTimeout, "forward-request-dial-timeout", config.DefaultForwardRequestDialTimeout, "请求转发链接超时时间")
	rootCmd.PersistentFlags().Int64Var(&apiConfig.ForwardRequestResponseHeaderTimeout, "forward-request-response-header-timeout", config.DefaultForwardRequestResponseHeaderTimeout, "请求转发Response超时时间")
}

func initEasyQMysqlConfig() {
	easyqMysqlConfig = new(config.MysqlConfig)

	rootCmd.PersistentFlags().StringVar(&easyqMysqlConfig.MysqlHost, "easyq-mysql-host", config.DefaultMysqlHost, "EasyQ Mysql默认链接使用的host")
	rootCmd.PersistentFlags().Int64Var(&easyqMysqlConfig.MysqlPort, "easyq-mysql-port", config.DefaultMysqlPort, "EasyQ Mysql默认需要链接的端口, 如果没有指定则动态通过命令获取")
	rootCmd.PersistentFlags().StringVar(&easyqMysqlConfig.MysqlUsername, "easyq-mysql-username", config.DefaultMysqlUsername, "EasyQ Mysql链接的用户名")
	rootCmd.PersistentFlags().StringVar(&easyqMysqlConfig.MysqlPassword, "easyq-mysql-password", config.DefaultMysqlPassword, "EasyQ Mysql链接的密码")
	rootCmd.PersistentFlags().StringVar(&easyqMysqlConfig.MysqlDatabase, "easyq-mysql-database", config.DefaultMysqlDatabase, "EasyQ Mysql链接的数据库名称")
	rootCmd.PersistentFlags().IntVar(&easyqMysqlConfig.MysqlConnTimeout, "easyq-mysql-conn-timeout", config.DefaultMysqlConnTimeout, "EasyQ Mysql链接超时时间. 单位(s)")
	rootCmd.PersistentFlags().StringVar(&easyqMysqlConfig.MysqlCharset, "easyq-mysql-charset", config.DefaultMysqlCharset, "EasyQ Mysql链接字符集")
	rootCmd.PersistentFlags().IntVar(&easyqMysqlConfig.MysqlMaxOpenConns, "easyq-mysql-max-open-conns", config.DefaultMysqlMaxOpenConns, "EasyQ Mysql最大链接数")
	rootCmd.PersistentFlags().IntVar(&easyqMysqlConfig.MysqlMaxIdleConns, "easyq-mysql-max-idle-conns", config.DefaultMysqlMaxIdleConns, "EasyQ Mysql最大空闲链接数")
	rootCmd.PersistentFlags().IntVar(&easyqMysqlConfig.MysqlAllowOldPasswords, "easyq-mysql-allow-old-passwords", config.DefaultMysqlAllowOldPasswords, "EasyQ Mysql是否兼容老密码链接方式")
	rootCmd.PersistentFlags().BoolVar(&easyqMysqlConfig.MysqlAutoCommit, "easyq-mysql-auto-commit", config.DefaultMysqlAutoCommit, "EasyQ Mysql自动提交")
}

func initEasyDBMysqlConfig() {
	easydbMysqlConfig = new(config.MysqlConfig)

	rootCmd.PersistentFlags().StringVar(&easydbMysqlConfig.MysqlHost, "easydb-mysql-host", config.DefaultMysqlHost, "EasyDB Mysql默认链接使用的host")
	rootCmd.PersistentFlags().Int64Var(&easydbMysqlConfig.MysqlPort, "easydb-mysql-port", config.DefaultMysqlPort, "EasyDB Mysql默认需要链接的端口, 如果没有指定则动态通过命令获取")
	rootCmd.PersistentFlags().StringVar(&easydbMysqlConfig.MysqlUsername, "easydb-mysql-username", config.DefaultMysqlUsername, "EasyDB Mysql链接的用户名")
	rootCmd.PersistentFlags().StringVar(&easydbMysqlConfig.MysqlPassword, "easydb-mysql-password", config.DefaultMysqlPassword, "EasyDB Mysql链接的密码")
	rootCmd.PersistentFlags().StringVar(&easydbMysqlConfig.MysqlDatabase, "easydb-mysql-database", config.DefaultMysqlDatabase, "EasyDB Mysql链接的数据库名称")
	rootCmd.PersistentFlags().IntVar(&easydbMysqlConfig.MysqlConnTimeout, "easydb-mysql-conn-timeout", config.DefaultMysqlConnTimeout, "EasyDB Mysql链接超时时间. 单位(s)")
	rootCmd.PersistentFlags().StringVar(&easydbMysqlConfig.MysqlCharset, "easydb-mysql-charset", config.DefaultMysqlCharset, "EasyDB Mysql链接字符集")
	rootCmd.PersistentFlags().IntVar(&easydbMysqlConfig.MysqlMaxOpenConns, "easydb-mysql-max-open-conns", config.DefaultMysqlMaxOpenConns, "EasyDB Mysql最大链接数")
	rootCmd.PersistentFlags().IntVar(&easydbMysqlConfig.MysqlMaxIdleConns, "easydb-mysql-max-idle-conns", config.DefaultMysqlMaxIdleConns, "EasyDB Mysql最大空闲链接数")
	rootCmd.PersistentFlags().IntVar(&easydbMysqlConfig.MysqlAllowOldPasswords, "easydb-mysql-allow-old-passwords", config.DefaultMysqlAllowOldPasswords, "EasyDB Mysql是否兼容老密码链接方式")
	rootCmd.PersistentFlags().BoolVar(&easydbMysqlConfig.MysqlAutoCommit, "easydb-mysql-auto-commit", config.DefaultMysqlAutoCommit, "EasyDB Mysql自动提交")
}

func initLogConfig() {
	logConfig = new(config.LogConfig)

	rootCmd.PersistentFlags().StringVar(&logConfig.LogFilename, "log-filename", config.DefaultLogFilename, "默认日志文件路径")
	rootCmd.PersistentFlags().StringVar(&logConfig.LogLevel, "log-level", config.DefaultLogLevel, "日志级别: debug, info, warn, error, dpanic, panic, fatal")
	rootCmd.PersistentFlags().IntVar(&logConfig.LogMaxSize, "log-max-size", config.DefaultLogMaxSize, "每个日志文件最大多大(单位:M)")
	rootCmd.PersistentFlags().IntVar(&logConfig.LogMaxBackups, "log-max-backups", config.DefaultLogMaxBackups, "日志文件最多保存多少个备份")
	rootCmd.PersistentFlags().IntVar(&logConfig.LogMaxAge, "log-max-age", config.DefaultLogMaxAge, "文件最多保存多少天(单位:天)")
	rootCmd.PersistentFlags().BoolVar(&logConfig.LogCompress, "log-compress", config.DefaultLogCompress, "是否压缩")
	rootCmd.PersistentFlags().BoolVar(&logConfig.LogConsole, "log-console", config.DefaultLogConsole, "是否打印到控制台")
}

func initExecConfig() {
	execConfig = new(config.ExecConfig)

	rootCmd.PersistentFlags().Int64Var(&execConfig.ExecMysqlExecTimeout, "exec-mysql-exec-timeout", config.DefaultExecMysqlExecTimeout, "mysql执行超时时间(单位:s)")
	rootCmd.PersistentFlags().Int64Var(&execConfig.ExecMysqlSelectLimit, "exec-mysql-select-limit", config.DefaultExecMysqlSelectLimit, "mysql执行sql select的默认limit数")
}

// 获取所有的命令行key
func getAllCmdKeys(
	apiConfig *config.ApiConfig,
	logConfig *config.LogConfig,
	easyqMysqlConfig *config.MysqlConfig,
	easydbMysqlConfig *config.MysqlConfig,
	execConfig *config.ExecConfig,
) ([]string, error) {
	startConfigFiledNames, err := utils.GetStructFieldNames(apiConfig)
	if err != nil {
		return nil, fmt.Errorf("获取ApiConfig所有字段名称失败. %v", err.Error())
	}
	logConfigFiledNames, err := utils.GetStructFieldNames(logConfig)
	if err != nil {
		return nil, fmt.Errorf("获取LogConfig所有字段名称失败. %v", err.Error())
	}
	easyqMysqlConfigFiledNames, err := utils.GetStructFieldNamesWithPrefix(easyqMysqlConfig, "easyq")
	if err != nil {
		return nil, fmt.Errorf("获取EasyQMysqlConfig所有字段名称失败. %v", err.Error())
	}
	easydbMysqlConfigFiledNames, err := utils.GetStructFieldNamesWithPrefix(easydbMysqlConfig, "easydb")
	if err != nil {
		return nil, fmt.Errorf("获取EasyDBMysqlConfig所有字段名称失败. %v", err.Error())
	}
	execConfigFiledNames, err := utils.GetStructFieldNames(execConfig)
	if err != nil {
		return nil, fmt.Errorf("获取ExecConfig所有字段名称失败. %v", err.Error())
	}

	allFieldNames := make([]string, 0, 10)
	allFieldNames = append(allFieldNames, startConfigFiledNames...)
	allFieldNames = append(allFieldNames, logConfigFiledNames...)
	allFieldNames = append(allFieldNames, easyqMysqlConfigFiledNames...)
	allFieldNames = append(allFieldNames, easydbMysqlConfigFiledNames...)
	allFieldNames = append(allFieldNames, execConfigFiledNames...)

	allSnakeFieldNames := utils.SnakeStrs(allFieldNames)
	allCmdKeys := utils.ReplaceAllStrs(allSnakeFieldNames, "_", "-")

	return allCmdKeys, nil
}

func getChangeCmdStrs(cmd *cobra.Command, cmdKeys []string) []string {
	changeCmdKeys := make([]string, 0, 10)
	f := cmd.Flags()
	for _, cmdKey := range cmdKeys {
		if changed := f.Changed(cmdKey); changed {
			changeCmdKeys = append(changeCmdKeys, cmdKey)
		}
	}

	return changeCmdKeys
}
