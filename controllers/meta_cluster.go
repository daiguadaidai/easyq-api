package controllers

import (
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/types"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/daiguadaidai/easyq-api/views/response"
	"strings"
)

type MetaClusterController struct {
	ctx *contexts.GlobalContext
}

func NewMetaClusterController(ctx *contexts.GlobalContext) *MetaClusterController {
	return &MetaClusterController{ctx: ctx}
}

// 获取cluster
func (this *MetaClusterController) Find(req *request.MCFindRequest) ([]*response.ClusterResponse, error) {
	clusters, err := dao.NewMetaClusterDao(this.ctx.EasydbDB).FindByKeyword(strings.TrimSpace(req.Keyword.String), req.Offset(), req.Limit())
	if err != nil {
		return nil, err
	}

	return helper.MetaClusterToClusterResponses(clusters), nil
}

// 获取cluster count
func (this *MetaClusterController) Count(req *request.MCFindRequest) (int, error) {
	return dao.NewMetaClusterDao(this.ctx.EasydbDB).CountByKeyword(strings.TrimSpace(req.Keyword.String))
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
func (this *MetaClusterController) All() ([]*response.ClusterResponse, error) {
	clusters, err := dao.NewMetaClusterDao(this.ctx.EasydbDB).All()
	if err != nil {
		return nil, err
	}
	return helper.MetaClusterToClusterResponses(clusters), nil
}

// 添加
func (this *MetaClusterController) Add(req *request.MCAddRequest) error {
	var cluster models.MetaCluster
	utils.CopyStruct(req, &cluster)
	// 通过逗号合并域名
	cluster.DomainName = types.NewNullString(types.JoinStr(req.DomainNames, ","), false)
	cluster.VipPort = types.NewNullString(types.JoinStr(req.VipPorts, ","), false)
	cluster.VpcgwVipPort = types.NewNullString(types.JoinStr(req.VpcgwVipPorts, ","), false)

	return dao.NewMetaClusterDao(this.ctx.EasydbDB).Create(&cluster)
}

// 通过id编辑
func (this *MetaClusterController) EditById(req *request.MCEditByIdRequest) error {
	var cluster models.MetaCluster
	utils.CopyStruct(req, &cluster)
	// 通过逗号合并域名
	cluster.DomainName = types.NewNullString(types.JoinStr(req.DomainNames, ","), false)
	cluster.VipPort = types.NewNullString(types.JoinStr(req.VipPorts, ","), false)
	cluster.VpcgwVipPort = types.NewNullString(types.JoinStr(req.VpcgwVipPorts, ","), false)

	return dao.NewMetaClusterDao(this.ctx.EasydbDB).UpdateById(&cluster)
}

// 通过id删除
func (this *MetaClusterController) DeleteById(req *request.MCDeleteByIdRequest) error {
	var cluster models.MetaCluster
	utils.CopyStruct(req, &cluster)
	cluster.IsDeleted = types.NewNullInt64(1, false)

	return dao.NewMetaClusterDao(this.ctx.EasydbDB).UpdateById(&cluster)
}
