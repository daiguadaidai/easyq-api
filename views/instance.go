package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/daiguadaidai/easyq-api/controllers"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/daiguadaidai/easyq-api/views/response"
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
	authGroup.POST("/find-master-slaves", this.FindMasterSlaves)

	// 需要auth校验, 和DBA权限
	authAndDBAGroup := group.Group("").Use(middlewares.JWTAuth(), middlewares.NeedRoleDBA())
	authAndDBAGroup.POST("/add", this.Add)
	authAndDBAGroup.POST("/edit-by-id", this.EditById)
	authAndDBAGroup.POST("/delete-by-id", this.DeleteById)
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

func (this *InstanceHandler) FindMasterSlaves(c *gin.Context) {
	// 解析 request参数
	var req request.InstanceFindMasterSlavesRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[InstanceHandler] FindMasterSlaves. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[InstanceHandler] FindMasterSlaves. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[InstanceHandler] FindMasterSlaves. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	var cnt int
	var resps []*response.InstanceClusterResponse
	controller := controllers.NewInstanceController(globalCtx)
	if !req.MetaClusterId.IsZero() { // 通过clusterId 获取所有实例
		resps, err = controller.FindInstanceClusterByClusterId(req.MetaClusterId.Int64)
		if err != nil {
			err = fmt.Errorf("集群id:%v, %s", req.MetaClusterId.Int64, err.Error())
			logger.M.Errorf("[InstanceHandler] FindMasterSlaves. %v", err.Error())
			utils.ReturnError(c, utils.ResponseCodeErr, err)
			return
		}
		cnt = len(resps)
	} else if !req.ClusterName.IsEmpty() { // 有指定集群名称
		resps, err = controller.FindInstanceClusterByClusterName(req.ClusterName.String)
		if err != nil {
			err = fmt.Errorf("集群名称:%v, %s", req.ClusterName.String, err.Error())
			logger.M.Errorf("[InstanceHandler] FindMasterSlaves. %v", err.Error())
			utils.ReturnError(c, utils.ResponseCodeErr, err)
			return
		}
		cnt = len(resps)
	} else if !req.SetName.IsEmpty() { // 指定setname
		resps, err = controller.FindInstanceClusterBySetName(req.SetName.String)
		if err != nil {
			err = fmt.Errorf("集群Set名称:%v, %s", req.SetName.String, err.Error())
			logger.M.Errorf("[InstanceHandler] FindMasterSlaves. %v", err.Error())
			utils.ReturnError(c, utils.ResponseCodeErr, err)
			return
		}
		cnt = len(resps)
	}

	utils.ReturnList(c, resps, cnt)
}

func (this *InstanceHandler) Add(c *gin.Context) {
	// 解析 request参数
	var req request.InstanceAddRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[InstanceHandler] Add. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[InstanceHandler] Add. req: %s", utils.ToJsonStr(req))
	if err := req.Check(); err != nil {
		logger.M.Errorf("[InstanceHandler] Add. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[InstanceHandler] Add. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewInstanceController(globalCtx).Add(&req); err != nil {
		err = fmt.Errorf("创建实例失败. %v", err.Error())
		logger.M.Errorf("[InstanceHandler] Add. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

func (this *InstanceHandler) EditById(c *gin.Context) {
	// 解析 request参数
	var req request.InstanceEditByIdRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[InstanceHandler] EditById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[InstanceHandler] EditById. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[InstanceHandler] EditById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewInstanceController(globalCtx).EditById(&req); err != nil {
		err = fmt.Errorf("通过Id修改实例信息出错. %v", err.Error())
		logger.M.Errorf("[InstanceHandler] EditById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

func (this *InstanceHandler) DeleteById(c *gin.Context) {
	// 解析 request参数
	var req request.InstanceDeleteByIdRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[InstanceHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[InstanceHandler] DeleteById. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[InstanceHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewInstanceController(globalCtx).DeleteById(&req); err != nil {
		err = fmt.Errorf("通过Id删除实例信息出错. %v", err.Error())
		logger.M.Errorf("[InstanceHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}
