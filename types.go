package gruby

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
