package middlewares

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/gin-gonic/gin"
)

const (
	TokenKey  = "X-Token"
	ClaimsKey = "claims"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(TokenKey)
		if token == "" {
			utils.ReturnError(c, utils.ResponseCodeTokenErr, fmt.Errorf("请求未携带token，无权限访问"))
			c.Abort()
			return
		}

		// 获取token用户
		claims, err := utils.TokenToClaims(token)
		if err != nil {
			utils.ReturnError(c, utils.ResponseCodeTokenErr, fmt.Errorf("%s. %s", err.Error(), token))
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set(ClaimsKey, claims)
		c.Next()
	}
}

// 获取用户信息
func GetClaims(ctx *gin.Context) (*utils.CustomClaims, error) {
	v, exists := ctx.Get(ClaimsKey)
	if !exists {
		return nil, fmt.Errorf("Claims 没有获取到用户信息")
	}
	vv, ok := v.(*utils.CustomClaims)
	if !ok {
		return nil, fmt.Errorf("用户信息转化为 CustomClaims 失败")
	}
	return vv, nil
}
