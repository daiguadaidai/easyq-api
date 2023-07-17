package views

import (
	"github.com/daiguadaidai/easyq-api/controllers"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/gin-gonic/gin"
)

func init() {
	handler := new(MysqlExecHandler)
	AddHandlerV1("/mysql-exec", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type MysqlExecHandler struct{}

// 注册route
func (this *MysqlExecHandler) RegisterV1(group *gin.RouterGroup) {
	// 需要auth校验
	authGroup := group.Group("").Use(middlewares.JWTAuth())
	authGroup.POST("/exec-sql", this.ExecSql)
}

// 执行SQL
func (this *MysqlExecHandler) ExecSql(c *gin.Context) {
	// 解析 request参数
	var req request.MysqlExecSqlRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlExecHandler] ExecSql. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlExecHandler] ExecSql. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[MysqlExecHandler] ExecSql. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlExecHandler] ExecSql. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	rs, err := controllers.NewMysqlExecController(globalCtx).ExecSql(c, &req)
	if err != nil {
		logger.M.Errorf("[MysqlExecHandler] ExecSql. %v", err.Error())
		rs.IsErr = true
		rs.ErrMsg = err.Error()
		return
	}

	utils.ReturnSuccess(c, rs)
}
