package view_models

import "github.com/daiguadaidai/easyq-api/types"

type NeedArchiveBusinessSpace struct {
	BusinessLine      types.NullString `json:"business_line" gorm:"column:business_line"`
	TableRows         types.NullInt64  `json:"table_rows" gorm:"column:table_rows"`
	DataLength        types.NullInt64  `json:"data_length" gorm:"column:data_length"`
	IndexLength       types.NullInt64  `json:"index_length" gorm:"column:index_length"`
	DataFree          types.NullInt64  `json:"data_free" gorm:"column:data_free"`
	BeforeTableRows   types.NullInt64  `json:"before_table_rows" gorm:"column:before_table_rows"`
	BeforeDataLength  types.NullInt64  `json:"before_data_length" gorm:"column:before_data_length"`
	BeforeIndexLength types.NullInt64  `json:"before_index_length" gorm:"column:before_index_length"`
	BeforeDataFree    types.NullInt64  `json:"before_data_free" gorm:"column:before_data_free"`
}
