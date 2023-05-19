package models

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/types"
)

const (
	// 实例角色
	InstanceRoleUnknow = ""
	InstanceRoleMaster = "master"
	InstanceRoleSlave  = "slave"
	InstanceRoleBackup = "backup"

	// 是否维护
	InstanceIsMaintenanceNo  = 0 // 否
	InstanceIsMaintenanceYes = 1 // 是
)

type Instance struct {
	ID            types.NullInt64  `json:"id" gorm:"column:id"`
	InstanceId    types.NullString `json:"instance_id" gorm:"column:instance_id;unique;not null;default:'';size:50"`
	InstanceName  types.NullString `json:"instance_name" gorm:"column:instance_name;not null;default:'';size:100"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" gorm:"column:meta_cluster_id;not null;default:0"`
	MachineHost   types.NullString `json:"machine_host" gorm:"column:machine_host;not null;default:'';size:15"`
	VpcgwRip      types.NullString `json:"vpcgw_rip" gorm:"column:vpcgw_rip;not null;default:'';size:15"`
	Port          types.NullInt64  `json:"port" gorm:"column:port;not null;default:0"`
	MasterHost    types.NullString `json:"master_host" gorm:"column:master_host;not null;default:'';size:15"`
	MasterPort    types.NullInt64  `json:"master_port" gorm:"column:master_port;not null;default:0"`
	VipPort       types.NullString `json:"vip_port" gorm:"column:vip_port;not null;default:'';size:30"`
	VpcgwVipPort  types.NullString `json:"vpcgw_vip_port" gorm:"column:vpcgw_vip_port;not null;default:'';size:30"`
	Role          types.NullString `json:"role" gorm:"column:role;not null;default:'';size:15"`
	Cpu           types.NullInt64  `json:"cpu" gorm:"column:cpu;not null;default:0"`
	Mem           types.NullInt64  `json:"mem" gorm:"column:mem;not null;default:0"`
	Disk          types.NullInt64  `json:"disk" gorm:"column:disk;not null;default:0"`
	Version       types.NullString `json:"version" gorm:"column:version;not null;default:'';size:15"`
	IsMaintenance types.NullInt64  `json:"is_maintenance" gorm:"column:is_maintenance;not null;default:0"`
	IsDeleted     types.NullInt64  `json:"is_deleted" gorm:"column:is_deleted;not null;default:0"`
	Descript      types.NullString `json:"descript" gorm:"column:descript"`
	UpdatedAt     types.NullTime   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt     types.NullTime   `json:"created_at" gorm:"column:created_at"`
}

func (Instance) TableName() string {
	return "instance"
}

func (Instance) TableNameAndAlias(alias string) string {
	return fmt.Sprintf("instance AS %s", alias)
}
