package models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

const (
	MysqlDBPrivApplyOrderStatusUnknow = iota
	MysqlDBPrivApplyOrderStatusApplying
	MysqlDBPrivApplyOrderStatusSuccess
	MysqlDBPrivApplyOrderStatusFail
)

type MysqlDBPrivApplyOrder struct {
	ID           types.NullInt64  `json:"id" gorm:"column:id"`
	OrderUUID    types.NullString `json:"order_uuid" gorm:"column:order_uuid"`
	UserId       types.NullInt64  `json:"user_id" gorm:"column:user_id"`
	Username     types.NullString `json:"username" gorm:"column:username;unique;not null;size:50"`
	NameZh       types.NullString `json:"name_zh" gorm:"column:name_zh"`
	ApplyStatus  types.NullInt64  `json:"apply_status" gorm:"column:apply_status"`
	ApplyReason  types.NullString `json:"apply_reason" gorm:"column:apply_reason"`
	ErrorMessage types.NullString `json:"error_message" gorm:"column:error_message"`
	UpdatedAt    types.NullTime   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt    types.NullTime   `json:"created_at" gorm:"column:created_at"`
}

func (MysqlDBPrivApplyOrder) TableName() string {
	return "mysql_db_priv_apply_order"
}
