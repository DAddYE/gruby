package gruby_test

type T int

const (
	Q = iota
	V
	c = "hello"
)

func (a T) Hello1(a, b int) {}

func (a T) Hello2() {}

func Hello3(a, b int) {}
