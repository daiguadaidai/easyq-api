package response

import (
	"github.com/daiguadaidai/easyq-api/types"
)

type UserLoginResponse struct {
	ID        types.NullInt64  `json:"id"`
	Username  types.NullString `json:"username"`
	Email     types.NullString `json:"email"`
	Mobile    types.NullString `json:"mobile"`
	NameEn    types.NullString `json:"name_en"`
	NameZh    types.NullString `json:"name_zh"`
	Role      types.NullString `json:"role"`
	IsDeleted types.NullInt64  `json:"is_deleted"`
	UpdatedAt types.NullTime   `json:"updated_at"`
	CreatedAt types.NullTime   `json:"created_at"`
	Token     types.NullString `json:"token"`
}
