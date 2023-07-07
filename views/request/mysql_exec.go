package request

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/types"
	"strings"
)

type MysqlExecSqlRequest struct {
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" form:"meta_cluster_id"`
	DBName        types.NullString `json:"db_name" form:"db_name"`
	Query         types.NullString `json:"query" form:"query"`
}

func (this *MysqlExecSqlRequest) Check() error {
	if this.MetaClusterId.IsZero() {
		return fmt.Errorf("集群信息不能为空")
	}
	if this.DBName.IsEmpty() {
		return fmt.Errorf("数据库名不能为空")
	}
	if strings.TrimSpace(this.Query.String) == "" {
		return fmt.Errorf("sql不能为空")

	}

	return nil
}
