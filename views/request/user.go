package request

import "github.com/daiguadaidai/easyq-api/types"

type UserFindRequest struct {
	Pager
	Keyword types.NullString `json:"keyword" form:"keyword"`
}

type UserRegisterRequest struct {
	Username types.NullString `json:"username" form:"username" binding:"required"`
	Password types.NullString `json:"password" form:"password"`
	Email    types.NullString `json:"email" form:"email"`
	NameEn   types.NullString `json:"name_en" form:"name_en"`
	NameZh   types.NullString `json:"name_zh" form:"name_zh"`
}

type UserLoginRequest struct {
	Username types.NullString `json:"username" form:"username" binding:"required"`
	Password types.NullString `json:"password" form:"password"`
}

type UserEditByIdRequest struct {
	ID       types.NullInt64  `json:"id" form:"id" binding:"required"`
	Username types.NullString `json:"username" form:"username"`
	Password types.NullString `json:"password" form:"password"`
	Email    types.NullString `json:"email" form:"email"`
	Mobile   types.NullString `json:"mobile" form:"mobile"`
	NameEn   types.NullString `json:"name_en" form:"name_en"`
	NameZh   types.NullString `json:"name_zh" form:"name_zh"`
	Role     types.NullString `json:"role" form:"role"`
}

type UserDeleteByIdRequest struct {
	ID types.NullInt64 `json:"id" form:"id" binding:"required"`
}
