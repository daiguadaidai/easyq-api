package dao

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/models"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

// 插入用户，先检查是否存在用户，如果没有则存入
func (this *UserDao) Create(user *models.User) error {
	if user.Username.IsEmpty() {
		return fmt.Errorf("用户名不能为空")
	}

	newUser, err := this.GetByUsername(user.Username.String)
	if err != nil {
		return err
	}
	if newUser != nil {
		return fmt.Errorf("用户 %s 已经存在", user.Username.String)
	}

	return this.db.Create(user).Error
}

func (this *UserDao) UpdateById(user *models.User) error {
	if user.ID.Int64 <= 0 {
		return fmt.Errorf("用户id不能<0")
	}

	return this.db.Model(user).Omit("id").Updates(user).Error
}

func (this *UserDao) getFindByKeywordQuery(keyword string) string {
	// 获取 LIKE OR LIKE WHERE 语句
	likeClauses := GetLikeClausesByKeyWords(keyword, "username", "name_en", "name_zh")
	likeClauseStr := JoinOrClauses(likeClauses...)

	// 获取其他语句
	otherClauseStr := "is_deleted=0"

	// LIKE 和 其他语句合并
	return JoinAndClauses(likeClauseStr, otherClauseStr)
}

func (this *UserDao) FindByKeyword(keyword string, offset, limit int) ([]*models.User, error) {
	query := this.getFindByKeywordQuery(keyword)
	logger.M.Debugf("[UserDao] FindByKeyword. query: %s", query)

	var users []*models.User
	if err := this.db.Where(query).Offset(offset).Limit(limit).Order("updated_at DESC").Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过keyword获取用户信息失败. %v", err.Error())
	}

	return users, nil
}

func (this *UserDao) CountByKeyword(keyword string) (int, error) {
	query := this.getFindByKeywordQuery(keyword)
	logger.M.Debugf("[UserDao] CountByKeyword. query: %s", query)

	var cnt int64
	if err := this.db.Model(&models.User{}).Where(query).Count(&cnt).Error; err != nil {
		return 0, fmt.Errorf("通过keyword获取用户信息数失败. %v", err.Error())
	}

	return int(cnt), nil
}

func (this *UserDao) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := this.db.Where("username = ? AND is_deleted=0", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过username获取用户信息失败. %v", err.Error())
	}

	return &user, nil
}

func (this *UserDao) GetByUsernameEmptyError(username string) (*models.User, error) {
	user, err := this.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("通过用户名没有获取到用户信息: username: %v", username)
	}

	return user, nil
}

func (this *UserDao) GetByUsernameAndPassword(username, password string) (*models.User, error) {
	var user models.User
	if err := this.db.Where("username = ? AND password = ? AND is_deleted=0", username, password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("通过username, password获取用户信息失败. %v", err.Error())
	}

	return &user, nil
}

// 获取所有用户
func (this *UserDao) All() ([]*models.User, error) {
	var users []*models.User
	if err := this.db.Where("is_deleted=0").Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("获取所有用户信息失败. %v", err.Error())
	}

	return users, nil
}
