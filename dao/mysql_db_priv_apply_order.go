package dao

import (
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/types"
	"gorm.io/gorm"
)

type MysqlDBPrivApplyOrderDao struct {
	DB *gorm.DB
}

func NewMysqlDBPrivApplyOrderDao(db *gorm.DB) *MysqlDBPrivApplyOrderDao {
	return &MysqlDBPrivApplyOrderDao{
		DB: db,
	}
}

func (this *MysqlDBPrivApplyOrderDao) GetByUUID(uuid string) (*models.MysqlDBPrivApplyOrder, error) {
	var order models.MysqlDBPrivApplyOrder
	if err := this.DB.Where("order_uuid=?", uuid).Limit(1).Find(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &order, nil
}

// 批量创建权限申请信息
func (this *MysqlDBPrivApplyOrderDao) Create(order *models.MysqlDBPrivApplyOrder) error {
	return this.DB.Create(order).Error
}

func (this *MysqlDBPrivApplyOrderDao) UpdateStatusApplying(id int64) error {
	order := &models.MysqlDBPrivApplyOrder{
		ID:          types.NewNullInt64(id, false),
		ApplyStatus: types.NewNullInt64(models.MysqlDBPrivApplyOrderStatusApplying, false),
	}
	return this.DB.Model(order).Updates(order).Error
}

func (this *MysqlDBPrivApplyOrderDao) UpdateStatusSuccess(id int64) error {
	order := &models.MysqlDBPrivApplyOrder{
		ID:          types.NewNullInt64(id, false),
		ApplyStatus: types.NewNullInt64(models.MysqlDBPrivApplyOrderStatusSuccess, false),
	}
	return this.DB.Model(order).Updates(order).Error
}

func (this *MysqlDBPrivApplyOrderDao) UpdateStatusFail(id int64) error {
	order := &models.MysqlDBPrivApplyOrder{
		ID:          types.NewNullInt64(id, false),
		ApplyStatus: types.NewNullInt64(models.MysqlDBPrivApplyOrderStatusFail, false),
	}
	return this.DB.Model(order).Updates(order).Error
}
