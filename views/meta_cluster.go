package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/daiguadaidai/easyq-api/controllers"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
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
	authGroup.POST("/find", this.Find)
	authGroup.GET("/all-name", this.AllName)
	authGroup.GET("/all", this.All)

	// 需要auth校验, 和DBA权限
	authAndDBAGroup := group.Group("").Use(middlewares.JWTAuth(), middlewares.NeedRoleDBA())
	authAndDBAGroup.POST("/add", this.Add)
	authAndDBAGroup.POST("/edit-by-id", this.EditById)
	authAndDBAGroup.POST("/delete-by-id", this.DeleteById)
}

func (this *ClusterHandler) Find(c *gin.Context) {
	// 解析 request参数
	var req request.MCFindRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[ClusterHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[ClusterHandler] Find. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	controller := controllers.NewMetaClusterController(globalCtx)
	// 获取集群列表
	clusters, err := controller.Find(&req)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	cnt, err := controller.Count(&req)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, clusters, cnt)
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

func (this *ClusterHandler) Add(c *gin.Context) {
	// 解析 request参数
	var req request.MCAddRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[ClusterHandler] Add. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[ClusterHandler] Add. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] Add. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewMetaClusterController(globalCtx).Add(&req); err != nil {
		err = fmt.Errorf("创建集群失败. %v", err.Error())
		logger.M.Errorf("[ClusterHandler] Add. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

func (this *ClusterHandler) EditById(c *gin.Context) {
	// 解析 request参数
	var req request.MCEditByIdRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[ClusterHandler] EditById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[ClusterHandler] EditById. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] EditById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewMetaClusterController(globalCtx).EditById(&req); err != nil {
		err = fmt.Errorf("通过Id修改集群信息出错. %v", err.Error())
		logger.M.Errorf("[ClusterHandler] EditById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

func (this *ClusterHandler) DeleteById(c *gin.Context) {
	// 解析 request参数
	var req request.MCDeleteByIdRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[ClusterHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[ClusterHandler] DeleteById. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[ClusterHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewMetaClusterController(globalCtx).DeleteById(&req); err != nil {
		err = fmt.Errorf("通过Id删除集群信息出错. %v", err.Error())
		logger.M.Errorf("[ClusterHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}
