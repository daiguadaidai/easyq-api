package view_models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

type ShardInfoDetailCluster struct {
	ID            types.NullInt64  `json:"id" gorm:"column:id"`
	ShardInfoId   types.NullInt64  `json:"shard_info_id" gorm:"shard_info_id:meta_cluster_id;not null"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" gorm:"column:meta_cluster_id;not null"`
	VpcgwVipPort  types.NullString `json:"vpcgw_vip_port" gorm:"column:vpcgw_vip_port;not null;default:'';size:30"`
	VipPort       types.NullString `json:"vip_port" gorm:"column:vip_port;not null;default:'';size:30"`
	ClusterName   types.NullString `json:"cluster_name" gorm:"column:cluster_name;not null;default:'';size:100"`
	ClusterId     types.NullString `json:"cluster_id" gorm:"column:cluster_id;not null;default:'';size:50"`
	SetName       types.NullString `json:"set_name" gorm:"column:set_name;not null;default:'';size:50"`
	DBNames       types.NullString `json:"db_names" gorm:"column:db_names;not null;default:''"`
	IsDeleted     types.NullInt64  `json:"is_deleted" gorm:"column:is_deleted;not null;default:0"`
	UpdatedAt     types.NullTime   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt     types.NullTime   `json:"created_at" gorm:"column:created_at"`
}
