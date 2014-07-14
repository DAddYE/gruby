package gruby

import (
	"go/ast"
	"go/token"
)

func (p *Printer) expr1(exp ast.Expr, prec1, depth int) {
	p.print(exp.Pos())
	switch x := exp.(type) {
	case *ast.BadExpr:
		panicf("print: got a bad expression %v", x)

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
	case *ast.BinaryExpr:
		if depth < 1 {
			panicf("depth < 1:", depth)
			depth = 1
		}
		p.binaryExpr(x, prec1, cutoff(x, depth), depth)

	// case *ast.KeyValueExpr:

	// case *ast.ArrayType:
	case *ast.StructType:
		p.print(STRUCT, DOT, NEW, token.LPAREN)
		p.fieldList(x.Fields)
		p.print(token.RPAREN)
	// case *ast.FuncType:
	// case *ast.InterfaceType:
	// case *ast.MapType:
	// case *ast.ChanType:

	case *ast.Ident:
		if p.context == inConst && x.Name == "iota" {
			p.expr(&ast.BasicLit{Value: "0", Kind: token.INT})
			p.context = inIota
		} else {
			p.print(x.Name)
		}
	default:
		panicf("print: not implemented yet (%T) %+v", x, x)
	}
}

func (p *Printer) fieldList(f *ast.FieldList) {
	if len(f.List) == 0 {
		return
	}
	names := []*ast.Ident{}
	for _, field := range f.List {
		if len(field.Names) == 0 {
			// names = append(names, field.Type)
			continue
		}
		for _, name := range field.Names {
			names = append(names, name)
		}
	}
	p.identListPrefixed(names, ":")
}

func (p *Printer) expr0(x ast.Expr, depth int) {
	p.expr1(x, token.LowestPrec, depth)
}

func (p *Printer) expr(x ast.Expr) {
	const depth = 1
	p.expr1(x, token.LowestPrec, depth)
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

func (p *Printer) binaryExpr(x *ast.BinaryExpr, prec1, cutoff, depth int) {
	prec := x.Op.Precedence()
	if prec < prec1 {
		p.print(token.LPAREN)
		p.expr0(x, reduceDepth(depth)) // parentheses undo one level of depth
		p.print(token.RPAREN)
		return
	}

	printBlank := prec < cutoff

	ws := indent
	p.expr1(x.X, prec, depth+diffPrec(x.X, prec))
	if printBlank {
		p.print(blank)
	}
	p.print(x.OpPos, x.Op)
	if printBlank {
		p.print(blank)
	}
	p.expr1(x.Y, prec+1, depth+1)
	if ws == ignore {
		p.print(dedent)
	}
}

func isBinary(expr ast.Expr) bool {
	_, ok := expr.(*ast.BinaryExpr)
	return ok
}

func walkBinary(e *ast.BinaryExpr) (has4, has5 bool, maxProblem int) {
	switch e.Op.Precedence() {
	case 4:
		has4 = true
	case 5:
		has5 = true
	}

	switch l := e.X.(type) {
	case *ast.BinaryExpr:
		if l.Op.Precedence() < e.Op.Precedence() {
			// parens will be inserted.
			// pretend this is an *ast.ParenExpr and do nothing.
			break
		}
		h4, h5, mp := walkBinary(l)
		has4 = has4 || h4
		has5 = has5 || h5
		if maxProblem < mp {
			maxProblem = mp
		}
	}

	switch r := e.Y.(type) {
	case *ast.BinaryExpr:
		if r.Op.Precedence() <= e.Op.Precedence() {
			// parens will be inserted.
			// pretend this is an *ast.ParenExpr and do nothing.
			break
		}
		h4, h5, mp := walkBinary(r)
		has4 = has4 || h4
		has5 = has5 || h5
		if maxProblem < mp {
			maxProblem = mp
		}

	case *ast.StarExpr:
		if e.Op == token.QUO { // `*/`
			maxProblem = 5
		}

	case *ast.UnaryExpr:
		switch e.Op.String() + r.Op.String() {
		case "/*", "&&", "&^":
			maxProblem = 5
		case "++", "--":
			if maxProblem < 4 {
				maxProblem = 4
			}
		}
	}
	return
}
