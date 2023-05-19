package services

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/gdbc"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/utils"
	"sync"
)

func RunApi(
	changeCmdKeys []string,
	apiConfigCmd *config.ApiConfig,
	logConfigCmd *config.LogConfig,
	easyqMysqlConfigCmd *config.MysqlConfig,
	easydbMysqlConfigCmd *config.MysqlConfig,
) error {
	serverConfigCmd := config.NewServerConfig(apiConfigCmd, logConfigCmd, easyqMysqlConfigCmd, easydbMysqlConfigCmd)
	serverConfig, err := config.GetConfigs(serverConfigCmd, changeCmdKeys)
	if err != nil {
		return fmt.Errorf("创建最终的配置文件出错. %v", err.Error())
	}

	// 初始化日志系统
	logger.InitLogger(serverConfig.LogConfig)

	logger.M.Info(utils.ToJsonStrPretty(serverConfig))

	// 获取 globalContext
	ctx, err := getCtx(serverConfig)
	if err != nil {
		logger.M.Errorf(err.Error())
		return err
	}

	wg := new(sync.WaitGroup)
	// 启动api service
	wg.Add(1)
	StartApiService(wg, ctx)

	go wg.Wait()

	return nil
}

func getCtx(serverConfig *config.ServerConfig) (*contexts.GlobalContext, error) {
	// 获取Easyq MySQL数据库连接池
	easyqDB, err := gdbc.GetOrmDB(serverConfig.EasyQMysqlConfig)
	if err != nil {
		return nil, err
	}
	logger.M.Infof("创建EasyQ gorm数据库连接池成功: %s", serverConfig.EasyQMysqlConfig.GetFuzzyDataSource())

	// 获取Easydb MySQL数据库连接池
	easydbDB, err := gdbc.GetOrmDB(serverConfig.EasyQMysqlConfig)
	if err != nil {
		return nil, err
	}
	logger.M.Infof("创建EasyDB gorm数据库连接池成功: %s", serverConfig.EasyDBMysqlConfig.GetFuzzyDataSource())

	return &contexts.GlobalContext{
		EasydbDB: easydbDB,
		EasyqDB:  easyqDB,
		Cfg:      serverConfig,
	}, nil
}
