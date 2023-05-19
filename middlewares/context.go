package middlewares

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/gin-gonic/gin"
)

// GlobalContextKey is the key in gin.Context for GlobalContext.
const GlobalContextKey = "global-contexts"

// WithGlobalContext injects the GlobalContext instance into gin.Context.
func WithGlobalContext(globalCtx *contexts.GlobalContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(GlobalContextKey, globalCtx)
		ctx.Next()
	}
}

// GetGlobalContext fetches the GlobalContext instance from middlewares.
func GetGlobalContext(ctx *gin.Context) (*contexts.GlobalContext, error) {
	v, exists := ctx.Get(GlobalContextKey)
	if !exists {
		return nil, fmt.Errorf("在gin.context自定义的GlobalContext不存在")
	}
	vv, ok := v.(*contexts.GlobalContext)
	if !ok {
		return nil, fmt.Errorf("gin.context中强制转化GlobaoContext失败: %v", v)
	}
	return vv, nil
}
