package controllers

import (
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/models/view_models"
	"github.com/daiguadaidai/easyq-api/types"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/daiguadaidai/easyq-api/views/response"
	"strings"
)

type InstanceController struct {
	ctx *contexts.GlobalContext
}

func NewInstanceController(ctx *contexts.GlobalContext) *InstanceController {
	return &InstanceController{ctx: ctx}
}

// 获取 instance
func (this *InstanceController) Find(req *request.InstanceFindRequest) ([]*view_models.InstanceCluster, error) {
	instances, err := dao.NewInstanceDao(this.ctx.EasydbDB).FindByKeyword(strings.TrimSpace(req.Keyword.String), req.Offset(), req.Limit())
	if err != nil {
		return nil, err
	}

	return instances, nil
}

// 获取 instance count
func (this *InstanceController) Count(req *request.InstanceFindRequest) (int, error) {
	return dao.NewInstanceDao(this.ctx.EasydbDB).CountByKeyword(strings.TrimSpace(req.Keyword.String))
}

// 通过集群id获取实例集群信息
func (this *InstanceController) FindInstanceClusterByClusterId(clusterId int64) ([]*response.InstanceClusterResponse, error) {
	instances, err := dao.NewInstanceDao(this.ctx.EasydbDB).FindInstanceClusterByClusterId(clusterId)
	if err != nil {
		return nil, err
	}

	instanceResps := helper.InstanceClustersToMasterSlavesResps(instances)

	return instanceResps, nil
}

// 通过集群名称获取实例集群信息
func (this *InstanceController) FindInstanceClusterByClusterName(clusterName string) ([]*response.InstanceClusterResponse, error) {
	instances, err := dao.NewInstanceDao(this.ctx.EasydbDB).FindInstanceClusterByClusterName(clusterName)
	if err != nil {
		return nil, err
	}

	instanceResps := helper.InstanceClustersToMasterSlavesResps(instances)

	return instanceResps, nil
}

// 通过集群set获取实例集群信息
func (this *InstanceController) FindInstanceClusterBySetName(setName string) ([]*response.InstanceClusterResponse, error) {
	instances, err := dao.NewInstanceDao(this.ctx.EasydbDB).FindInstanceClusterBySetName(setName)
	if err != nil {
		return nil, err
	}

	instanceResps := helper.InstanceClustersToMasterSlavesResps(instances)

	return instanceResps, nil
}

// 添加
func (this *InstanceController) Add(req *request.InstanceAddRequest) error {
	var instance models.Instance
	utils.CopyStruct(req, &instance)

	return dao.NewInstanceDao(this.ctx.EasydbDB).Create(&instance)
}

// 通过id编辑
func (this *InstanceController) EditById(req *request.InstanceEditByIdRequest) error {
	var insdtance models.Instance
	utils.CopyStruct(req, &insdtance)

	return dao.NewInstanceDao(this.ctx.EasydbDB).UpdateById(&insdtance)
}

// 通过id删除
func (this *InstanceController) DeleteById(req *request.InstanceDeleteByIdRequest) error {
	var instance models.Instance
	utils.CopyStruct(req, &instance)
	instance.IsDeleted = types.NewNullInt64(1, false)

	return dao.NewInstanceDao(this.ctx.EasydbDB).UpdateById(&instance)
}
