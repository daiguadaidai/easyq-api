package response

import "github.com/daiguadaidai/easyq-api/types"

type MysqlPrivsTreeResponse struct {
	ID            types.NullInt64  `json:"id"`
	MetaClusterId types.NullInt64  `json:"meta_cluster_id"`
	ClusterName   types.NullString `json:"cluster_name"`
	DBName        types.NullString `json:"db_name"`
	VipPort       types.NullString `json:"vip_port"`
}
