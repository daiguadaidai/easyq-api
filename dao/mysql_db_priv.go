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

func (this *MysqlDBPrivDao) FindPrivsTreeByUsername(username string) ([]*models.MysqlDBPriv, error) {
	var privs []*models.MysqlDBPriv
	if err := this.DB.Model(&models.MysqlDBPriv{}).
		Select("id, meta_cluster_id, cluster_name, db_name, vip_port").
		Where("username = ?", username).
		Find(&privs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return privs, nil
}

func (this *MysqlDBPrivDao) GetByUsernameClusterDB(username string, meta_cluster_id int64, db_name string) (*models.MysqlDBPriv, error) {
	var priv models.MysqlDBPriv
	if err := this.DB.Model(&models.MysqlDBPriv{}).
		Where("username = ? AND meta_cluster_id = ? AND db_name = ?", username, meta_cluster_id, db_name).
		First(&priv).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &priv, nil
}

func (this *MysqlDBPrivDao) CountByUsernameClusterDB(username string, meta_cluster_id int64, db_name string) (int64, error) {
	var cnt int64
	if err := this.DB.Model(&models.MysqlDBPriv{}).
		Where("username = ? AND meta_cluster_id = ? AND db_name = ?", username, meta_cluster_id, db_name).
		Count(&cnt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}

		return 0, err
	}

	return cnt, nil
}

func (this *MysqlDBPrivDao) GetById(privId int64) (*models.MysqlDBPriv, error) {
	var priv models.MysqlDBPriv
	if err := this.DB.Model(&models.MysqlDBPriv{}).Where("id = ?", privId).First(&priv).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &priv, nil
}
