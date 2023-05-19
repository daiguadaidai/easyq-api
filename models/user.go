package models

import (
	"github.com/daiguadaidai/easyq-api/types"
)

const (
	RoleDBA   = "dba"
	RoleDev   = "dev"
	RoleGuest = "guest"
)

type User struct {
	ID        types.NullInt64  `json:"id" gorm:"column:id"`
	Username  types.NullString `json:"username" gorm:"column:username;unique;not null;size:50"`
	Password  types.NullString `json:"password" gorm:"column:password;not null;default:'';size:128"`
	Email     types.NullString `json:"email" gorm:"column:email;not null;default:''"`
	Mobile    types.NullString `json:"mobile" gorm:"column:mobile;not null;default:''"`
	NameEn    types.NullString `json:"name_en" gorm:"column:name_en;not null;default:''"`
	NameZh    types.NullString `json:"name_zh" gorm:"column:name_zh;not null;default:''"`
	Role      types.NullString `json:"role" gorm:"column:role"`
	IsDeleted types.NullInt64  `json:"is_deleted" gorm:"column:is_deleted;not null;default:0"`
	UpdatedAt types.NullTime   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt types.NullTime   `json:"created_at" gorm:"column:created_at"`
}

func (User) TableName() string {
	return "user"
}
