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
	handler := new(InstanceHandler)
	AddHandlerV1("/instance", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type InstanceHandler struct{}

// 注册route
func (this *InstanceHandler) RegisterV1(group *gin.RouterGroup) {
	// 需要auth校验
	authGroup := group.Group("").Use(middlewares.JWTAuth())
	authGroup.POST("/find", this.Find)
}

func (this *InstanceHandler) Find(c *gin.Context) {
	// 解析 request参数
	var req request.InstanceFindRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[InstanceHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[InstanceHandler] Find. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[InstanceHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	controller := controllers.NewInstanceController(globalCtx)
	// 获取集群列表
	instances, err := controller.Find(&req)
	if err != nil {
		logger.M.Errorf("[InstanceHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取总共集群数
	cnt, err := controller.Count(&req)
	if err != nil {
		logger.M.Errorf("[InstanceHandler] Find. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, instances, cnt)
}