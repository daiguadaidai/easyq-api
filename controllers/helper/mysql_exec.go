package helper

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/config"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/executor"
	"github.com/daiguadaidai/easyq-api/gdbc"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/utils"
	"strings"
)

func CheckExecDBPriv(ctx *contexts.GlobalContext, dbNames []string, metaClusterId int64, username string) error {
	privDao := dao.NewMysqlDBPrivDao(ctx.EasyqDB)
	for _, dbName := range dbNames {
		cnt, err := privDao.CountByUsernameClusterDB(username, metaClusterId, dbName)
		if err != nil {
			return fmt.Errorf("获取执行sql中涉及到到数据库权限出错. 数据库: %v. %v", dbName, err.Error())
		}
		if cnt == 0 {
			return fmt.Errorf("没有数据库权限. 数据库: %v", dbName)
		}
	}

	return nil
}

func getExecInstance(ctx *contexts.GlobalContext, priv *models.MysqlDBPriv) (*models.Instance, error) {
	instances, err := dao.NewInstanceDao(ctx.EasydbDB).FindByClusterId(priv.MetaClusterId.Int64)
	if err != nil {
		return nil, fmt.Errorf("获取所有可执行的实例出错. meta_cluster_id: %v, %v", priv.MetaClusterId.Int64, err.Error())
	}

	if len(instances) == 0 {
		return nil, fmt.Errorf("没有获取到可执行到实例. meta_cluster_id: %v", priv.MetaClusterId.Int64)
	}

	// 检测和获取可执行的实例
	masters, slaves, backups, others := splitInstances(instances)

	// 检测backups
	backup, backupErr := checkAndGetInstance(backups, "backup", ctx.Cfg.ApiConfig.QueryMysqlUser, ctx.Cfg.ApiConfig.QueryMysqlPassword, priv.DBName.String)
	if backupErr == nil { // 没有错误说明backup实例可用
		return backup, nil
	}

	// 检测 slaves
	slave, slaveErr := checkAndGetInstance(slaves, "slave", ctx.Cfg.ApiConfig.QueryMysqlUser, ctx.Cfg.ApiConfig.QueryMysqlPassword, priv.DBName.String)
	if slaveErr == nil {
		return slave, nil
	}

	// 检测 master
	master, masterErr := checkAndGetInstance(masters, "master", ctx.Cfg.ApiConfig.QueryMysqlUser, ctx.Cfg.ApiConfig.QueryMysqlPassword, priv.DBName.String)
	if masterErr == nil {
		return master, nil
	}

	// 检测 others
	other, otherErr := checkAndGetInstance(others, "other", ctx.Cfg.ApiConfig.QueryMysqlUser, ctx.Cfg.ApiConfig.QueryMysqlPassword, priv.DBName.String)
	if otherErr == nil {
		logger.M.Warnf("检测到有未知角色实例可以使用, 使用未知角色实例. 用户: %v, 集群id: %v, 数据库: %v, %v", priv.Username.String, priv.MetaClusterId.Int64, priv.DBName.String, utils.ToJsonStr(other))
		return other, nil
	}

	return nil, utils.ErrorsToError(backupErr, slaveErr, masterErr, otherErr)
}

func checkAndGetInstance(instances []*models.Instance, tagStr string, username, password, database string) (*models.Instance, error) {
	if len(instances) == 0 {
		return nil, fmt.Errorf("%v 没有实例", tagStr)
	}

	errStrs := make([]string, 0, len(instances))
	for _, instance := range instances {
		mysqlCfg := config.NewMysqlConfig(instance.MachineHost.String, instance.Port.Int64, username, password, database)
		db, err := gdbc.GetMySQLDB(mysqlCfg)
		if err != nil {
			errStr := fmt.Sprintf("%v 检测实例出错. %v:%v. %v", tagStr, instance.MachineHost.String, instance.Port.Int64, err.Error())
			errStrs = append(errStrs, errStr)
			continue
		}
		db.Close()

		return instance, nil
	}

	return nil, fmt.Errorf("%v", strings.Join(errStrs, "\n"))
}

// 拆分实例, master, slave, backup, 其他角色实例
func splitInstances(instances []*models.Instance) ([]*models.Instance, []*models.Instance, []*models.Instance, []*models.Instance) {
	masters := make([]*models.Instance, 0, 2)
	slaves := make([]*models.Instance, 0, 2)
	backups := make([]*models.Instance, 0, 2)
	others := make([]*models.Instance, 0, 2)

	for _, instance := range instances {
		if instance.Role.String == models.InstanceRoleMaster {
			masters = append(masters, instance)
		} else if instance.Role.String == models.InstanceRoleSlave {
			slaves = append(slaves, instance)
		} else if instance.Role.String == models.InstanceRoleBackup {
			backups = append(backups, instance)
		} else {
			others = append(others, instance)
		}
	}

	return masters, slaves, backups, others
}

// 指定单实例mysql
func StartExecSingleMysqlSql(ctx *contexts.GlobalContext, priv *models.MysqlDBPriv, query string) ([]map[string]interface{}, []string, error) {
	// 获取执行的实例
	execInstance, err := getExecInstance(ctx, priv)
	if err != nil {
		return nil, nil, fmt.Errorf("获取可以执行实例出错. %v", err.Error())
	}

	mysqlCfg := config.NewMysqlConfig(execInstance.MachineHost.String, execInstance.Port.Int64, ctx.Cfg.ApiConfig.QueryMysqlUser, ctx.Cfg.ApiConfig.QueryMysqlPassword, priv.DBName.String)
	mysqlExecutor := executor.NewMysqlExcutor(ctx.Cfg.ExecConfig, mysqlCfg, query)

	return mysqlExecutor.Execute()
}
