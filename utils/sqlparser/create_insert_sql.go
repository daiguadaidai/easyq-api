package sqlparser

import (
	"fmt"
	"github.com/daiguadaidai/parser/ast"
)

func CreateInsertSqlsByRows(insertStmt *ast.InsertStmt, rows [][]interface{}, batchCount int64, endStr string) ([]string, error) {
	insertStmts, err := CreateInsertStmtsByRows(insertStmt, rows, batchCount)
	if err != nil {
		return nil, err
	}

	statements := make([]string, 0, len(insertStmts))
	for _, insertStmt := range insertStmts {
		statement, err := RestoreSql(insertStmt, endStr)
		if err != nil {
			return nil, fmt.Errorf("Insert. %s", err.Error())
		}
		statements = append(statements, statement)
	}

	return statements, nil
}

func CreateInsertStmtsByRows(insertStmt *ast.InsertStmt, rows [][]interface{}, batchCount int64) ([]*ast.InsertStmt, error) {
	// 计算要循环多少次
	rowCount := int64(len(rows))       // 一共有多少条数据
	integer := rowCount / batchCount   // 整数
	remainder := rowCount % batchCount // 余数

	stmtCount := integer // 有多少个语句
	if remainder > 0 {
		stmtCount++
	}
	insertStmts := make([]*ast.InsertStmt, 0, stmtCount)

	var startIndex int64
	for i := int64(0); i < integer; i++ {
		endIndex := startIndex + batchCount
		list := NewRows(rows[startIndex:endIndex])
		newInsertStmt := &ast.InsertStmt{
			IsReplace:      insertStmt.IsReplace,
			IgnoreErr:      insertStmt.IgnoreErr,
			Table:          insertStmt.Table,
			Columns:        insertStmt.Columns,
			Lists:          list,
			Setlist:        insertStmt.Setlist,
			Priority:       insertStmt.Priority,
			OnDuplicate:    insertStmt.OnDuplicate,
			TableHints:     insertStmt.TableHints,
			PartitionNames: insertStmt.PartitionNames,
		}
		insertStmts = append(insertStmts, newInsertStmt)
		startIndex = endIndex
	}

	// 还有剩余的批量没有添加
	if remainder > 0 {
		list := NewRows(rows[startIndex:])
		newInsertStmt := &ast.InsertStmt{
			IsReplace:      insertStmt.IsReplace,
			IgnoreErr:      insertStmt.IgnoreErr,
			Table:          insertStmt.Table,
			Columns:        insertStmt.Columns,
			Lists:          list,
			Setlist:        insertStmt.Setlist,
			Priority:       insertStmt.Priority,
			OnDuplicate:    insertStmt.OnDuplicate,
			TableHints:     insertStmt.TableHints,
			PartitionNames: insertStmt.PartitionNames,
		}
		insertStmts = append(insertStmts, newInsertStmt)
	}

	return insertStmts, nil
}
