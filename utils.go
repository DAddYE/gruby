package gruby

import (
	"fmt"
	"go/ast"
	"strings"
	"unicode"
)

func panicf(format string, args ...interface{}) {
	s := fmt.Errorf(format, args...)
	panic(s)
}

func inspect(args ...interface{}) {
	const inspectFmt = "%#+v\n"
	for _, arg := range args {
		fmt.Printf(inspectFmt, arg)
	}
}

func inspect1(args ...interface{}) {
	const inspectFmt = "%+v\n"
	for _, arg := range args {
		fmt.Printf(inspectFmt, arg)
	}
}

// Functions to modify "names"
func constantize(s string) string {
	return strings.ToUpper(underscore(s))
}

func classify(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, "")
}

func orignalName(s string) string {
	return s
}

func symbolize(s string) string {
	return ":" + underscore(s)
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

func stripParens(x ast.Expr) ast.Expr {
	if x, ok := x.(*ast.ParenExpr); ok {
		return stripParens(x.X)
	}
	return x
}

func cutoff(e *ast.BinaryExpr, depth int) int {
	has4, has5, maxProblem := walkBinary(e)
	if maxProblem > 0 {
		return maxProblem + 1
	}
	if has4 && has5 {
		if depth == 1 {
			return 5
		}
		return 4
	}
	if depth == 1 {
		return 6
	}
	return 4
}

func diffPrec(expr ast.Expr, prec int) int {
	x, ok := expr.(*ast.BinaryExpr)
	if !ok || prec != x.Op.Precedence() {
		return 1
	}
	return 0
}

func reduceDepth(depth int) int {
	depth--
	if depth < 1 {
		depth = 1
	}
	return depth
}

func isPrivate(ident *ast.Ident) bool {
	return unicode.IsLower([]rune(ident.Name)[0])
}
