package gruby

type GrubyToken int

const (
	ILLEGAL GrubyToken = iota
	CLASS
	END
	DEF
	DOT
	NIL
	REQUIRE
	INHERIT
	SEMI
	VAR
)

var tokens = [...]string{
	ILLEGAL: "",
	CLASS:   "class",
	END:     "end",
	DEF:     "def",
	DOT:     ".",
	NIL:     "nil",
	REQUIRE: "require",
	INHERIT: "<",
	SEMI:    ";",
	VAR:     "@",
}

func (t GrubyToken) String() string { return tokens[t] }

// Language specific

var goTypeToRuby = map[string]string{
	// Numeric types
	"uint8":  "Fixnum",
	"uint16": "Fixnum",
	"uint32": "Fixnum",
	"uint64": "Bignum",

	"int8":  "Fixnum",
	"int16": "Fixnum",
	"int32": "Fixnum",
	"int64": "Fixnum",

	"float32": "Float",
	"float64": "Float",

	"complex64":  "Float",
	"complex128": "Float",

	"byte": "Fixnum",
	"rune": "Fixnum",

	"int":  "Fixnum",
	"uint": "Fixnum",
}
