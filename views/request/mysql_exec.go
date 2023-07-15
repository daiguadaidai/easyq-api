package request

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/types"
	"strings"
)

type MysqlExecSqlRequest struct {
	PrivId types.NullInt64  `json:"priv_id" form:"priv_id"`
	Query  types.NullString `json:"query" form:"query"`
}

func (this *MysqlExecSqlRequest) Check() error {
	if this.PrivId.IsZero() {
		return fmt.Errorf("集群信息不能为空")
	}
	if strings.TrimSpace(this.Query.String) == "" {
		return fmt.Errorf("sql不能为空")

	}

	return nil
}
