package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/gin-gonic/gin"
)

type MysqlExecController struct {
	ctx *contexts.GlobalContext
}

func NewMysqlExecController(ctx *contexts.GlobalContext) *MysqlExecController {
	return &MysqlExecController{ctx: ctx}
}

func (this *MysqlExecController) ExecSql(c *gin.Context, req *request.MysqlExecSqlRequest) error {
	// 检测权限是否存在
	priv, err := helper.CheckMysqlPriv(c, this.ctx, req.DBName.String, req.MetaClusterId.Int64)
	if err != nil {
		return fmt.Errorf("检测用户数据库权限出错. %v", err.Error())
	}

	logger.M.Infof(utils.ToJsonStr(priv))

	return nil
}
