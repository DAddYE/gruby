package gruby

import (
	"fmt"
	"go/ast"
	"os"
	"strings"
)

const DEBUG = true

func debug(args ...string) {
	if !DEBUG {
		return
	}
	i := 0
	for ; i < len(args); i++ {
		fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m", args[i])
	}
	i--
	if args[i][len(args[i])-1] != '\n' {
		fmt.Fprint(os.Stderr, "\n")
	}
}

func debugf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	debug(str)
}

func panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	panic(s)
}

var constantize = strings.ToUpper

func classify(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, "")
}

// Remove parens
func stripParens(x ast.Expr) ast.Expr {
	if x, ok := x.(*ast.ParenExpr); ok {
		return stripParens(x.X)
	}
	return x
}
