package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/types"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/daiguadaidai/easyq-api/views/response"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type UserController struct {
	ctx *contexts.GlobalContext
}

func NewUserController(ctx *contexts.GlobalContext) *UserController {
	return &UserController{ctx: ctx}
}

// 注册
func (this *UserController) Register(req *request.UserRegisterRequest) (*models.User, error) {
	userDao := dao.NewUserDao(this.ctx.EasyqDB)

	user := new(models.User)
	utils.CopyStruct(req, user)
	user.Role = types.NewNullString(models.RoleDev, true)

	if err := userDao.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// 登录
func (this *UserController) Login(req *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	// 获取数据库中的用户
	user, err := dao.NewUserDao(this.ctx.EasyqDB).GetByUsername(req.Username.String)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("用户名 %v 不存在", req.Username.String)
	}
	if req.Password.String != user.Password.String {
		return nil, fmt.Errorf("用户 %v 存在, 但是登陆密码输入错误, 密码: %v", req.Username.String, req.Password.String)
	}

	// 生成登入 token
	token, err := this.generateToken(user)
	if err != nil {
		return nil, fmt.Errorf("生成token失败, username: %s, role: %s", user.Username.String, user.Password.String)
	}

	var resp response.UserLoginResponse
	utils.CopyStruct(user, &resp)
	resp.Token = types.NewNullString(token, true)

	return &resp, nil
}

// 获取用户信息
func (this *UserController) Find(req *request.UserFindRequest) ([]*models.User, error) {
	users, err := dao.NewUserDao(this.ctx.EasyqDB).FindByKeyword(strings.TrimSpace(req.Keyword.String), req.Offset(), req.Limit())
	if err != nil {
		return nil, err
	}

	return users, nil
}

// 获取user count
func (this *UserController) Count(req *request.UserFindRequest) (int, error) {
	return dao.NewUserDao(this.ctx.EasyqDB).CountByKeyword(strings.TrimSpace(req.Keyword.String))
}

// 通过id编辑
func (this *UserController) EditById(req *request.UserEditByIdRequest) error {
	var user models.User
	utils.CopyStruct(req, &user)

	return dao.NewUserDao(this.ctx.EasyqDB).UpdateById(&user)
}

// 通过id删除
func (this *UserController) DeleteById(req *request.UserDeleteByIdRequest) error {
	var user models.User
	utils.CopyStruct(req, &user)
	user.IsDeleted = types.NewNullInt64(1, false)

	return dao.NewUserDao(this.ctx.EasyqDB).UpdateById(&user)
}

// 通过id删除
func (this *UserController) All() ([]*models.User, error) {
	return dao.NewUserDao(this.ctx.EasyqDB).All()
}

// 生成令牌
func (this *UserController) generateToken(user *models.User) (string, error) {
	j := utils.NewJWT()

	claims := utils.CustomClaims{
		Username: user.Username.String,
		Password: user.Password.String,
		Role:     user.Role.String,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),                               // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + this.ctx.Cfg.ApiConfig.TokenExpire), // 过期时间1天
			Issuer:    "HH",                                                          // 签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (this *UserController) GetUserByUsername(username string) (*models.User, error) {
	return dao.NewUserDao(this.ctx.EasyqDB).GetByUsernameEmptyError(username)
}
