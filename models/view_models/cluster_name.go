package view_models

import "github.com/daiguadaidai/easyq-api/types"

type ClusterNameView struct {
	MetaClusterId types.NullInt64  `json:"meta_cluster_id"`
	ClusterName   types.NullString `json:"cluster_name"`
}
