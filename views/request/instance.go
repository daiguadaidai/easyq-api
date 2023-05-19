package request

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/types"
)

type InstanceFindRequest struct {
	Pager
	Keyword types.NullString `json:"keyword" form:"keyword"`
}

type InstanceFindMasterSlavesRequest struct {
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" form:"meta_cluster_id"`
	ClusterName   types.NullString `json:"cluster_name" form:"cluster_name"`
	SetName       types.NullString `json:"set_name" form:"set_name"`
}

type InstanceAddRequest struct {
	InstanceId    types.NullString `json:"instance_id" form:"instance_id"`
	InstanceName  types.NullString `json:"instance_name" form:"instance_name"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" form:"meta_cluster_id" binding:"required"`
	MachineHost   types.NullString `json:"machine_host" form:"machine_host"`
	VpcgwRip      types.NullString `json:"vpcgw_rip" form:"vpcgw_rip"`
	Port          types.NullInt64  `json:"port" form:"port"`
	MasterHost    types.NullString `json:"master_host" form:"master_host"`
	MasterPort    types.NullInt64  `json:"master_port" form:"master_port"`
	VipPort       types.NullString `json:"vip_port" form:"vip_port"`
	VpcgwVipPort  types.NullString `json:"vpcgw_vip_port" form:"vpcgw_vip_port"`
	Role          types.NullString `json:"role" form:"role" binding:"required"`
	Cpu           types.NullInt64  `json:"cpu" form:"cpu"`
	Mem           types.NullInt64  `json:"mem" form:"mem"`
	Disk          types.NullInt64  `json:"disk" form:"disk"`
	Version       types.NullString `json:"version" form:"version"`
	IsMaintenance types.NullInt64  `json:"is_maintenance" form:"is_maintenance"`
	Descript      types.NullString `json:"descript" form:"descript"`
}

func (this *InstanceAddRequest) Check() error {
	if this.InstanceId.IsEmpty() && this.InstanceName.IsEmpty() && this.MachineHost.IsEmpty() && this.VpcgwRip.IsEmpty() {
		return fmt.Errorf("实例id/实例名称/机器host/vpcgw host 4个字段必须要有一个有值")
	}

	if this.MetaClusterId.IsZero() {
		return fmt.Errorf("必须选择一个集群")
	}

	if this.Role.IsEmpty() {
		return fmt.Errorf("必须选择实例角色")
	}

	return nil
}

type InstanceEditByIdRequest struct {
	ID            types.NullInt64  `json:"id" form:"id" binding:"required"`
	InstanceId    types.NullString `json:"instance_id" form:"instance_id"`
	InstanceName  types.NullString `json:"instance_name" form:"instance_name"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id" form:"meta_cluster_id"`
	MachineHost   types.NullString `json:"machine_host" form:"machine_host"`
	VpcgwRip      types.NullString `json:"vpcgw_rip" form:"vpcgw_rip"`
	Port          types.NullInt64  `json:"port" form:"port"`
	MasterHost    types.NullString `json:"master_host" form:"master_host"`
	MasterPort    types.NullInt64  `json:"master_port" form:"master_port"`
	VipPort       types.NullString `json:"vip_port" form:"vip_port"`
	VpcgwVipPort  types.NullString `json:"vpcgw_vip_port" form:"vpcgw_vip_port"`
	Role          types.NullString `json:"role" form:"role"`
	Cpu           types.NullInt64  `json:"cpu" form:"cpu"`
	Mem           types.NullInt64  `json:"mem" form:"mem"`
	Disk          types.NullInt64  `json:"disk" form:"disk"`
	Version       types.NullString `json:"version" form:"version"`
	IsMaintenance types.NullInt64  `json:"is_maintenance" form:"is_maintenance"`
	Descript      types.NullString `json:"descript" form:"descript"`
}

type InstanceDeleteByIdRequest struct {
	ID types.NullInt64 `json:"id" form:"id" binding:"required"`
}
