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
	STRUCT
	NEW
	PLUS
	ATTR_ACCESSOR
)

var tokens = [...]string{
	ILLEGAL:       "",
	CLASS:         "class",
	END:           "end",
	DEF:           "def",
	DOT:           ".",
	NIL:           "nil",
	REQUIRE:       "require",
	INHERIT:       "<",
	SEMI:          ";",
	VAR:           "@",
	PRIVATE:       "private",
	STRUCT:        "Struct",
	NEW:           "new",
	PLUS:          "+",
	ATTR_ACCESSOR: "attr_accessor",
}

func (t GrubyToken) String() string { return tokens[t] }
