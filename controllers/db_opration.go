package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/views/request"
)

type DBOperationController struct {
	ctx *contexts.GlobalContext
}

func NewDBOperationController(ctx *contexts.GlobalContext) *DBOperationController {
	return &DBOperationController{ctx: ctx}
}

func (this *DBOperationController) FindClusterDBNames(req *request.DBOperationClusterDBNamesRequest) ([]string, error) {
	// 获取集群
	metaCluster, err := dao.NewMetaClusterDao(this.ctx.EasydbDB).GetById(req.MetaClusterId.Int64)
	if err != nil {
		return nil, fmt.Errorf("通过id获取集群失败. id: %v, %v", req.MetaClusterId.Int64, err)
	}
	if metaCluster == nil {
		return nil, fmt.Errorf("集群不存在. id: %v", req.MetaClusterId.Int64)
	}

	dbNames := make([]string, 0, 10)
	// tdsql 分库分表
	if metaCluster.Category.Int64 == models.ClusterCategoryTceTDSQL &&
		metaCluster.IsShard.Int64 == models.IsShardYes &&
		metaCluster.ShardType.String == models.ShardTypeTDSQL {
		// 获取数据库名
		dbNames, err = helper.FindDBNameByVipPort(this.ctx, metaCluster.VipPort.String)
		if err != nil {
			return nil, fmt.Errorf("分布式tdsql获取数据库名出错. %v", err)
		}
	} else {
		// 获取 单实例master
		master, err := dao.NewInstanceDao(this.ctx.EasydbDB).GetMasterByMetaClusterId(req.MetaClusterId.Int64)
		if err != nil {
			return nil, fmt.Errorf("获取集群master出错. metaClusterId: %v. %v", req.MetaClusterId.Int64, err)
		}

		// 获取数据库名
		dbNames, err = helper.FindDBNameByHostPort(this.ctx, master.MachineHost.String, master.Port.Int64)
		if err != nil {
			return nil, fmt.Errorf("通过 Master 获取数据库名出错. metaClusterId: %v. vip:port: %v. %v", req.MetaClusterId.Int64, master.MachineHost.String, master.Port.Int64, err)
		}
	}

	return helper.FilterIgnoreDatabases(dbNames), nil
}
