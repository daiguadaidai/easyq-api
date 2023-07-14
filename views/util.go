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
	handler := new(UtilHandler)
	AddHandlerV1("/utils", handler) // 添加当前页面的uri路径之前都要添加上这个
}

// 注册route
func (this *UtilHandler) RegisterV1(group *gin.RouterGroup) {
	noAuthGroup := group.Group("")
	noAuthGroup.GET("/encrypt", this.Encrypt)
	noAuthGroup.GET("/decrypt", this.Decrypt)
	noAuthGroup.POST("/sql-fingerprint", this.SqlFingerprint)
	noAuthGroup.POST("/db-result", this.DBResult)
	noAuthGroup.POST("/text-to-sqls", this.TextToSqls)

	authGroup := group.Group("").Use(middlewares.JWTAuth())
	authGroup.GET("/jwt-auth-test", this.JWTAuthTest)
}

type UtilHandler struct{}

// 加密
func (this *UtilHandler) Encrypt(c *gin.Context) {
	var req request.UtilEncreptRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	data, err := controllers.NewUtilController(globalCtx).Encrypt(&req)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
	}
	utils.ReturnSuccess(c, data)
}

// 解密
func (this *UtilHandler) Decrypt(c *gin.Context) {
	var req request.UtilDecryptRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	data, err := controllers.NewUtilController(globalCtx).Decrypt(&req)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
	}
	utils.ReturnSuccess(c, data)
}

// 测试Token
func (this *UtilHandler) JWTAuthTest(c *gin.Context) {
	utils.ReturnSuccess(c, nil)
}

// 解密
func (this *UtilHandler) SqlFingerprint(c *gin.Context) {
	var req request.UtilSqlFingerprintRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	fingers, err := controllers.NewDefaultUtilController(globalCtx).GetSqlFingerprints(&req)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	utils.ReturnList(c, fingers, len(fingers))
}

func (this *UtilHandler) DBResult(c *gin.Context) {
	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	result, err := controllers.NewUtilController(globalCtx).DBResult()
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnSuccess(c, result)
}

func (this *UtilHandler) TextToSqls(c *gin.Context) {
	var req request.UtilTextToSqlsRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.M.Errorf("[UtilHandler] TextToSqls. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}
	if err := req.Check(); err != nil {
		logger.M.Errorf("[UtilHandler] TextToSqls. %v", err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	globalCtx, err := middlewares.GetGlobalContext(c)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	sqls, err := controllers.NewUtilController(globalCtx).TextToSqls(&req)
	if err != nil {
		logger.M.Error(err.Error())
		utils.ReturnError(c, utils.ResponseCodeErr, err)
		return
	}

	utils.ReturnList(c, sqls, len(sqls))
}
