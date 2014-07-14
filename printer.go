package gruby

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"
)

type whiteSpace byte

const (
	ignore  = whiteSpace(0)
	blank   = whiteSpace(' ')
	newline = whiteSpace('\n')
	indent  = whiteSpace('>')
	dedent  = whiteSpace('<')
)

type context int

const (
	inDefault context = iota
	inConst
	inIota
)

// This represent the source of destination
type Printer struct {
	fset       *token.FileSet
	output     []byte // raw output
	indent     uint   // indent level
	indentSize uint   // size of the indentation level
	indentCh   rune   // char to use as indent
	context    context

	// Positions
	pos  token.Position // position in ast
	out  token.Position // position in output
	last token.Position // last position
}

func NewPrinter(fset *token.FileSet) *Printer {
	return &Printer{
		fset:       fset,
		indentSize: 2,
		indentCh:   ' ',
		pos:        token.Position{Line: 1, Column: 1},
		out:        token.Position{Line: 1, Column: 1},
	}
}

func (p *Printer) printNode(node interface{}) {
	file, ok := node.(*ast.File)
	if !ok {
		panicf("print: unsupported argument %v (%T)\n", file, file)
	}

	p.setComment(file.Doc)
	p.print(file.Pos(), CLASS, blank)
	p.print(classify(file.Name.Name))
	p.print(indent)

	funcs := map[string][]*ast.FuncDecl{}
	types := []*ast.TypeSpec{}

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv != nil {
				recvType := d.Recv.List[0].Type
				if ident, ok := recvType.(*ast.Ident); ok {
					funcs[ident.Name] = append(funcs[ident.Name], d)
					goto out
				}
			}
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				if s, ok := spec.(*ast.TypeSpec); ok {
					types = append(types, s)
					goto out
				}
			}
		}
		p.decl(decl)
		p.print(newline)
	out:
	}

	for _, t := range types {
		if f, ok := funcs[t.Name.Name]; ok {
			p.specClass(t, f)
		} else {
			p.spec(t)
		}
		p.print(newline)
	}

	p.print(dedent, END, newline)
}

func (p *Printer) Fprint(output io.Writer, node interface{}) {
	p.printNode(node)
	fmt.Fprint(output, string(p.output))
}

func (p *Printer) print(args ...interface{}) error {
	for _, arg := range args {
		switch x := arg.(type) {
		case token.Pos:
			if x.IsValid() {
				p.pos = p.fset.Position(x)
			}
			continue
		case GrubyToken:
			p.print(x.String())
		case token.Token:
			p.print(x.String())
		case whiteSpace:
			p.printWhitespace(x)
		case *ast.Ident:
			p.print(x.Name)
		case *ast.BasicLit:
			p.print(x.Value)
		case string:
			p.output = append(p.output, x...)
		default:
			fmt.Fprintf(os.Stderr, "print: unsupported argument %v (%T)\n", arg, arg)
			panic("gruby/printer type")
		}
	}
	return nil
}

func (p *Printer) backIndent() {
	for i := len(p.output) - 1; i > -1; i-- {
		if p.output[i] == byte(p.indentCh) {
			p.output = p.output[:i]
		} else {
			break
		}
	}
}

func (p *Printer) printWhitespace(ws whiteSpace) {
	switch ws {
	case newline:
		p.pos.Line++
		p.out.Line++
		p.pos.Column = 1
		p.out.Column = 1

		// Clean previous whitespaces
		p.backIndent()

		// Add new line
		if s := len(p.output) - 1; s > 1 &&
			p.output[s] == '\n' &&
			p.output[s-1] == '\n' {
		} else {
			p.output = append(p.output, '\n')
		}

		// Prepare indentation
		var i uint
		for i = 0; i < p.indent*p.indentSize; i++ {
			p.output = append(p.output, byte(p.indentCh))
		}
	case indent, dedent:
		switch ws {
		case indent:
			p.indent++
			p.pos.Column++
			p.out.Column++
		case dedent:
			p.indent--
			p.pos.Column--
			p.out.Column--
			p.backIndent()
			if s := len(p.output) - 1; p.output[s] == '\n' {
				p.output = p.output[:s]
			}
		}
		p.print(newline)

	default:
		p.pos.Column++
		p.out.Column++
		p.output = append(p.output, byte(ws))
	}
}

func (p *Printer) setComment(cg *ast.CommentGroup) {
	// TODO
}

// func (p *Printer) print(token interface{}) {
// 	fmt.Fprint(p.output, token)
// }
