package controllers

import (
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/views/response"
)

type MetaClusterController struct {
	ctx *contexts.GlobalContext
}

func NewMetaClusterController(ctx *contexts.GlobalContext) *MetaClusterController {
	return &MetaClusterController{ctx: ctx}
}

// 获取所有集群名
func (this *MetaClusterController) AllName() ([]*response.ClusterNameResponse, error) {
	clusters, err := dao.NewMetaClusterDao(this.ctx.EasydbDB).AllName()
	if err != nil {
		return nil, err
	}

	return helper.MetaClusterToNameResponses(clusters), nil
}

// 获取所有集群
func (this *MetaClusterController) All() ([]*models.MetaCluster, error) {
	return dao.NewMetaClusterDao(this.ctx.EasydbDB).All()
}
