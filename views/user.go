package views

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/controllers"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/gin-gonic/gin"
)

func init() {
	handler := new(UserHandler)
	AddHandlerV1("/user", handler) // 添加当前页面的uri路径之前都要添加上这个
}

type UserHandler struct{}

// 注册route
func (this *UserHandler) RegisterV1(group *gin.RouterGroup) {
	noAuthGroup := group.Group("")
	noAuthGroup.POST("/register", this.Register)
	noAuthGroup.POST("/login", this.Login)
	noAuthGroup.POST("/logout", this.Logout)
	noAuthGroup.GET("/all", this.All)

	authGroup := group.Group("").Use(middlewares.JWTAuth(), middlewares.NeedRoleDBA())
	authGroup.POST("/find", this.Find)
	authGroup.POST("/edit-by-id", this.EditById)
	authGroup.POST("/delete-by-id", this.DeleteById)
}

func (this *UserHandler) Register(c *gin.Context) {
	// 解析 request参数
	var req request.UserRegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[UserHandler] Register. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[UserHandler] Register. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[UserHandler] Register. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 注册
	user, err := controllers.NewUserController(globalCtx).Register(&req)
	if err != nil {
		logger.M.Errorf("[UserHandler] Register. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, user)
}

func (this *UserHandler) Login(c *gin.Context) {
	// 解析 request参数
	var req request.UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[UserHandler] Login. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[UserHandler] Login. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[UserHandler] Login. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	resp, err := controllers.NewUserController(globalCtx).Login(&req)
	if err != nil {
		logger.M.Errorf("[UserHandler] Login. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, resp)
}

func (this *UserHandler) Logout(c *gin.Context) {
	utils.ReturnSuccess(c, nil)
}

func (this *UserHandler) Find(c *gin.Context) {
	// 解析 request参数
	var req request.UserFindRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[UserHandler] FindUsers. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[UserHandler] FindUsers. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[UserHandler] FindUsers. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	controller := controllers.NewUserController(globalCtx)
	// 获取集群列表
	users, err := controller.Find(&req)
	if err != nil {
		logger.M.Errorf("[UserHandler] FindUsers. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取总共集群数
	cnt, err := controller.Count(&req)
	if err != nil {
		logger.M.Errorf("[UserHandler] FindUsers. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, users, cnt)
}

func (this *UserHandler) EditById(c *gin.Context) {
	// 解析 request参数
	var req request.UserEditByIdRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[UserHandler] EditById. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[UserHandler] EditById. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[UserHandler] EditById. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewUserController(globalCtx).EditById(&req); err != nil {
		err = fmt.Errorf("通过Id修改用户信息出错. %v", err.Error())
		logger.M.Errorf("[UserHandler] EditById. %s", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

func (this *UserHandler) DeleteById(c *gin.Context) {
	// 解析 request参数
	var req request.UserDeleteByIdRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[UserHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	logger.M.Infof("[UserHandler] DeleteById. req: %s", utils.ToJsonStr(req))

	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[UserHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	if err := controllers.NewUserController(globalCtx).DeleteById(&req); err != nil {
		err = fmt.Errorf("通过Id删除用户信息出错. %v", err.Error())
		logger.M.Errorf("[UserHandler] DeleteById. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, nil)
}

func (this *UserHandler) All(c *gin.Context) {
	// 获取context
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Errorf("[UserHandler] All. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	// 获取集群列表
	users, err := controllers.NewUserController(globalCtx).All()
	if err != nil {
		err = fmt.Errorf("获取所有用户失败 %v", err.Error())
		logger.M.Errorf("[UserHandler] All. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, users, len(users))
}
