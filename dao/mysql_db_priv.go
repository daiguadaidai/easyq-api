package dao

import (
	"github.com/daiguadaidai/easyq-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlDBPrivDao struct {
	DB *gorm.DB
}

func NewMysqlDBPrivDao(db *gorm.DB) *MysqlDBPrivDao {
	return &MysqlDBPrivDao{
		DB: db,
	}
}

// 批量创建权限信息
func (this *MysqlDBPrivDao) BatchReplace(privs []*models.MysqlDBPriv) error {
	return this.DB.Model(&models.MysqlDBPriv{}).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"user_id", "order_uuid", "username", "name_zh", "meta_cluster_id", "cluster_name", "db_name", "vip_port"}),
	}).Create(&privs).Error
}