package views

import (
	"github.com/daiguadaidai/easyq-api/controllers"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	handler := new(ClusterHandler)
	AddHandlerV1("/cluster", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type ClusterHandler struct{}

// 注册route
func (this *ClusterHandler) RegisterV1(group *gin.RouterGroup) {
	// 需要auth校验
	authGroup := group.Group("").Use(middlewares.JWTAuth())
	authGroup.GET("/all-name", this.AllName)
	authGroup.GET("/all", this.All)
}

// 通过集群名称搜索集群
func (this *ClusterHandler) AllName(c *gin.Context) {
	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] AllName. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	controller := controllers.NewMetaClusterController(globalCtx)
	// 获取集群列表
	clusters, err := controller.AllName()
	if err != nil {
		logger.M.Errorf("[ClusterHandler] AllName. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, clusters, len(clusters))
}

// 所有集群
func (this *ClusterHandler) All(c *gin.Context) {
	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] All. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	controller := controllers.NewMetaClusterController(globalCtx)
	// 获取集群列表
	clusters, err := controller.All()
	if err != nil {
		logger.M.Errorf("[ClusterHandler] All. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, clusters, len(clusters))
}
