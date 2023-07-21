package sqlparser

import (
	"fmt"
	"github.com/daiguadaidai/parser"
	"github.com/daiguadaidai/parser/ast"
	"strings"
)

func NewRow(fields []interface{}) []ast.ExprNode {
	row := make([]ast.ExprNode, 0, len(fields))

	for _, field := range fields {
		row = append(row, ast.NewValueExpr(field, "", ""))
	}

	return row
}

func NewRows(dataz [][]interface{}) [][]ast.ExprNode {
	rows := make([][]ast.ExprNode, 0, len(dataz))
	for _, fields := range dataz {
		row := NewRow(fields)
		rows = append(rows, row)
	}

	return rows
}

func CreateInsertStmtTemplate(dbName, tableName string, columnNames []string) (*ast.InsertStmt, error) {
	dbTableName := ""
	if dbName == "" {
		dbTableName = tableName
	} else {
		dbTableName = fmt.Sprintf("`%v`.`%v`", dbName, tableName)
	}

	columnNamesStr := strings.Join(columnNames, ", ")

	insertStr := fmt.Sprintf("INSERT INTO %v(%v) VALUES()", dbTableName, columnNamesStr)

	ps := parser.New()

	stmtNode, err := ps.ParseOneStmt(insertStr, "", "")
	if err != nil {
		return nil, fmt.Errorf("解析insert模板sql出错. %v", err.Error())
	}

	return stmtNode.(*ast.InsertStmt), nil
}
