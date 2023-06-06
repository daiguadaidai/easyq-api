package controllers

import (
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/models/view_models"
	"github.com/daiguadaidai/easyq-api/views/request"
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

func (this *InstanceController) Count(req *request.InstanceFindRequest) (int, error) {
	return dao.NewInstanceDao(this.ctx.EasydbDB).CountByKeyword(strings.TrimSpace(req.Keyword.String))
}
