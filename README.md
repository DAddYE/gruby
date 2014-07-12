## Gruby

Go-Lang to Ruby transpiler.

This is an attempt to translate Go source code to (readable) Ruby.

The aim is to bring the simplicity of the `go` syntax and the compiler capabilities (like type
checks) to legacy environments that runs on ruby.

This is still a research project, isn't completed yet, but I'll glad to know your feedback.

To give you an idea right now it translates:

```go
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
```

Into:

```rb
class GrubyTest
  Q = 0
  V = 1
  C = "hello"

  def Hello3(a, b)
  end

  class T < Fixnum
    def Hello1(a, b)
    end

    def Hello2()
    end

  end
end
```

## License

MIT