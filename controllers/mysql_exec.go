package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/utils"
	"github.com/daiguadaidai/easyq-api/utils/sqlparser"
	"github.com/daiguadaidai/easyq-api/views/request"
	"github.com/gin-gonic/gin"
)

type MysqlExecController struct {
	ctx *contexts.GlobalContext
}

func NewMysqlExecController(ctx *contexts.GlobalContext) *MysqlExecController {
	return &MysqlExecController{ctx: ctx}
}

func (this *MysqlExecController) ExecSql(c *gin.Context, req *request.MysqlExecSqlRequest) error {
	// 检测权限是否存在
	priv, err := helper.CheckMysqlPrivById(c, this.ctx, req.PrivId.Int64)
	if err != nil {
		return fmt.Errorf("检测用户数据库权限出错. %v", err.Error())
	}

	// 获取解析后的sql节点
	stmtNode, err := sqlparser.ParseOneStmt(req.Query.String)
	if err != nil {
		return fmt.Errorf("解析sql语句出错. %v", err.Error())
	}

	// 获取查询类型
	stmtType := sqlparser.GetStmtType(stmtNode)
	if stmtType != sqlparser.StmtTypeExplain && stmtType != sqlparser.StmtTypeSelect {
		return fmt.Errorf("只允许 SELECT 和 EXPLAIN 语句")
	}

	// 获取数据库查询的所有数据库
	dbNames := sqlparser.FindDBNamesByStmtNode(stmtNode)
	// 将公共库的权限过滤掉
	dbNames = helper.FilterIgnoreDatabases(dbNames)

	// 检测这些库的权限是否都有

	// SELECT 语句查看是否有LIMIT, 没有LIMIT 自动添加, 默认 LIMIT 2000

	// 获取语句前缀 注释

	// 从新生成语句 并且 添加上前前缀注释

	logger.M.Infof(utils.ToJsonStr(dbNames))
	logger.M.Infof(utils.ToJsonStr(priv))

	return nil
}
