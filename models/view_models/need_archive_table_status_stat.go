package view_models

import "github.com/daiguadaidai/easyq-api/types"

type NeedArchiveTableStatusStat struct {
	TotalArchiveStatusUnrunCount      types.NullInt64 `json:"total_archive_status_unrun_count"`
	TotalArchiveStatusRunningCount    types.NullInt64 `json:"total_archive_status_running_count"`
	TotalArchiveStatusSchedulingCount types.NullInt64 `json:"total_archive_status_scheduling_count"`
}
