package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/daddye/gruby"
)

type T int

func (x T) tester(a, b int) {

}

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	f, err := parser.ParseFile(fset, "example_test.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup the destination
	p := gruby.NewPrinter(fset)
	p.Fprint(os.Stdout, f)

	// Print name of package
	// p.Class(f.Name.Name)
	// p.Indent()
	// for _, node := range f.Decls {
	// 	fmt.Printf("%#v\n", node)
	// 	switch node.(type) {
	// 	case *ast.GenDecl:
	// 	}
	// }
	// p.Dedent()
	// p.End()
}
