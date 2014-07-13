package gruby

import (
	"fmt"
	"go/ast"
	"strings"
	"unicode"
)

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

func underscore(s string) string {
	b := []rune(s)

	for i := 0; i < len(b); i++ {
		ch := b[i]

		if i == 0 {
			if unicode.IsUpper(ch) {
				b[0] = unicode.ToLower(ch)
			}
			continue
		}

		if unicode.IsUpper(ch) {
			b[i] = unicode.ToLower(ch)
			b = append(b, 0)
			copy(b[i+1:], b[i:])
			b[i] = '_'
		}
	}

	return string(b)
}

// Remove parens
func stripParens(x ast.Expr) ast.Expr {
	if x, ok := x.(*ast.ParenExpr); ok {
		return stripParens(x.X)
	}
	return x
}
