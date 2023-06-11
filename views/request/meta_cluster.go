package request

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/types"
)

type MCFindRequest struct {
	Pager
	Keyword types.NullString `json:"keyword" form:"keyword"`
}

type MCFindNameByNameRequest struct {
	Name  types.NullString `json:"name" form:"name"`
	Limit types.NullInt64  `json:"limit" form:"limit"`
}

func (this *MCFindNameByNameRequest) GetLimit() int64 {
	if this.Limit.IsZero() {
		return DefaultLimit
	}
	return DefaultLimit
}

type MCAddRequest struct {
	Name          types.NullString   `json:"name" form:"name"`
	ClusterId     types.NullString   `json:"cluster_id" form:"cluster_id"`
	Owner         types.NullString   `json:"owner" form:"owner"`
	BusinessLine  types.NullString   `json:"business_line" form:"business_line"`
	DomainNames   []types.NullString `json:"domain_names" form:"domain_names"`
	VipPorts      []types.NullString `json:"vip_ports" form:"vip_ports"`
	VpcgwVipPorts []types.NullString `json:"vpcgw_vip_ports" form:"vpcgw_vip_ports"`
	ReadHostPort  types.NullString   `json:"read_host_port" form:"read_host_port"`
	IsShard       types.NullInt64    `json:"is_shard" form:"is_shard"`
	ShardType     types.NullString   `json:"shard_type" form:"shard_type"`
	Category      types.NullInt64    `json:"category" form:"category"`
	SetName       types.NullString   `json:"set_name" form:"set_name"`
}

func (this *MCAddRequest) Check() error {
	if this.ClusterId.IsEmpty() {
		return fmt.Errorf("集群名不能为空")
	}

	return nil
}

type MCEditByIdRequest struct {
	ID            types.NullInt64    `json:"id" form:"id" binding:"required"`
	Name          types.NullString   `json:"name" form:"name"`
	ClusterId     types.NullString   `json:"cluster_id" form:"cluster_id"`
	BusinessLine  types.NullString   `json:"business_line" form:"business_line"`
	Owner         types.NullString   `json:"owner" form:"owner"`
	DomainNames   []types.NullString `json:"domain_names" form:"domain_names"`
	VipPorts      []types.NullString `json:"vip_ports" form:"vip_ports"`
	VpcgwVipPorts []types.NullString `json:"vpcgw_vip_ports" form:"vpcgw_vip_ports"`
	ReadHostPort  types.NullString   `json:"read_host_port" form:"read_host_port"`
	IsShard       types.NullInt64    `json:"is_shard" form:"is_shard"`
	ShardType     types.NullString   `json:"shard_type" form:"shard_type"`
	Category      types.NullInt64    `json:"category" form:"category"`
	IsDeleted     types.NullInt64    `json:"is_deleted" form:"is_deleted"`
	SetName       types.NullString   `json:"set_name" form:"set_name"`
}

type MCDeleteByIdRequest struct {
	ID types.NullInt64 `json:"id" form:"id" binding:"required"`
}
