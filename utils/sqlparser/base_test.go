package sqlparser

import (
	"fmt"
	"testing"
)

func TestCreateInsertStmtTemplate(t *testing.T) {
	columnNames := []string{"col_1", "col_2", "col_3"}
	insertStmt, err := CreateInsertStmtTemplate("easyq", "emp", columnNames)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(insertStmt)
}
