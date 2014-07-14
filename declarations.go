package gruby

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strconv"
)

func (p *Printer) decl(decl ast.Decl) {
	switch x := decl.(type) {
	case *ast.GenDecl:
		p.genDecl(x)
		p.print(newline)
	case *ast.FuncDecl:
		p.funcDecl(x)
	default:
		fmt.Fprintf(os.Stderr, "print: not implemented yet %+v", x)
		panic("gruby/nodes decl")
	}
}

func (p *Printer) declList(list []ast.Decl) {
	for _, decl := range list {
		p.decl(decl)
		p.print(newline)
	}
}

// Declarations
func (p *Printer) genDecl(d *ast.GenDecl) {
	p.setComment(d.Doc)
	p.print(d.Pos())

	for i, spec := range d.Specs {
		switch d.Tok {
		case token.CONST:
			p.context = inConst
			valueSpec := spec.(*ast.ValueSpec)
			exprList := make([]ast.Expr, len(valueSpec.Names))
			copy(exprList, valueSpec.Values)

			for i, ident := range valueSpec.Names {
				p.print(ident.Pos())
				p.print(constantize(ident.Name))
				if i+1 < len(valueSpec.Names) {
					p.print(token.COMMA, blank)
				}
			}
			p.print(blank, token.ASSIGN, blank)

			for _, expr := range exprList {
				if expr != nil {
					p.expr(expr)
				} else if p.context == inIota {
					p.print(strconv.Itoa(i))
				} else {
					p.print(NIL)
				}
				if i+1 < len(exprList) {
					p.print(token.COMMA, blank)
				}
			}
		default:
			p.spec(spec)
		}

		p.print(newline)
	}

	p.context = inDefault
}

func (p *Printer) specClass(spec *ast.TypeSpec, funcs []*ast.FuncDecl) {
	p.print(CLASS, blank, spec.Name)
	t := spec.Type.(*ast.Ident)
	p.print(blank, INHERIT, blank)
	if class, found := goTypeToRuby[t.Name]; found {
		p.print(class)
	} else {
		p.print(classify(t.Name))
	}
	p.print(indent)
	for _, fn := range funcs {
		p.funcDecl(fn)
		p.print(newline)
	}
	p.print(dedent, END)
	p.print(newline)
}

func (p *Printer) spec(spec ast.Spec) {
	switch s := spec.(type) {
	case *ast.ImportSpec:
		p.setComment(s.Doc)
		p.print(REQUIRE, blank)
		if s.Name != nil {
			p.expr(s.Name)
			p.print(blank)
		}
		p.expr(s.Path)
		p.setComment(s.Comment)
		p.print(s.EndPos)

	case *ast.ValueSpec:
		if s.Values == nil {
			break
		}
		p.setComment(s.Doc)
		p.identList(s.Names)
		p.print(blank, token.ASSIGN, blank)
		p.exprList(s.Values)
		p.setComment(s.Comment)

	case *ast.TypeSpec:
		p.setComment(s.Doc)
		p.print(CLASS, blank)
		p.expr(s.Name)
		p.print(blank, INHERIT, blank)
		p.expr(s.Type)
		// t := s.Type.(*ast.Ident)
		// if class, found := goTypeToRuby[t.Name]; found {
		// 	p.print(class)
		// } else {
		// 	p.print(classify(t.Name))
		// }
		p.print(SEMI, blank, END)
		p.setComment(s.Comment)

	default:
		panic("unreachable")
	}
}

func (p *Printer) funcDecl(d *ast.FuncDecl) {
	p.setComment(d.Doc)
	nname := underscore(d.Name.Name)
	p.print(d.Pos(), DEF, blank)
	p.print(nname)
	p.signature(d.Type.Params, d.Type.Results)
	p.print(indent)
	p.stmtList(d.Body.List)
	p.print(dedent, END)
	if isPrivate(d.Name) {
		p.print(newline)
		p.print(PRIVATE, blank, token.COLON, nname)
	}
	p.print(newline)
}

func (p *Printer) parameters(fields *ast.FieldList) {
	p.print(fields.Opening, token.LPAREN)
	for _, par := range fields.List {
		if len(par.Names) == 0 {
			panic("you must provide some parameters")
		}
		p.identList(par.Names)
	}
	p.print(fields.Closing, token.RPAREN)
}

func (p *Printer) signature(params, result *ast.FieldList) {
	if params != nil {
		p.parameters(params)
	}
}

func (p *Printer) identList(list []*ast.Ident) {
	l := len(list) - 1
	for i, x := range list {
		p.print(x)
		if i < l {
			p.print(token.COMMA, blank)
		}
	}
}

func (p *Printer) identListPrefixed(list []*ast.Ident, prefix string) {
	l := len(list) - 1
	for i, x := range list {
		p.print(prefix)
		p.print(x)
		if i < l {
			p.print(token.COMMA, blank)
		}
	}
}
