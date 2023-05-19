package middlewares

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/gin-gonic/gin"
)

// 判断是否是DBA
func NeedRoleDBA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 查看是不是开发环境, 开发环境就不需要校验
		// 获取context
		globalCtx, err := GetGlobalContext(c)
		if err != nil {
			utils.ReturnError(c, utils.ResponseCodeErr, fmt.Errorf("[NeedRoleDBA] middlewares 获取全局context出错. %v", err.Error()))
			c.Abort()
			return
		}

		// 只有生产才需要校验DBA权限
		if globalCtx.Cfg.ApiConfig.Env == "prod" {
			// 检测用户权限是否是DBA, 非DBA不能编辑
			claims, err := GetClaims(c)
			if err != nil {
				utils.ReturnError(c, utils.ResponseCodeErr, fmt.Errorf("判断用户是否是DBA出错. %s", err.Error()))
				c.Abort()
				return
			}

			if claims.Role != models.RoleDBA {
				utils.ReturnError(c, utils.ResponseCodeErr, fmt.Errorf("用户不是DBA, 不允许操作"))
				c.Abort()
				return
			}
		}

		c.Next() // 处理请求
	}
}
