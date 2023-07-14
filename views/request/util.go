package request

import (
	"fmt"
	"strings"
)

type UtilEncreptRequest struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type UtilDecryptRequest struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type UtilSqlFingerprintRequest struct {
	Statements []string `json:"statements" form:"statements" binding:"required"`
}

type UtilTextToSqlsRequest struct {
	Text string `json:"text" form:"text" binding:"required"`
}

func (this *UtilTextToSqlsRequest) Check() error {
	if strings.TrimSpace(this.Text) == "" {
		return fmt.Errorf("请输入需要拆分多sql文本")
	}

	return nil
}
