package dao

import (
	"github.com/daiguadaidai/easyq-api/models"
	"gorm.io/gorm"
)

type MysqlDBPrivApplyDao struct {
	DB *gorm.DB
}

func NewMysqlDBPrivApplyDao(db *gorm.DB) *MysqlDBPrivApplyDao {
	return &MysqlDBPrivApplyDao{
		DB: db,
	}
}

// 批量创建权限申请信息
func (this *MysqlDBPrivApplyDao) BatchCreate(applys []*models.MysqlDBPrivApply) error {
	return this.DB.CreateInBatches(applys, 100).Error
}

// 通过uuid 获取申请权限
func (this *MysqlDBPrivApplyDao) FindByUUID(uuid string) ([]*models.MysqlDBPrivApply, error) {
	var applys []*models.MysqlDBPrivApply

	if err := this.DB.Model(&models.MysqlDBPrivApply{}).Where("order_uuid=?", uuid).Find(&applys).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil ,nil
		}

		return nil, err
	}

	return applys, nil
}
