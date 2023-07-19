package controllers

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/contexts"
	"github.com/daiguadaidai/easyq-api/controllers/helper"
	"github.com/daiguadaidai/easyq-api/logger"
	"github.com/daiguadaidai/easyq-api/models/view_models"
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

func (this *MysqlExecController) ExecSql(c *gin.Context, req *request.MysqlExecSqlRequest) (*view_models.MysqlExecResult, error) {
	// 检测权限是否存在
	priv, err := helper.CheckMysqlPrivById(c, this.ctx, req.PrivId.Int64)
	if err != nil {
		return nil, fmt.Errorf("检测用户数据库权限出错. %v", err.Error())
	}

	// 获取解析后的sql节点
	stmtNode, err := sqlparser.ParseOneStmt(req.Query.String)
	if err != nil {
		return nil, fmt.Errorf("解析sql语句出错. %v", err.Error())
	}

	// 获取查询类型
	stmtType := sqlparser.GetStmtType(stmtNode)
	if stmtType != sqlparser.StmtTypeExplain && stmtType != sqlparser.StmtTypeSelect {
		return nil, fmt.Errorf("只允许 SELECT 和 EXPLAIN 语句")
	}

	// 获取数据库查询的所有数据库
	dbNames := sqlparser.FindDBNamesByStmtNode(stmtNode)
	// 将公共库的权限过滤掉
	dbNames = helper.FilterIgnoreDatabases(dbNames)

	// 检测这些库的权限是否都有
	if err := helper.CheckExecDBPriv(this.ctx, dbNames, priv.MetaClusterId.Int64, priv.Username.String); err != nil {
		return nil, fmt.Errorf("检测执行sql需要到所有数据库权限失败. %v", err.Error())
	}

	// SELECT 语句查看是否有LIMIT, 没有LIMIT 自动添加, 默认 LIMIT 2000
	newNodeStmt := sqlparser.ResetSelectLimitAndGet(stmtNode, this.ctx.Cfg.ExecConfig.ExecMysqlSelectLimit)

	// 获取语句前缀 注释
	prefixStr := sqlparser.GetSQLStmtHearderComment(stmtNode.Text())

	// 从新生成语句
	newQuery, err := sqlparser.RestoreSql(newNodeStmt, "")
	if err != nil {
		return nil, fmt.Errorf("重写执行sql出错. %v", err.Error())
	}
	// 添加上前前缀注释
	newQuery = fmt.Sprintf("%v%v", prefixStr, newQuery)
	logger.M.Infof("用户: %v, 数据库: %v, 集群id: %v. \n重写前sql: %v\n重写后sql: %v\n", priv.Username.String, priv.DBName.String, priv.MetaClusterId.Int64, stmtNode.Text(), newQuery)

	// 开始执行sql
	rows, columns, err := helper.StartExecSingleMysqlSql(this.ctx, priv, newQuery)
	if err != nil {
		return nil, err
	}

	rs := &view_models.MysqlExecResult{
		ExecSql:     newQuery,
		ColumnNames: columns,
		Rows:        rows,
	}

	return rs, nil
}
