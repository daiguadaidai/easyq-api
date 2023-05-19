package view_models

import "github.com/daiguadaidai/easyq-api/types"

type NeedArchiveTableStat struct {
	TableCount                         types.NullInt64 `json:"table_count"`
	NewTableCount                      types.NullInt64 `json:"new_table_count"`
	DataLength                         types.NullInt64 `json:"data_length"`
	IndexLength                        types.NullInt64 `json:"index_length"`
	DataFree                           types.NullInt64 `json:"data_free"`
	FilterArchiveStatusUnrunCount      types.NullInt64 `json:"filter_archive_status_unrun_count"`
	FilterArchiveStatusRunningCount    types.NullInt64 `json:"filter_archive_status_running_count"`
	FilterArchiveStatusSchedulingCount types.NullInt64 `json:"filter_archive_status_scheduling_count"`
}
