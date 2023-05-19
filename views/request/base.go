package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/daiguadaidai/easyq-api/logger"
	"strings"
)

const (
	DefaultPageSize = 50
	MaxPageSize     = 500
)

func NewRequest(c *gin.Context, obj interface{}) (interface{}, error) {
	if err := c.ShouldBind(obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// 获取参数
func GetParam(c *gin.Context, key string) (string, error) {
	v := c.Param(key)
	if strings.TrimSpace(v) == "" {
		return "", fmt.Errorf("必须输入参数 %s 值")
	}
	return v, nil
}

// 获取参数
func GetParamInt64(c *gin.Context, key string) (int64, error) {
	v, err := GetParam(c, key)
	if err != nil {
		return 0, err
	}

	i, err := com.StrTo(v).Int64()
	if err != nil {
		return 0, err
	}
	return i, nil
}

type Pager struct {
	Current  int `json:"current" form:"current"`   // 第一页从1开始
	PageSize int `json:"pageSize" form:"pageSize"` // 每业大小
}

func (this *Pager) Check() error {
	if this.PageSize > MaxPageSize {
		return fmt.Errorf("每页最大不能超过 %d", MaxPageSize)
	}

	return nil
}

func (this *Pager) Offset() int {
	offset := (this.Current - 1) * this.Limit()
	if offset < 0 {
		logger.M.Warnf("[Pager] Offset < 0 自动设置成0, Current: %v, PageSize: %v, offset: %v, limit: %v",
			this.Current, this.Limit(), 0, this.Limit())
		return 0
	}

	return offset
}

func (this *Pager) Limit() int {
	if this.PageSize <= 0 {
		logger.M.Warnf("[Pager] Limit 使用默认 PageSize: %v", DefaultPageSize)
		return DefaultPageSize
	}

	return this.PageSize
}
