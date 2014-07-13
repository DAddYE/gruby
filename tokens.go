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
	PRIVATE
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
	PRIVATE: "private",
}

func (t GrubyToken) String() string { return tokens[t] }
