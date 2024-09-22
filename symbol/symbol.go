package symbol

type SymbolType string

type Symbol struct {
	Type    SymbolType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Non-alphanumeric characters that are not operators or grouping
	EOF     = "EOF"     // End of file

	LITERAL = "LITERAL" // Alphanumeric characters

	// Operators
	KSTAR = "*" // Kleene star: zero or more

)
