package sqlparser

import (
	"fmt"
	"github.com/daiguadaidai/easyq-api/utils"
	"testing"
)

func Test_SqlToMulti(t *testing.T) {
	sqlStr := `
ALTER TABLE emp ADD COLUMN col_01 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '测试';
ALTER TABLE emp ADD COLUMN col_02 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '测试';
ALTER TABLE emp ADD COLUMN col_03 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '测试';
ALTER TABLE emp ADD INDEX idx_col_01(col_01);
ALTER TABLE emp DROP COLUMN col_test;
`
	sqls, err := SqlToMulti(sqlStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("%v", utils.ToJsonStrPretty(sqls))
}

func Test_ReplaceDDLTableName(t *testing.T) {
	sqlStr := "ALTER TABLE emp ADD COLUMN col_01 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '测试';"
	replaceTableName := "employees"

	newSql, err := ReplaceDDLTableName(sqlStr, replaceTableName)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(newSql)
}

func Test_ReplaceDDLsTableName(t *testing.T) {
	sqlStr := `
ALTER TABLE emp
ADD COLUMN col_01 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '测试';
ALTER TABLE user ADD COLUMN col_02 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '测试';
ALTER TABLE emp ADD COLUMN col_03 VARCHAR(20) NOT NULL DEFAULT '' COMMENT '测试';
ALTER TABLE emp ADD INDEX idx_col_01(col_01);
ALTER TABLE emp DROP COLUMN col_test;
`
	sqls, err := SqlToMulti(sqlStr)
	if err != nil {
		t.Fatal(err.Error())
	}
	logicRealTableNameMap := map[string]string{
		"emp":  "employees",
		"user": "user_02",
	}

	newSqls, err := ReplaceDDLsTableName(sqls, logicRealTableNameMap)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(utils.ToJsonStrPretty(newSqls))
}

func TestGetStmtType2(t *testing.T) {
	sqlStr := `/**/SELECT 1`
	stmtType := GetStmtType2(sqlStr)
	fmt.Println(stmtType)

	sqlStr = `/*  */SElect 1`
	stmtType = GetStmtType2(sqlStr)
	fmt.Println(stmtType)

	sqlStr = `/* */ SELECT 1`
	stmtType = GetStmtType2(sqlStr)
	fmt.Println(stmtType)

	sqlStr = `  SeLECT 1`
	stmtType = GetStmtType2(sqlStr)
	fmt.Println(stmtType)

	sqlStr = `/* */ INSERT 1`
	stmtType = GetStmtType2(sqlStr)
	fmt.Println(stmtType)
}

func Test_GetSQLStmtHearderCommentEndPos(t *testing.T) {
	sqlStr := `/**/SELECT 1`
	endPos := GetSQLStmtHearderCommentEndPos(sqlStr)
	fmt.Printf("pos: %v. sql: %v. 取消注释sql: %v\n", endPos, sqlStr, sqlStr[endPos+1:])

	sqlStr = `/*  */SELECT 1`
	endPos = GetSQLStmtHearderCommentEndPos(sqlStr)
	fmt.Printf("pos: %v. sql: %v. 取消注释sql: %v\n", endPos, sqlStr, sqlStr[endPos+1:])

	sqlStr = `/* */ SELECT 1`
	endPos = GetSQLStmtHearderCommentEndPos(sqlStr)
	fmt.Printf("pos: %v. sql: %v. 取消注释sql: %v\n", endPos, sqlStr, sqlStr[endPos+1:])

	sqlStr = `SELECT 1`
	endPos = GetSQLStmtHearderCommentEndPos(sqlStr)
	fmt.Printf("pos: %v. sql: %v. 取消注释sql: %v\n", endPos, sqlStr, sqlStr[endPos+1:])
}
