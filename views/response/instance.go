package response

import "github.com/daiguadaidai/easyq-api/types"

type InstanceClusterResponse struct {
	ID            types.NullInt64            `json:"id"`
	InstanceId    types.NullString           `json:"instance_id"`
	InstanceName  types.NullString           `json:"instance_name"`
	MetaClusterId types.NullInt64            `json:"meta_cluster_id"`
	MachineHost   types.NullString           `json:"machine_host"`
	VpcgwRip      types.NullString           `json:"vpcgw_rip"`
	Port          types.NullInt64            `json:"port"`
	MasterHost    types.NullString           `json:"master_host"`
	MasterPort    types.NullInt64            `json:"master_port"`
	VipPort       types.NullString           `json:"vip_port"`
	VpcgwVipPort  types.NullString           `json:"vpcgw_vip_port"`
	Role          types.NullString           `json:"role"`
	Cpu           types.NullInt64            `json:"cpu"`
	Mem           types.NullInt64            `json:"mem"`
	Disk          types.NullInt64            `json:"disk"`
	Version       types.NullString           `json:"version"`
	IsMaintenance types.NullInt64            `json:"is_maintenance"`
	IsDeleted     types.NullInt64            `json:"is_deleted"`
	Descript      types.NullString           `json:"descript"`
	UpdatedAt     types.NullTime             `json:"updated_at"`
	CreatedAt     types.NullTime             `json:"created_at"`
	ClusterName   types.NullString           `json:"cluster_name"`
	ClusterId     types.NullString           `json:"cluster_id"`
	SetName       types.NullString           `json:"set_name"`
	Category      types.NullInt64            `json:"category"`
	IsShard       types.NullInt64            `json:"is_shard"`
	ShardType     types.NullString           `json:"shard_type"`
	Slaves        []*InstanceClusterResponse `json:"slaves"`
}
