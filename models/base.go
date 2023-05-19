package models

const (
	// 是否删除
	IsDeletedNo  = 0 // 否
	IsDeletedYes = 1 // 是

	// 是否分库分表
	IsShardNo  = 0 // 否
	IsShardYes = 1 // 是

	StorageTypeSingle = "single"
	StorageTypeShard  = "shard"
	StorageTypeTIDB   = "tidb"

	Underscore = "_"
)
