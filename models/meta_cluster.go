package models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

const (
	// 集群类别
	ClusterCategoryUnknow          = 0 // 未知
	ClusterCategoryTceSelfMySQL    = 1 // TCE自建MySQL
	ClusterCategoryPublicSelfMySQL = 2 // 公有云自建MySQL
	ClusterCategoryTceTDSQL        = 3 // TEC-TDSQL
	ClusterCategoryPublicTDSQLC    = 4 // 公有云TDSQL-C
	ClusterCategoryPublicMySQL     = 5 // 公有云MySQL
	ClusterCategoryFuZhouMySQL     = 6 // 福州MySQL(自建)
	ClusterCategoryTIDB            = 7 // tidb

	// 分库分表类型
	ShardTypeNone               = ""
	ShardTypeMycat              = "mycat"
	ShardTypeTDSQL              = "tdsql"
	ShardTypeZebra              = "zebra"
	ShardTypeShardingJdbc       = "sharding-jdbc"
	ShardTypeShardingJdbcOnlyDB = "sharding-jdbc-only-db"

	// 分库分表后缀递增类型
	ShardIncrementTypeNone   = ""
	ShardIncrementTypeGlobal = "global" // 全局递增
	ShardIncrementTypeOneDB  = "one-db" // 一个数据库中递增
)

type MetaCluster struct {
	ID           types.NullInt64  `json:"id" gorm:"column:id"`
	Name         types.NullString `json:"name" gorm:"column:name;unique;not null;size:50"`
	ClusterId    types.NullString `json:"cluster_id" gorm:"column:cluster_id;unique;not null;default:'';size:50"`
	BusinessLine types.NullString `json:"business_line" gorm:"column:business_line;unique;not null;default:'';size:100"`
	Owner        types.NullString `json:"owner" gorm:"column:owner;unique;not null;default:'';size:50"`
	IsDeleted    types.NullInt64  `json:"is_deleted" gorm:"column:is_deleted;not null;default:0"`
	DomainName   types.NullString `json:"domain_name" gorm:"column:domain_name;not null;default:'';size:500"`
	VipPort      types.NullString `json:"vip_port" gorm:"column:vip_port;not null;default:'';size:30"`
	VpcgwVipPort types.NullString `json:"vpcgw_vip_port" gorm:"column:vpcgw_vip_port;not null;default:'';size:30"`
	ReadHostPort types.NullString `json:"read_host_port" gorm:"column:read_host_port;not null;default:'';size:30"`
	IsShard      types.NullInt64  `json:"is_shard" gorm:"column:is_shard;not null;default:0"`
	ShardType    types.NullString `json:"shard_type" gorm:"column:shard_type;not null;default:'';size:20"`
	Category     types.NullInt64  `json:"category" gorm:"column:category;not null;default:0"`
	SetName      types.NullString `json:"set_name" gorm:"column:set_name;not null;default:'';size:30"`
	UpdatedAt    types.NullTime   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt    types.NullTime   `json:"created_at" gorm:"column:created_at"`
}

func (MetaCluster) TableName() string {
	return "meta_cluster"
}
