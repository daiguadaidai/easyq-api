package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/types"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/gin-gonic/gin"
)

type PrivsController struct {
	ctx *contexts.GlobalContext
}

func NewPrivsController(ctx *contexts.GlobalContext) *PrivsController {
	return &PrivsController{ctx: ctx}
}

func (this *PrivsController) ApplyMySQLPriv(c *gin.Context, req *request.PrivsApplyMysqlPrivRequest) error {
	// 获取用户信息
	clainms, err := middlewares.GetClaims(c)
	if err != nil {
		return fmt.Errorf("通过token获取用户信息出错: %v", err)
	}

	// 获取数据库用户信息
	user, err := dao.NewUserDao(this.ctx.EasyqDB).GetByUsername(clainms.Username)
	if err != nil {
		return fmt.Errorf("通过token解析出的用户获取数据库用户出错. username: %v. %v", clainms.Username, err)
	}

	// 生成申请单
	order := &models.MysqlDBPrivApplyOrder{
		UserId:      user.ID,
		Username:    user.Username,
		NameZh:      user.NameZh,
		ApplyStatus: types.NewNullInt64(models.MysqlDBPrivApplyOrderStatusApplying, false),
		ApplyReason: req.ApplyReason,
		OrderUUID:   types.NewNullString(utils.RandString(20), false),
	}

	// 生成申请权限
	applyPrivs := make([]*models.MysqlDBPrivApply, 0, len(req.Privs))
	for _, priv := range req.Privs {
		applyPrivs = append(applyPrivs, &models.MysqlDBPrivApply{
			OrderUUID:     order.OrderUUID,
			UserId:        user.ID,
			Username:      user.Username,
			NameZh:        user.NameZh,
			MetaClusterId: priv.MetaClusterId,
			ClusterName:   priv.ClusterName,
			DBName:        priv.DBName,
			VipPort:       priv.VipPort,
		})
	}

	// 权限去重复
	applyPrivs = helper.UniqueMysqlPrivApplys(applyPrivs)

	// 创建工单
	if err := dao.NewMysqlDBPrivApplyOrderDao(this.ctx.EasyqDB).Create(order); err != nil {
		return fmt.Errorf("创建申请单出错. %v. %v", utils.ToJsonStr(order), err)
	}

	// 批量创建申请权限
	if err := dao.NewMysqlDBPrivApplyDao(this.ctx.EasyqDB).BatchCreate(applyPrivs); err != nil {
		return fmt.Errorf("批量创建申请权限出错. %v. %v", utils.ToJsonStr(applyPrivs), err)
	}

	return nil
}

func (this *PrivsController) ApplyMysqlPrivSuccess(c *gin.Context, req *request.PrivsApplyPrivSuccessRequest) error {
	// 获取用户信息
	clainms, err := middlewares.GetClaims(c)
	if err != nil {
		return fmt.Errorf("通过token获取用户信息出错: %v", err)
	}

	// 获取数据库用户信息
	user, err := dao.NewUserDao(this.ctx.EasyqDB).GetByUsername(clainms.Username)
	if err != nil {
		return fmt.Errorf("通过token解析出的用户获取数据库用户出错. username: %v. %v", clainms.Username, err)
	}

	if user.Role.String != models.RoleDBA {
		return fmt.Errorf("不是DBA不允许审批")
	}

	// 获取单子
	order, err := dao.NewMysqlDBPrivApplyOrderDao(this.ctx.EasyqDB).GetByUUID(req.OrderUUID.String)
	if err != nil {
		return fmt.Errorf("通过uuid获取工单失败. uuid: %v. %v", req.OrderUUID.String, err)
	}
	if order == nil {
		return fmt.Errorf("申请单不存在, uuid: %v. %v", req.OrderUUID.String)
	}

	// 获取单子对应到申请权限
	applyPrivs, err := dao.NewMysqlDBPrivApplyDao(this.ctx.EasyqDB).FindByUUID(req.OrderUUID.String)
	if err != nil {
		return fmt.Errorf("通过uuid获取申请权限失败. uuid: %v. %v", req.OrderUUID.String, err)
	}
	if len(applyPrivs) == 0 {
		return fmt.Errorf("通过uuid没有获取到申请权限, uuid: %v. %v", req.OrderUUID.String)
	}

	// 权限申请权限转化成权限
	privs := models.PrivApplysToPrivs(applyPrivs)

	// 批量保存权限
	if err := dao.NewMysqlDBPrivDao(this.ctx.EasyqDB).BatchReplace(privs); err != nil {
		return fmt.Errorf("批量报错权限出错. %v, ", err, utils.ToJsonStr(privs))
	}

	return nil
}
func (this *PrivsController) ApplyMysqlPrivOrder(req *request.PrivsApplyMysqlPrivOrderRequest) ([]*models.MysqlDBPrivApplyOrder, error) {
	orderPram := &models.MysqlDBPrivApplyOrder{
		OrderUUID:   req.OrderUUID,
		Username:    req.Username,
		ApplyStatus: req.ApplyStatus,
	}

	return dao.NewMysqlDBPrivApplyOrderDao(this.ctx.EasyqDB).Find(orderPram, req.Offset(), req.Limit())
}
