package dao

import (
	"github.com/daiguadaidai/easyq-api/models"
	"gorm.io/gorm"
)

type MysqlDBPrivTrashDao struct {
	DB *gorm.DB
}

func NewMysqlDBPrivTrashDao(db *gorm.DB) *MysqlDBPrivTrashDao {
	return &MysqlDBPrivTrashDao{
		DB: db,
	}
}

// 批量创建权限申请信息
func (this *MysqlDBPrivTrashDao) BatchCreate(trashs []*models.MysqlDBPrivTrash) error {
	return this.DB.CreateInBatches(trashs, 100).Error
}
