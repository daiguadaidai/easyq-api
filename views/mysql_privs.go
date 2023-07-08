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
	handler := new(MysqlPrivsHandler)
	AddHandlerV1("/mysql-privs", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type MysqlPrivsHandler struct{}

// 注册route
func (this *MysqlPrivsHandler) RegisterV1(group *gin.RouterGroup) {
	// 需要auth校验
	authGroup := group.Group("").Use(middlewares.JWTAuth())
	authGroup.POST("/apply-mysql-priv", this.ApplyMySQLPriv)
	authGroup.POST("/apply-mysql-priv-order", this.ApplyMysqlPrivOrder)
	authGroup.POST("/apply-mysql-priv-find-by-uuid", this.ApplyMysqlPrivFindByUUID)
	authGroup.POST("/find-privs-tree-by-username", this.FindPrivsTreeByUsername)
	authGroup.POST("/find-tables-by-user", this.FindTablesByUser)

	// 需要auth校验, 和DBA权限
	authAndDBAGroup := group.Group("").Use(middlewares.JWTAuth(), middlewares.NeedRoleDBA())
	authAndDBAGroup.POST("/apply-mysql-priv-success", this.ApplyMysqlPrivSuccess)
}

// 申请MySQL权限
func (this *MysqlPrivsHandler) ApplyMySQLPriv(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsApplyMysqlPrivRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMySQLPriv. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlPrivsHandler] ApplyMySQLPriv. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMySQLPriv. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMySQLPriv. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	if err := controllers.NewMysqlPrivsController(globalCtx).ApplyMySQLPriv(c, &req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMySQLPriv. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

// 申请MySQL权限
func (this *MysqlPrivsHandler) ApplyMysqlPrivSuccess(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsApplyPrivSuccessRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyPrivSuccess. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlPrivsHandler] ApplyPrivSuccess. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyPrivSuccess. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyPrivSuccess. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	if err := controllers.NewMysqlPrivsController(globalCtx).ApplyMysqlPrivSuccess(c, &req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyPrivSuccess. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

// 获取mysql工单列表
func (this *MysqlPrivsHandler) ApplyMysqlPrivOrder(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsApplyMysqlPrivOrderRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrder. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlPrivsHandler] ApplyMysqlPrivOrder. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrder. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	controller := controllers.NewMysqlPrivsController(globalCtx)
	// 获取列表
	orders, err := controller.ApplyMysqlPrivOrder(&req)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrder. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	// 获取数量
	cnt, err := controller.Count(&req)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrder. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, orders, int(cnt))
}

// 获取mysql工单列表
func (this *MysqlPrivsHandler) ApplyMysqlPrivFindByUUID(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsApplyMysqlPrivByUUIDRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivFindByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlPrivsHandler] ApplyMysqlPrivFindByUUID. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivFindByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivFindByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取列表
	applyPrivs, err := controllers.NewMysqlPrivsController(globalCtx).ApplyPrivsFindByUUID(&req)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivFindByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, applyPrivs, len(applyPrivs))
}

// 申请MySQL权限
func (this *MysqlPrivsHandler) ApplyMysqlPrivOrderEditByUUID(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsApplyMysqlPrivOrderEditByUUIDRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrderEditByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlPrivsHandler] ApplyMysqlPrivOrderEditByUUID. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrderEditByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrderEditByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	if err := controllers.NewMysqlPrivsController(globalCtx).ApplyMysqlPrivOrderEditByUUID(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] ApplyMysqlPrivOrderEditByUUID. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

// 通过用户获取所有的数据库权限
func (this *MysqlPrivsHandler) FindPrivsTreeByUsername(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsMysqlFindTreeByUsername
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindPrivsTreeByUsername. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlPrivsHandler] FindPrivsTreeByUsername. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindPrivsTreeByUsername. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindPrivsTreeByUsername. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	privs, err := controllers.NewMysqlPrivsController(globalCtx).FindPrivsTreeByUsername(&req)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindPrivsTreeByUsername. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, privs, len(privs))
}

// 通过用户获取所有的数据库权限
func (this *MysqlPrivsHandler) FindTablesByUser(c *gin.Context) {
	// 解析 request参数
	var req request.PrivsMysqlFindTablesByUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindTablesByUser. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[MysqlPrivsHandler] FindTablesByUser. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindTablesByUser. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindTablesByUser. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 创建申请工单
	tableNames, err := controllers.NewMysqlPrivsController(globalCtx).FindTablesByUser(c, &req)
	if err != nil {
		logger.M.Errorf("[MysqlPrivsHandler] FindTablesByUser. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, tableNames, len(tableNames))
}
