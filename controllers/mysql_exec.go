package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/middlewares"
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

	// 通过集群id和数据库名和用户获取数据库权限, 判断用户是否有权限
	priv, err := dao.NewMysqlDBPrivDao(this.ctx.EasyqDB).GetByUsernameClusterDB(user.Username.String, req.MetaClusterId.Int64, req.DBName.String)
	if err != nil {
		return fmt.Errorf("获取数据库权限出错. username: %v, meta_cluster_id: %v, db_name: %v, %v", user.Username.String, req.MetaClusterId.Int64, req.DBName.String, err.Error())
	}

	logger.M.Infof(utils.ToJsonStr(priv))

	return nil
}
