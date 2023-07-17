package dao

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/models/view_models"
	"github.com/daiguadaidai/easyq-api/utils"
	"gorm.io/gorm"
)

type InstanceDao struct {
	db *gorm.DB
}

func NewInstanceDao(db *gorm.DB) *InstanceDao {
	return &InstanceDao{
		db: db,
	}
}

// 插入用户，先检查是否存在用户，如果没有则存入
func (this *InstanceDao) Create(instance *models.Instance) error {
	return this.db.Create(instance).Error
}

func (this *InstanceDao) UpdateById(instance *models.Instance) error {
	if instance.ID.Int64 <= 0 {
		return fmt.Errorf("实例id不能<0")
	}

	return this.db.Model(instance).Omit("id").Updates(instance).Error
}

func (this *InstanceDao) getFindByKeywordQuery(keyword string) string {
	// 获取 LIKE OR LIKE WHERE 语句
	likeClauses := make([]string, 0, 5)
	host, port, isHostPort := utils.AddrToHostPort(keyword)
	if isHostPort { // 是hostport的字符串
		machineClause := fmt.Sprintf("(i.machine_host = %#v AND i.port = %v)", host, port)
		vpcgwRipClause := fmt.Sprintf("(i.vpcgw_rip = %#v AND i.port = %v)", host, port)
		likeClauses = append(likeClauses, machineClause, vpcgwRipClause)
	} else { // 如果不是host port 就将machine_host和vpcgw_rip变成vpcgw ip
		hostLikeClauses := GetLikeClausesByKeyWords(keyword,
			"c.name", "c.cluster_id", "c.set_name", "i.instance_id", "i.instance_name", "i.machine_host",
			"i.vpcgw_rip", "i.vip_port", "i.vpcgw_vip_port")
		likeClauses = append(likeClauses, hostLikeClauses...)
	}

	likeClauseStr := JoinOrClauses(likeClauses...)

	// 获取其他语句
	otherClauseStr := "i.is_deleted=0"

	// LIKE 和 其他语句合并
	return JoinAndClauses(likeClauseStr, otherClauseStr)
}

func (this *InstanceDao) FindByKeyword(keyword string, offset, limit int) ([]*view_models.InstanceCluster, error) {
	// LIKE 和 其他语句合并
	query := this.getFindByKeywordQuery(keyword)
	logger.M.Debugf("[InstanceDao] FindByKeyword. query: %s", query)

	var instances []*view_models.InstanceCluster
	if err := this.db.Table(models.Instance{}.TableNameAndAlias("i")).
		Select("i.*, c.name AS cluster_name, c.cluster_id, c.set_name, c.category, c.is_shard, c.shard_type").
		Joins("LEFT JOIN meta_cluster AS c ON i.meta_cluster_id = c.id").
		Where(query).
		Order("i.updated_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&instances).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过keyword获取实例集群信息失败. %v", err.Error())
	}

	return instances, nil
}

func (this *InstanceDao) CountByKeyword(keyword string) (int, error) {
	query := this.getFindByKeywordQuery(keyword)
	logger.M.Debugf("[InstanceDao] CountByKeyword. query: %s", query)

	var cnt int64
	if err := this.db.Table(models.Instance{}.TableNameAndAlias("i")).
		Joins("LEFT JOIN meta_cluster AS c ON i.meta_cluster_id = c.id").
		Where(query).
		Count(&cnt).Error; err != nil {

		return 0, fmt.Errorf("通过keyword获取实例信息数失败. %v", err.Error())
	}

	return int(cnt), nil
}

// 通过集群id获取所有的实例
func (this *InstanceDao) FindByClusterId(clusterId int64) ([]*models.Instance, error) {
	var instances []*models.Instance
	if err := this.db.Model(&models.Instance{}).
		Where("meta_cluster_id = ? AND is_deleted=0", clusterId).
		Order("updated_at DESC").
		Find(&instances).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过集群id获取实例信息失败. %v", err.Error())
	}

	return instances, nil
}

// 通过集群id获取所有的实例
func (this *InstanceDao) FindInstanceClusterByClusterId(clusterId int64) ([]*view_models.InstanceCluster, error) {
	var instances []*view_models.InstanceCluster
	if err := this.db.Table(models.Instance{}.TableNameAndAlias("i")).
		Select("i.*, c.name AS cluster_name, c.cluster_id, c.set_name, c.category, c.is_shard, c.shard_type").
		Joins("LEFT JOIN meta_cluster AS c ON i.meta_cluster_id = c.id").
		Where("i.meta_cluster_id = ? AND i.is_deleted=0", clusterId).
		Order("i.updated_at DESC").
		Find(&instances).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过集群id获取实例集群信息失败. %v", err.Error())
	}

	return instances, nil
}

// 通过集群Name获取所有的实例
func (this *InstanceDao) FindInstanceClusterByClusterName(clusterName string) ([]*view_models.InstanceCluster, error) {
	var instances []*view_models.InstanceCluster
	if err := this.db.Table(models.Instance{}.TableNameAndAlias("i")).
		Select("i.*, c.name AS cluster_name, c.cluster_id, c.set_name, c.category, c.is_shard, c.shard_type").
		Joins("LEFT JOIN meta_cluster AS c ON i.meta_cluster_id = c.id").
		Where("c.name = ? AND i.is_deleted=0", clusterName).
		Order("i.updated_at DESC").
		Find(&instances).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过集群名称获取实例集群信息失败. %v", err.Error())
	}

	return instances, nil
}

// 通过集群Name获取所有的实例
func (this *InstanceDao) FindInstanceClusterBySetName(setName string) ([]*view_models.InstanceCluster, error) {
	var instances []*view_models.InstanceCluster
	if err := this.db.Table(models.Instance{}.TableNameAndAlias("i")).
		Select("i.*, c.name AS cluster_name, c.cluster_id, c.set_name, c.category, c.is_shard, c.shard_type").
		Joins("LEFT JOIN meta_cluster AS c ON i.meta_cluster_id = c.id").
		Where("c.set_name = ? AND i.is_deleted=0", setName).
		Order("i.updated_at DESC").
		Find(&instances).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过集群set获取实例集群信息失败. %v", err.Error())
	}

	return instances, nil
}

// 获取一个实例信息, 通过 实例id 或 实例名称 或 机器ip, 或 vpcgw_rip
func (this *InstanceDao) GetByInstance(filterInstance *models.Instance) (*models.Instance, error) {
	// 构造where条件
	filterClauses := make([]string, 0, 4)
	if !filterInstance.InstanceId.IsEmpty() {
		filterClauses = append(filterClauses, fmt.Sprintf("instance_id = %#v", filterInstance.InstanceId.String))
	}
	if !filterInstance.InstanceName.IsEmpty() {
		filterClauses = append(filterClauses, fmt.Sprintf("instance_name = %#v", filterInstance.InstanceName.String))
	}
	if !filterInstance.MachineHost.IsEmpty() {
		filterClauses = append(filterClauses, fmt.Sprintf(
			"machine_host = %#v AND port = %#v",
			filterInstance.MachineHost.String, filterInstance.Port.Int64,
		))
	}
	if !filterInstance.VpcgwRip.IsEmpty() {
		filterClauses = append(filterClauses, fmt.Sprintf(
			"vpcgw_rip = %#v AND port = %#v",
			filterInstance.VpcgwRip.String, filterInstance.Port.Int64,
		))
	}

	// 构造or语句
	clause := JoinOrClauses(filterClauses...)

	// 构造and语句
	query := JoinAndClauses(clause, "is_deleted=0")
	logger.M.Debugf("[InstanceDao] GetByInstance. query: %s", query)

	var instance models.Instance
	if err := this.db.Where(query).First(&instance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("获取实例信息失败(实例id/实例名称/机器ip/vpcgw_rip). %v", err.Error())
	}

	return &instance, nil
}

func (this *InstanceDao) GetByInstanceId(instanceId string) (*models.Instance, error) {
	var instance models.Instance
	if err := this.db.Where("instance_id = ? AND is_deleted = 0", instanceId).First(&instance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("通过实例id获取实例信息失败. %v", err.Error())
	}

	return &instance, nil
}

func (this *InstanceDao) GetByInstanceName(instanceName string) (*models.Instance, error) {
	var instance models.Instance
	if err := this.db.Where("instance_name = ? AND is_deleted = 0", instanceName).First(&instance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("通过实例名称获取实例信息失败. %v", err.Error())
	}

	return &instance, nil
}

func (this *InstanceDao) GetByMachineHostAndPort(machineHost string, port int64) (*models.Instance, error) {
	var instance models.Instance
	if err := this.db.Where("machine_host = ? AND port = ? AND is_deleted = 0", machineHost, port).First(&instance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("通过 machine_host 和 port 获取实例信息失败. %v", err.Error())
	}

	return &instance, nil
}

func (this *InstanceDao) GetByVpcgwRipAndPort(vpcgwRip string, port int64) (*models.Instance, error) {
	var instance models.Instance
	if err := this.db.Where("vpcgw_rip = ? AND port = ? AND is_deleted = 0", vpcgwRip, port).First(&instance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("通过 vpcgw_rip 和 port 获取实例信息失败. %v", err.Error())
	}

	return &instance, nil
}

func (this *InstanceDao) GetMasterByMetaClusterId(metaClusterId int64) (*models.Instance, error) {
	var instance models.Instance
	if err := this.db.Where("meta_cluster_id = ? AND role='master' AND is_deleted = 0", metaClusterId).First(&instance).Error; err != nil {
		return nil, fmt.Errorf("通过 meta_cluster_id 获取master实例信息失败. %v", err.Error())
	}

	return &instance, nil
}

func (this *InstanceDao) FindMasterByMetaClusterId(metaClusterId int64) ([]*models.Instance, error) {
	var instances []*models.Instance
	if err := this.db.Where("meta_cluster_id = ? AND role='master' AND is_deleted = 0", metaClusterId).Find(&instances).Error; err != nil {
		return nil, fmt.Errorf("通过 meta_cluster_id 获取所有master实例信息失败. %v", err.Error())
	}

	return instances, nil
}
