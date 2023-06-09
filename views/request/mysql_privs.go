package request

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/types"
)

type PrivsApplyMysqlPriv struct {
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" form:"meta_cluster_id"`
	ClusterName   types.NullString `json:"cluster_name" form:"cluster_name"`
	DBName        types.NullString `json:"db_name" form:"db_name"`
	VipPort       types.NullString `json:"vip_port" form:"vip_port"`
}

type PrivsApplyMysqlPrivRequest struct {
	ApplyReason types.NullString       `json:"apply_reason" form:"apply_reason"`
	Privs       []*PrivsApplyMysqlPriv `json:"privs" form:"privs"`
}

func (this *PrivsApplyMysqlPrivRequest) Check() error {
	if len(this.Privs) == 0 {
		return fmt.Errorf("申请权限列表为空")
	}

	for _, priv := range this.Privs {
		if priv.MetaClusterId.IsZero() {
			return fmt.Errorf("meta_cluster_id不允许为空")
		}
		if priv.DBName.IsEmpty() {
			return fmt.Errorf("db_name 不能为空")
		}
	}

	return nil
}

type PrivsApplyPrivSuccessRequest struct {
	OrderUUID types.NullString `json:"order_uuid" form:"order_uuid"`
}

func (this *PrivsApplyPrivSuccessRequest) Check() error {
	if this.OrderUUID.IsEmpty() {
		return fmt.Errorf("申请单不能为空")
	}

	return nil
}

type PrivsApplyMysqlPrivOrderRequest struct {
	Pager
	OrderUUID   types.NullString `json:"order_uuid" form:"order_uuid"`
	Username    types.NullString `json:"username" form:"username"`
	ApplyStatus types.NullInt64  `json:"apply_status" form:"apply_status"`
}

type PrivsApplyMysqlPrivByUUIDRequest struct {
	OrderUUID types.NullString `json:"order_uuid" form:"order_uuid"`
}

func (this *PrivsApplyMysqlPrivByUUIDRequest) Check() error {
	if this.OrderUUID.IsEmpty() {
		return fmt.Errorf("工单uuid不能为空")
	}

	return nil
}

type PrivsApplyMysqlPrivOrderEditByUUIDRequest struct {
	OrderUUID   types.NullString `json:"order_uuid" form:"order_uuid"`
	ApplyStatus types.NullInt64  `json:"apply_status" form:"apply_status"`
}

func (this *PrivsApplyMysqlPrivOrderEditByUUIDRequest) Check() error {
	if this.OrderUUID.IsEmpty() {
		return fmt.Errorf("工单uuid不能为空")
	}

	if this.ApplyStatus.IsZero() {
		return fmt.Errorf("工单状态不能为空")
	}

	return nil
}

type PrivsMysqlFindTreeByUsername struct {
	Username types.NullString `json:"username" form:"username"`
}

func (this *PrivsMysqlFindTreeByUsername) Check() error {
	if this.Username.IsEmpty() {
		return fmt.Errorf("获取数据库权限信息, 用户信息不能为空")
	}

	return nil
}

type PrivsMysqlFindTablesByUserRequest struct {
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" form:"meta_cluster_id"`
	DBName        types.NullString `json:"db_name" form:"db_name"`
}

func (this *PrivsMysqlFindTablesByUserRequest) Check() error {
	if this.MetaClusterId.IsZero() {
		return fmt.Errorf("集群信息不能为空")
	}
	if this.DBName.IsEmpty() {
		return fmt.Errorf("数据库名不能为空")
	}

	return nil
}
