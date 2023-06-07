package request

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/types"
)

var ignoreDatabases = map[string]struct{}{
	"information_schema": {},
	"mysql":              {},
	"performance_schema": {},
	"sys":                {},
	"test":               {},
	"sysdb":              {},
}

// 过滤掉可以忽略掉数据库
func FilterIgnoreDatabases(databases []string) []string {
	newDatabases := make([]string, 0, len(databases))
	for _, database := range databases {
		if _, ok := ignoreDatabases[database]; !ok {
			newDatabases = append(newDatabases, database)
		}
	}

	return newDatabases
}

type DBOperationClusterDBNamesRequest struct {
	MetaClusterId types.NullInt64 `json:"meta_cluster_id" form:"meta_cluster_id"`
}

func (this *DBOperationClusterDBNamesRequest) Check() error {
	if this.MetaClusterId.IsZero() {
		return fmt.Errorf("集群名不能为空")
	}

	return nil
}
