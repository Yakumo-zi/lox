package errors

import (
	"fmt"
	"lox/token"
)

var Errors []error = make([]error, 0, 10)

func Report(line int, where, msg string) {
	Errors = append(Errors, fmt.Errorf("[line %d] Error %s:%s", line, where, msg))
}
func Error(tok *token.Token, msg string) {
	if tok.Typ == token.EOF {
		Report(tok.Line, " at end", msg)
	} else {
		Report(tok.Line, " at '"+tok.Lexeme+"'", msg)
	}
}
