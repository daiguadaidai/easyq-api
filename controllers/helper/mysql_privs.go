package helper

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/models"
)

// 申请权限去重
func UniqueMysqlPrivApplys(applys []*models.MysqlDBPrivApply) []*models.MysqlDBPrivApply {
	privMap := make(map[string]*models.MysqlDBPrivApply)
	uniqueApplys := make([]*models.MysqlDBPrivApply, 0, len(applys))

	for _, apply := range applys {
		key := fmt.Sprintf("%v#%v#v", apply.MetaClusterId.Int64, apply.DBName.String)
		if _, ok := privMap[key]; !ok {
			privMap[key] = apply
			uniqueApplys = append(uniqueApplys, apply)
		}
	}

	return uniqueApplys
}
