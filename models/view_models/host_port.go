package view_models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

type HostPort struct {
	ID            types.NullInt64  `json:"id" gorm:"column:id"`
	Host          types.NullString `json:"host" gorm:"column:host"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" gorm:"column:meta_cluster_id"`
	Port          types.NullInt64  `json:"port" gorm:"column:port"`
	Role          types.NullString `json:"role" gorm:"column:role"`
}
