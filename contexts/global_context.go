package contexts

import (
	"github.com/daiguadaidai/easyq-api/config"
	"gorm.io/gorm"
)

type GlobalContext struct {
	Cfg      *config.ServerConfig `json:"cfg"`
	EasydbDB *gorm.DB             `json:"easydb_db"`
	EasyqDB  *gorm.DB             `json:"easyq_db"`
}
