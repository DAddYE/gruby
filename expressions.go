package gruby

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
)

func (p *Printer) expr(exp ast.Expr) {
	p.print(exp.Pos())
	switch x := exp.(type) {
	case *ast.BadExpr:
		fmt.Fprintf(os.Stderr, "print: got a bad expression %v", x)
		panic("gruby/nodes expr")

	// case *ast.Ellipsis:
	case *ast.BasicLit:
		p.print(x)
	// case *ast.FuncLit:
	// case *ast.CompositeLit:

	// case *ast.ParenExpr:
	// case *ast.SelectorExpr:
	// case *ast.IndexExpr:
	// case *ast.SliceExpr:
	// case *ast.StarExpr:
	// case *ast.TypeAssertExpr:
	// case *ast.CallExpr:
	// case *ast.UnaryExpr:
	// case *ast.BinaryExpr:
	// case *ast.KeyValueExpr:

	// case *ast.ArrayType:
	// case *ast.StructType:
	// case *ast.FuncType:
	// case *ast.InterfaceType:
	// case *ast.MapType:
	// case *ast.ChanType:

	case *ast.Ident:
		p.print(x.Name)
	default:
		fmt.Fprintf(os.Stderr, "print: not implemented yet (%T) %+v", x, x)
		panic("gruby/nodes expr")
	}
}

func (p *Printer) exprList(exprs []ast.Expr) {
	size := len(exprs)
	for i, expr := range exprs {
		p.expr(expr)
		if i != size-1 {
			p.print(blank, token.COMMA)
		}
	}
}
