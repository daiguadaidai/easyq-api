package view_models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

type MySQLKillTaskView struct {
	ID                      types.NullInt64  `json:"id"`
	ExecUserId              types.NullInt64  `json:"exec_user_id"`
	ExecUserNameZh          types.NullString `json:"exec_user_name_zh"`
	MetaClusterId           types.NullInt64  `json:"meta_cluster_id"`
	InstanceId              types.NullInt64  `json:"instance_id"`
	ClusterName             types.NullString `json:"cluster_name"`
	ExecStatus              types.NullInt64  `json:"exec_status"`
	InstanceHost            types.NullString `json:"instance_host"`
	InstancePort            types.NullInt64  `json:"instance_port"`
	InstanceHostPort        types.NullString `json:"instance_host_port"`
	ProcesslistHosts        []string         `json:"processlist_hosts"`
	ProcesslistUsers        []string         `json:"processlist_users"`
	ProcesslistDBs          []string         `json:"processlist_dbs"`
	ProcesslistCommandTypes []string         `json:"processlist_command_types"`
	ProcesslistStmtTypes    []string         `json:"processlist_stmt_types"`
	ProcesslistTime         types.NullInt64  `json:"processlist_time"`
	Duration                types.NullInt64  `json:"duration"`
	Interval                types.NullInt64  `json:"interval"`
	ExecCnt                 types.NullInt64  `json:"exec_cnt"`
	KillLogFile             types.NullString `json:"kill_log_file"`
	ErrorMessage            types.NullString `json:"error_message"`
	StartTime               types.NullTime   `json:"start_time"`
	EndTime                 types.NullTime   `json:"end_time"`
	UpdatedAt               types.NullTime   `json:"updated_at"`
	CreatedAt               types.NullTime   `json:"created_at"`
}
