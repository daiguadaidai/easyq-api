package services

import (
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/views"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// 启动http服务
func StartApiService(wg *sync.WaitGroup, ctx *contexts.GlobalContext) {
	defer wg.Done()

	// 注册路由
	router := gin.New()
	router.Use(middlewares.ZapLogger(), middlewares.GinRecovery(true))
	router.Use(middlewares.WithGlobalContext(ctx))
	router.Use(middlewares.Cors())
	views.Register(router)

	// 获取pala启动配置信息
	s := &http.Server{
		Addr:           ctx.Cfg.ApiConfig.Address(),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.M.Infof("EasyQ Api 监听地址为: %v", ctx.Cfg.ApiConfig.Address())
	err := s.ListenAndServe()
	if err != nil {
		logger.M.Errorf("EasyQ Api 启动服务出错: %v", err)
	}
}
