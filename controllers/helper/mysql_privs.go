package helper

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/dao"
	"github.com/daiguadaidai/easyq-api/middlewares"
	"github.com/daiguadaidai/easyq-api/models"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/views/response"
	"github.com/gin-gonic/gin"
)

// 申请权限去重
func UniqueMysqlPrivApplys(applys []*models.MysqlDBPrivApply) []*models.MysqlDBPrivApply {
	privMap := make(map[string]*models.MysqlDBPrivApply)
	uniqueApplys := make([]*models.MysqlDBPrivApply, 0, len(applys))

	for _, apply := range applys {
		key := fmt.Sprintf("%v#%v#v", apply.MetaClusterId.Int64, apply.DBName.String)
		if _, ok := privMap[key]; !ok {
			privMap[key] = apply
			uniqueApplys = append(uniqueApplys, apply)
		}
	}

	return uniqueApplys
}

func MysqlPrivsToTree(privs []*models.MysqlDBPriv) []*response.MysqlPrivsTreeResponse {
	tree := make([]*response.MysqlPrivsTreeResponse, 0, len(privs))
	for _, priv := range privs {
		var p response.MysqlPrivsTreeResponse
		utils.CopyStruct(priv, &p)
		tree = append(tree, &p)
	}

	return tree
}

func CheckMysqlPriv(c *gin.Context, ctx *contexts.GlobalContext, db_name string, meta_cluster_id int64) (*models.MysqlDBPriv, error) {
	// 获取用户信息
	clainms, err := middlewares.GetClaims(c)
	if err != nil {
		return nil, fmt.Errorf("通过token获取用户信息出错: %v", err)
	}

	// 获取数据库用户信息
	user, err := dao.NewUserDao(ctx.EasyqDB).GetByUsername(clainms.Username)
	if err != nil {
		return nil, fmt.Errorf("通过token解析出的用户获取数据库用户出错. username: %v. %v", clainms.Username, err)
	}

	// 通过集群id和数据库名和用户获取数据库权限, 判断用户是否有权限
	priv, err := dao.NewMysqlDBPrivDao(ctx.EasyqDB).GetByUsernameClusterDB(user.Username.String, meta_cluster_id, db_name)
	if err != nil {
		return nil, fmt.Errorf("获取数据库权限出错. username: %v, meta_cluster_id: %v, db_name: %v, %v", user.Username.String, meta_cluster_id, db_name, err.Error())
	}

	if priv == nil {
		return nil, fmt.Errorf("用户没有该数据库查询权限, username: %v, meta_cluster_id: %v, db_name: %v", user.Username.String, meta_cluster_id, db_name)
	}

	return priv, nil
}

func CheckMysqlPrivById(c *gin.Context, ctx *contexts.GlobalContext, privId int64) (*models.MysqlDBPriv, error) {
	// 获取用户信息
	clainms, err := middlewares.GetClaims(c)
	if err != nil {
		return nil, fmt.Errorf("通过token获取用户信息出错: %v", err)
	}

	// 获取数据库用户信息
	user, err := dao.NewUserDao(ctx.EasyqDB).GetByUsername(clainms.Username)
	if err != nil {
		return nil, fmt.Errorf("通过token解析出的用户获取数据库用户出错. username: %v. %v", clainms.Username, err)
	}

	// 通过id获取数据库权限, 判断用户是否有权限
	priv, err := dao.NewMysqlDBPrivDao(ctx.EasyqDB).GetById(privId)
	if err != nil {
		return nil, fmt.Errorf("获取数据库权限出错. 权限id: %v, %v", privId, err.Error())
	}

	if priv == nil {
		return nil, fmt.Errorf("用户没有该数据库查询权限, 权限id: %v", privId)
	}

	if user.Username.String != priv.Username.String {
		return nil, fmt.Errorf("用户没有该数据库查询权限, username: %v, meta_cluster_id: %v, db_name: %v, cluster_name: %v", user.Username.String, priv.MetaClusterId.Int64, priv.DBName.String, priv.ClusterName.String)
	}

	return priv, nil
}
