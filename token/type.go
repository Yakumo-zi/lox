package token

import "fmt"

type Token struct {
	typ     TokenType
	lexeme  string
	literal any
	line    int
}

func NewToken(typ TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		typ:     typ,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}
func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.typ, t.lexeme, t.literal)
}
