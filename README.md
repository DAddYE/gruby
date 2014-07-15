## Gruby

Go-Lang to Ruby transpiler.

This is an attempt to translate Go source code to (readable) Ruby.

The aim is to bring the simplicity of the `go` syntax and the compiler capabilities (like type
checks) to legacy environments that runs on ruby.

This is still a research project, isn't completed yet, but I'll glad to know your feedback.

To give you an idea right now it translates:

```go
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
```

Into:

```rb
class GrubyTest
  Q = 1 << 0
  V = Q + 1
  C = "hello"
  Q = -1

  attr_accessor :a, :b
  attr_accessor :c

  def hi_there(a, b)
  end
  private :hi_there

  class A < Fixnum; end
  class B < Array; end
  class C < Hash; end
  class D < Proc; end
  class E; end
  class F; end
  class S < Struct.new(:a, :b, :t, :c, :z)
    def hello_world1(a, b)
    end

    def hello_world2()
    end
    private :hello_world2

  end

end
```

## License

MIT
