package models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

type MysqlDBPriv struct {
	ID            types.NullInt64  `json:"id" gorm:"column:id"`
	OrderUUID     types.NullString `json:"order_uuid" gorm:"column:order_uuid"`
	UserId        types.NullInt64  `json:"user_id" gorm:"column:user_id"`
	Username      types.NullString `json:"username" gorm:"column:username;unique;not null;size:50"`
	NameZh        types.NullString `json:"name_zh" gorm:"column:name_zh"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" gorm:"column:meta_cluster_id"`
	ClusterName   types.NullString `json:"cluster_name" gorm:"column:cluster_name"`
	DBName        types.NullString `json:"db_name" gorm:"column:db_name"`
	VipPort       types.NullString `json:"vip_port" gorm:"column:vip_port"`
	UpdatedAt     types.NullTime   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt     types.NullTime   `json:"created_at" gorm:"column:created_at"`
}

func (MysqlDBPriv) TableName() string {
	return "mysql_db_priv"
}

func PrivsToPrivTrashs(privs []*MysqlDBPriv) []*MysqlDBPrivTrash {
	trashs := make([]*MysqlDBPrivTrash, 0, len(privs))
	for _, priv := range privs {
		trashs = append(trashs, PrivToPrivTrash(priv))
	}

	return trashs
}

func PrivToPrivTrash(priv *MysqlDBPriv) *MysqlDBPrivTrash {
	return &MysqlDBPrivTrash{
		MysqlDBPrivId: priv.ID,
		OrderUUID:     priv.OrderUUID,
		UserId:        priv.UserId,
		Username:      priv.Username,
		NameZh:        priv.NameZh,
		MetaClusterId: priv.MetaClusterId,
		ClusterName:   priv.ClusterName,
		DBName:        priv.DBName,
		VipPort:       priv.VipPort,
	}
}
