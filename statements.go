package gruby

import (
	"fmt"
	"go/ast"
	"os"
)

func (p *Printer) stmt(stmt ast.Stmt) {
	switch x := stmt.(type) {
	// case *ast.BadStmt:
	// case *ast.DeclStmt:
	// case *ast.EmptyStmt:
	// case *ast.LabeledStmt:
	// case *ast.ExprStmt:
	// case *ast.SendStmt:
	// case *ast.IncDecStmt:
	// case *ast.AssignStmt:
	// case *ast.GoStmt:
	// case *ast.DeferStmt:
	// case *ast.ReturnStmt:
	// case *ast.BranchStmt:
	// case *ast.BlockStmt:
	// case *ast.IfStmt:
	// case *ast.CaseClause:
	// case *ast.SwitchStmt:
	// case *ast.TypeSwitchStmt:
	// case *ast.CommClause:
	// case *ast.SelectStmt:
	// case *ast.ForStmt:
	// case *ast.RangeStmt:
	default:
		fmt.Fprintf(os.Stderr, "print: not implemented yet %+v", x)
		panic("gruby/nodes stmt")
	}
}

func (p *Printer) stmtList(list []ast.Stmt) {

}
