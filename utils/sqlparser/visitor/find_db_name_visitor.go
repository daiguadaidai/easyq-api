package visitor

import (
	"github.com/daiguadaidai/parser/ast"
	_ "github.com/daiguadaidai/parser/test_driver"
)

type FindDBNameVisitor struct {
	DBNames []string
}

func NewFindDBNameVisitor() *FindDBNameVisitor {
	findDBNameVisitor := new(FindDBNameVisitor)
	findDBNameVisitor.DBNames = make([]string, 0, 5)

	return findDBNameVisitor
}

func (this *FindDBNameVisitor) Enter(in ast.Node) (out ast.Node, skipChildren bool) {
	switch stmt := in.(type) {
	case *ast.TableName:
		if stmt.Schema.String() != "" {
			this.DBNames = append(this.DBNames, stmt.Schema.String())
		}
	}

	return in, false
}

func (this *FindDBNameVisitor) Leave(in ast.Node) (out ast.Node, ok bool) {
	return in, true
}
