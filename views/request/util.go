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

type UtilGetBatchInsertSqlRequest struct {
	TableName   string                   `json:"table_name" form:"table_name"`
	ColumnNames []string                 `json:"column_names" form:"column_names"`
	Rows        []map[string]interface{} `json:"rows" form:"rows"`
}

func (this *UtilGetBatchInsertSqlRequest) Check() error {
	if len(this.ColumnNames) == 0 {
		return fmt.Errorf("字段名不能为空")
	}

	if len(this.Rows) == 0 {
		return fmt.Errorf("数据不能为空")
	}

	if strings.TrimSpace(this.TableName) == "" {
		return fmt.Errorf("表名不能为空")
	}

	return nil
}
