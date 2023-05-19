package view_models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

type GrantInfo struct {
	InstanceId    types.NullInt64  `json:"instance_id" gorm:"column:instance_id"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" gorm:"column:meta_cluster_id"`
	DBHost        types.NullString `json:"db_host" gorm:"column:db_host"`
	DBPort        types.NullInt64  `json:"db_port" gorm:"column:db_port"`
	GrantUser     types.NullString `json:"grant_user" gorm:"column:grant_user"`
	GrantHost     types.NullString `json:"grant_host" gorm:"column:grant_host"`
	GrantInfos    []string         `json:"grant_infos" form:"grant_infos"`
}
