package view_models

import "github.com/daiguadaidai/easyq-api/types"

type ArchiveTableStatusCount struct {
	ArchiveTableStatus types.NullInt64 `json:"archive_task_status" gorm:"column:archive_task_status"`
	Cnt                types.NullInt64 `json:"cnt" gorm:"column:cnt"`
}
