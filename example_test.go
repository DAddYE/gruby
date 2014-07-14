package gruby_test

type T int

const (
	Q = 1 << iota
	V
	c = "hello"
)

type S struct {
	a, b int
	T
	c string
}

func (a T) HelloWorld1(a, b int) {}

func (a T) helloWorld2() {}

func HelloWorld3(a, b int) {}
