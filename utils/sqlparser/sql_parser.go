package sqlparser

import (
	"fmt"
	"github.com/daiguadaidai/parser"
	"github.com/daiguadaidai/parser/ast"
	"github.com/daiguadaidai/parser/format"
	"github.com/daiguadaidai/parser/model"
	_ "github.com/daiguadaidai/parser/test_driver"
	"regexp"
	"strings"
)

func IsAllDDL(sqlStr string) (bool, []error, error) {
	ps := parser.New()
	stmts, _, err := ps.Parse(sqlStr, "", "")
	if err != nil {
		return false, nil, fmt.Errorf("解析sql失败: %v", err.Error())
	}

	for _, stmt := range stmts {
		switch stmt.(type) {
		case *ast.CreateTableStmt, *ast.AlterTableStmt:
		default:
			return false, nil, nil
		}
	}

	return true, nil, nil
}

func SqlToMulti(sqlStr string) ([]string, error) {
	ps := parser.New()
	stmts, _, err := ps.Parse(sqlStr, "", "")
	if err != nil {
		return nil, fmt.Errorf("解析sql失败: %v", err.Error())
	}

	sqls := make([]string, 0, len(stmts))
	for _, stmt := range stmts {
		sqls = append(sqls, stmt.Text())
	}

	return sqls, nil
}

// 获取所有Alter table表名
func GetAllAlterTableNames(sqlStrs []string) ([]string, error) {
	tableNames := make([]string, 0, len(sqlStrs))

	for _, sqlStr := range sqlStrs {
		ps := parser.New()
		stmts, _, err := ps.Parse(sqlStr, "", "")
		if err != nil {
			return nil, fmt.Errorf("解析sql失败. %v %v", sqlStr, err.Error())
		}

		for _, stmt := range stmts {
			switch stmtNode := stmt.(type) {
			case *ast.AlterTableStmt:
				tableNames = append(tableNames, stmtNode.Table.Name.String())
			case *ast.CreateTableStmt:
			default:
				return nil, fmt.Errorf("只支持 ALTER, CREATE TABLE 语句")
			}
		}
	}

	return tableNames, nil
}

// 获取所有Create table表名
func GetAllCreateTableNames(sqlStrs []string) ([]string, error) {
	tableNames := make([]string, 0, len(sqlStrs))

	for _, sqlStr := range sqlStrs {
		ps := parser.New()
		stmts, _, err := ps.Parse(sqlStr, "", "")
		if err != nil {
			return nil, fmt.Errorf("解析sql失败. %v %v", sqlStr, err.Error())
		}

		for _, stmt := range stmts {
			switch stmtNode := stmt.(type) {
			case *ast.AlterTableStmt:
			case *ast.CreateTableStmt:
				tableNames = append(tableNames, stmtNode.Table.Name.String())
			default:
				return nil, fmt.Errorf("只支持 ALTER, CREATE TABLE 语句")
			}
		}
	}

	return tableNames, nil
}

// 获取DDL表名
func GetDDLTableName(sqlStr string) (string, error) {
	ps := parser.New()
	stmt, err := ps.ParseOneStmt(sqlStr, "", "")
	if err != nil {
		return "", fmt.Errorf("解析sql失败. %v %v", sqlStr, err.Error())
	}

	switch stmtNode := stmt.(type) {
	case *ast.AlterTableStmt:
		return stmtNode.Table.Name.String(), nil
	case *ast.CreateTableStmt:
		return stmtNode.Table.Name.String(), nil
	default:
		return "", fmt.Errorf("只支持 ALTER, CREATE 语句")
	}
}

// 替换所有DDL表名
func ReplaceDDLsTableName(sqlStrs []string, logicRealTableNameMap map[string]string) ([]string, error) {
	newSqls := make([]string, 0, len(sqlStrs))
	for _, sqlStr := range sqlStrs {
		logicTableName, err := GetDDLTableName(sqlStr)
		if err != nil {
			return nil, fmt.Errorf("获取DDL表名失败. %v %v", sqlStr, err.Error())
		}
		realTableName, ok := logicRealTableNameMap[logicTableName]
		if !ok {
			return nil, fmt.Errorf("通过逻辑表名(%v), 无法获取到真实表名(%v) %v", logicTableName, realTableName, sqlStr)
		}

		newSql, err := ReplaceDDLTableName(sqlStr, realTableName)
		if err != nil {
			return nil, fmt.Errorf("重写DDL失败. 将逻辑表名(%v)替换为真实表名(%v) %v", logicTableName, realTableName, sqlStr)
		}
		newSqls = append(newSqls, newSql)
	}

	return newSqls, nil
}

func ReplaceDDLTableNames(sqlStr string, replaceTableNames []string) ([]string, error) {
	newDDLs := make([]string, 0, len(replaceTableNames))
	for _, replaceTableName := range replaceTableNames {
		newDDL, err := ReplaceDDLTableName(sqlStr, replaceTableName)
		if err != nil {
			return nil, err
		}
		newDDLs = append(newDDLs, newDDL)
	}

	return newDDLs, nil
}

func ReplaceDDLTableName(sqlStr string, replaceTableName string) (string, error) {
	ps := parser.New()
	stmt, err := ps.ParseOneStmt(sqlStr, "", "")
	if err != nil {
		return "", fmt.Errorf("解析DDL语句失败: %v %v", sqlStr, err.Error())
	}

	// 新建一个表名结构
	tableNameNode := &ast.TableName{
		Name: model.CIStr{
			L: replaceTableName,
			O: replaceTableName,
		},
	}

	switch stmtNode := stmt.(type) {
	case *ast.CreateTableStmt:
		stmtNode.Table = tableNameNode
	case *ast.AlterTableStmt:
		stmtNode.Table = tableNameNode
	default:
		return "", fmt.Errorf("只支持 ALTER, CREATE 语句")
	}

	restoreSql, err := RestoreSql(stmt, "")
	if err != nil {
		return "", fmt.Errorf("replaceTableName: %v, sql: %v, %v", sqlStr, replaceTableName, err.Error())
	}

	return restoreSql, nil
}

// 重写sql
func RestoreSql(node ast.Node, endStr string) (string, error) {
	// 重写并美化
	var sb strings.Builder
	if err := node.Restore(format.NewRestoreCtx(format.DefaultRestoreFlags, &sb)); err != nil {
		return "", fmt.Errorf("重写SQL出错. %s", err.Error())
	}

	if _, err := sb.WriteString(endStr); err != nil {
		return "", fmt.Errorf("重写sql添加结尾符号出错. %s", err.Error())
	}

	return sb.String(), nil
}

const (
	StmtTypeUnknow = "Unknow"
	StmtTypeSelect = "Select"
	StmtTypeUpdate = "Update"
	StmtTypeDelete = "Delete"
	StmtTypeInsert = "Insert"
	StmtTypeCreate = "Create"
	StmtTypeAlter  = "Alter"
)

func GetStmtType(sqlStr string) (string, error) {
	ps := parser.New()
	stmt, err := ps.ParseOneStmt(sqlStr, "", "")
	if err != nil {
		return "", err
	}

	switch stmt.(type) {
	case *ast.SelectStmt:
		return StmtTypeSelect, nil
	case *ast.UpdateStmt:
		return StmtTypeUpdate, nil
	case *ast.DeleteStmt:
		return StmtTypeDelete, nil
	case *ast.InsertStmt:
		return StmtTypeInsert, nil
	case *ast.AlterTableStmt:
		return StmtTypeAlter, nil
	case *ast.CreateTableStmt, *ast.CreateIndexStmt, *ast.CreateBindingStmt, *ast.CreateDatabaseStmt, *ast.CreateSequenceStmt, *ast.CreateUserStmt, *ast.CreateViewStmt:
		return StmtTypeCreate, nil
	}

	return StmtTypeUnknow, nil
}

func GetStmtType2(sqlStr string) string {
	commentEndPos := GetSQLStmtHearderCommentEndPos(sqlStr)
	newSqlStr := sqlStr[commentEndPos+1:]

	if matched, _ := regexp.MatchString(`(?i)^\s*SELECT`, newSqlStr); matched {
		return StmtTypeSelect
	} else if matched, _ := regexp.MatchString(`(?i)^\s*INSERT`, newSqlStr); matched {
		return StmtTypeInsert
	} else if matched, _ := regexp.MatchString(`(?i)^\s*UPDATE`, newSqlStr); matched {
		return StmtTypeUpdate
	} else if matched, _ := regexp.MatchString(`(?i)^\s*DELETE`, newSqlStr); matched {
		return StmtTypeDelete
	} else if matched, _ := regexp.MatchString(`(?i)^\s*CREATE`, newSqlStr); matched {
		return StmtTypeCreate
	} else if matched, _ := regexp.MatchString(`(?i)^\s*ALTER`, newSqlStr); matched {
		return StmtTypeAlter
	}

	return StmtTypeUnknow
}

func GetSQLStmtHearderComment(stmt string) string {
	var commentBegin bool
	var commentEnd bool
	var meetBeginRod bool
	var meetEndAsterisk bool
	var startContentPos int
	var endContentPos int

	for i, item := range stmt {
		if !commentBegin { // 注释没开始
			if meetBeginRod { // 开始的反斜杆之后必须是 星号 '*' ascii: 42
				if item != 42 {
					return ""
				}
				commentBegin = true     // 设置注释已经开始
				startContentPos = i + 1 // 设置注释内容开始的位点
			} else { // 上一个字符没有碰到反斜杆 '/' ascii: 47
				switch item {
				case 9, 10, 13, 32: // 空白符
					continue
				case 47: // 反斜杆
					meetBeginRod = true //
				default:
					return ""
				}
			}
		} else { // 注释开始, 获取注释内容结束位点
			if meetEndAsterisk { // 碰到星号 '*' ascii: 42 需要检测是否注释结束
				if item == 47 { // 碰到了  */ 注释结束
					endContentPos = i - 1 // 获取注释内容结束位点
					commentEnd = true
					break
				} else if item == 42 { // 还是星号进行下一次字符判断
					continue
				}

				// 星号后面接的不是 '/'
				meetEndAsterisk = false
			} else { // 没有遇到星号
				if item == 42 {
					meetEndAsterisk = true
				}
			}
		}
	}

	if commentEnd {
		return stmt[startContentPos:endContentPos]
	}

	return ""
}

func GetSQLStmtHearderCommentEndPos(stmt string) int {
	var commentBegin bool
	var endCommentPos int = -1
	var meetBeginRod bool
	var meetEndAsterisk bool

	for i, item := range stmt {
		if !commentBegin { // 注释没开始
			if meetBeginRod { // 开始的反斜杆之后必须是 星号 '*' ascii: 42
				if item != 42 {
					return endCommentPos
				}
				commentBegin = true // 设置注释已经开始
			} else { // 上一个字符没有碰到反斜杆 '/' ascii: 47
				switch item {
				case 9, 10, 13, 32: // 空白符
					continue
				case 47: // 反斜杆
					meetBeginRod = true //
				default:
					return endCommentPos
				}
			}
		} else { // 注释开始, 获取注释内容结束位点
			if meetEndAsterisk { // 碰到星号 '*' ascii: 42 需要检测是否注释结束
				if item == 47 { // 碰到了  */ 注释结束
					endCommentPos = i
					break
				} else if item == 42 { // 还是星号进行下一次字符判断
					continue
				}

				// 星号后面接的不是 '/'
				meetEndAsterisk = false
			} else { // 没有遇到星号
				if item == 42 {
					meetEndAsterisk = true
				}
			}
		}
	}

	return endCommentPos
}

func ExistsCreateTableStmt(sqlStr string) (bool, error) {
	ps := parser.New()
	stmts, _, err := ps.Parse(sqlStr, "", "")
	if err != nil {
		return false, err
	}

	for _, stmt := range stmts {
		switch stmt.(type) {
		case *ast.CreateTableStmt:
			return true, nil
		}
	}

	return false, nil
}

func IsCreateTableStmt(sqlStr string) (bool, error) {
	ps := parser.New()
	stmt, err := ps.ParseOneStmt(sqlStr, "", "")
	if err != nil {
		return false, err
	}

	switch stmt.(type) {
	case *ast.CreateTableStmt:
		return true, nil
	}

	return false, nil
}

func NormalizeDigest(sqlStr string) (string, string) {
	nor, dig :=parser.NormalizeDigest(strings.TrimSpace(sqlStr))

	return nor, dig.String()
}
