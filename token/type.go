package token

import "fmt"

type Token struct {
	Typ     TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(typ TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		Typ:     typ,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.Typ, t.Lexeme, t.Literal)
}
