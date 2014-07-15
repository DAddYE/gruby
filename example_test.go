package gruby_test

type A int
type B []string
type C map[string]int
type D func(string) string
type E chan<- string
type F interface {
	isUseless()
}

const (
	Q = 1 << iota
	V
	c = "hello"
	q = -1
)

var (
	a, b int
	c    string
)

type S struct {
	a, b int
	T
	c string
	z func()
}

func (a S) HelloWorld1(a, b int) {
	x := 1
}

func (a S) helloWorld2() {}

func hiThere(a, b int) {}
