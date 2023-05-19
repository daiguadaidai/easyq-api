package response

import "github.com/daiguadaidai/easyq-api/types"

type ClusterResponse struct {
	ID            types.NullInt64    `json:"id"`
	Name          types.NullString   `json:"name"`
	ClusterId     types.NullString   `json:"cluster_id" form:"cluster_id"`
	BusinessLine  types.NullString   `json:"business_line"`
	Owner         types.NullString   `json:"owner"`
	IsDeleted     types.NullInt64    `json:"is_deleted"`
	DomainNames   []types.NullString `json:"domain_names"`
	VipPorts      []types.NullString `json:"vip_ports"`
	VpcgwVipPorts []types.NullString `json:"vpcgw_vip_ports"`
	ReadHostPort  types.NullString   `json:"read_host_port"`
	IsShard       types.NullInt64    `json:"is_shard"`
	ShardType     types.NullString   `json:"shard_type" form:"shard_type"`
	Category      types.NullInt64    `json:"category"`
	SetName       types.NullString   `json:"set_name"`
	UpdatedAt     types.NullTime     `json:"updated_at"`
	CreatedAt     types.NullTime     `json:"created_at"`
}

type ClusterNameResponse struct {
	ID   types.NullInt64  `json:"id"`
	Name types.NullString `json:"name"`
}
