package dao

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/models"
	"gorm.io/gorm"
)

type MetaClusterDao struct {
	db *gorm.DB
}

func NewMetaClusterDao(db *gorm.DB) *MetaClusterDao {
	return &MetaClusterDao{
		db: db,
	}
}

// 插入用户，先检查是否存在用户，如果没有则存入
func (this *MetaClusterDao) Create(cluster *models.MetaCluster) error {
	// 通过名称获取集群
	oldCluster, err := this.GetByName(cluster.Name.String)
	if err != nil {
		return err
	}
	if oldCluster != nil {
		return fmt.Errorf("集群已经存在: %v", cluster.Name.String)
	}

	return this.db.Create(cluster).Error
}

func (this *MetaClusterDao) UpdateById(cluster *models.MetaCluster) error {
	if cluster.ID.Int64 <= 0 {
		return fmt.Errorf("集群id不能<0")
	}

	return this.db.Model(cluster).Omit("id").Updates(cluster).Error
}

func (this *MetaClusterDao) getFindByKeywordQuery(keyword string) string {
	// 获取 LIKE OR LIKE WHERE 语句
	likeClauses := GetLikeClausesByKeyWords(keyword, "name", "cluster_id", "domain_name", "vip_port", "vpcgw_vip_port", "set_name", "read_host_port")
	likeClauseStr := JoinOrClauses(likeClauses...)

	// 获取其他语句
	otherClauseStr := "is_deleted=0"

	// LIKE 和 其他语句合并
	return JoinAndClauses(likeClauseStr, otherClauseStr)
}

func (this *MetaClusterDao) FindByKeyword(keyword string, offset, limit int) ([]*models.MetaCluster, error) {
	query := this.getFindByKeywordQuery(keyword)
	logger.M.Debugf("[MetaClusterDao] FindByKeyword. query: %s", query)

	var clusters []*models.MetaCluster
	if err := this.db.Where(query).Offset(offset).Limit(limit).Order("updated_at DESC").Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过keyword获取集群信息失败. %v", err.Error())
	}

	return clusters, nil
}

func (this *MetaClusterDao) CountByKeyword(keyword string) (int, error) {
	query := this.getFindByKeywordQuery(keyword)
	logger.M.Debugf("[MetaClusterDao] CountByKeyword. query: %s", query)

	var cnt int64
	if err := this.db.Model(&models.MetaCluster{}).Where(query).Count(&cnt).Error; err != nil {
		return 0, fmt.Errorf("通过keyword获取集群信息数失败. %v", err.Error())
	}

	return int(cnt), nil
}

func (this *MetaClusterDao) FindNameByName(name string, limit int64) ([]*models.MetaCluster, error) {
	query := fmt.Sprintf("name LIKE '%%%v%%' AND is_deleted = 0", name)
	logger.M.Debugf("[MetaClusterDao] FindByName. query: %s", query)

	var clusters []*models.MetaCluster
	if err := this.db.Select("id, name").Where(query).Limit(int(limit)).Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过name搜索集群名失败. %v", err.Error())
	}

	return clusters, nil
}

func (this *MetaClusterDao) AllName() ([]*models.MetaCluster, error) {
	var clusters []*models.MetaCluster
	if err := this.db.Select("id, name").Where("is_deleted=0").Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("获取所有集群名. %v", err.Error())
	}

	return clusters, nil
}

func (this *MetaClusterDao) All() ([]*models.MetaCluster, error) {
	var clusters []*models.MetaCluster
	if err := this.db.Where("is_deleted=0").Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("获取所有集群. %v", err.Error())
	}

	return clusters, nil
}

func (this *MetaClusterDao) GetByName(name string) (*models.MetaCluster, error) {
	var cluster models.MetaCluster
	if err := this.db.Where("name = ? AND is_deleted = 0", name).First(&cluster).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("通过集群名获取集群信息失败. %v", err.Error())
	}

	return &cluster, nil
}

func (this *MetaClusterDao) GetById(id int64) (*models.MetaCluster, error) {
	var cluster models.MetaCluster
	if err := this.db.Where("id = ? AND is_deleted = 0", id).First(&cluster).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("通过集群id取集群信息失败. %v", err.Error())
	}

	return &cluster, nil
}

func (this *MetaClusterDao) GetByIdEmptyError(id int64) (*models.MetaCluster, error) {
	cluster, err := this.GetById(id)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, fmt.Errorf("cluster Id 无法获取到集群信息, clusterId: %v", id)
	}

	return cluster, nil
}

func (this *MetaClusterDao) FindByIds(ids []int64) ([]*models.MetaCluster, error) {
	var clusters []*models.MetaCluster
	if err := this.db.Where("id IN (?)", ids).Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("通过自增ids获取集群信息失败. %v", err.Error())
	}

	return clusters, nil
}
