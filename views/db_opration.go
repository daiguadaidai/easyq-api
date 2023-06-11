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
	handler := new(DBOprationHandler)
	AddHandlerV1("/db-operation", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type DBOprationHandler struct{}

// 注册route
func (this *DBOprationHandler) RegisterV1(group *gin.RouterGroup) {
	// 需要auth校验
	authGroup := group.Group("").Use(middlewares.JWTAuth())
	authGroup.POST("/db-names", this.ClusterDBNames)
}

// 获取集群所有数据库名
func (this *DBOprationHandler) ClusterDBNames(c *gin.Context) {
	// 解析 request参数
	var req request.DBOperationClusterDBNamesRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[DBOprationHandler] ClusterDBNames. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[DBOprationHandler] ClusterDBNames. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[DBOprationHandler] ClusterDBNames. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[DBOprationHandler] ClusterDBNames. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	controller := controllers.NewDBOperationController(globalCtx)
	// 获取集群列表
	dbNames, err := controller.FindClusterDBNames(&req)
	if err != nil {
		logger.M.Errorf("[DBOprationHandler] ClusterDBNames. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, dbNames, len(dbNames))
}
