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
	handler := new(PrivsHandler)
	AddHandlerV1("/privs", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type PrivsHandler struct{}

// 注册route
func (this *PrivsHandler) RegisterV1(group *gin.RouterGroup) {
	// 需要auth校验
	authGroup := group.Group("").Use(middlewares.JWTAuth())
	authGroup.POST("/apply-mysql-priv", this.ApplyMySQLPriv)

	// 需要auth校验, 和DBA权限
	authAndDBAGroup := group.Group("").Use(middlewares.JWTAuth(), middlewares.NeedRoleDBA())
	authAndDBAGroup.POST("/apply-mysql-priv-success", this.ApplyMysqlPrivSuccess)
}

// 申请MySQL权限
func (this *PrivsHandler) ApplyMySQLPriv(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsApplyMysqlPrivRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[PrivsHandler] ApplyMySQLPriv. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[PrivsHandler] ApplyMySQLPriv. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[PrivsHandler] ApplyMySQLPriv. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	if err := controllers.NewPrivsController(globalCtx).ApplyMySQLPriv(c, req); err != nil {
		logger.M.Errorf("[PrivsHandler] ApplyMySQLPriv. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

// 申请MySQL权限
func (this *PrivsHandler) ApplyMysqlPrivSuccess(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsApplyPrivSuccessRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[PrivsHandler] ApplyPrivSuccess. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[PrivsHandler] ApplyPrivSuccess. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[PrivsHandler] ApplyPrivSuccess. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	if err := controllers.NewPrivsController(globalCtx).ApplyMysqlPrivSuccess(c, req); err != nil {
		logger.M.Errorf("[PrivsHandler] ApplyPrivSuccess. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}
