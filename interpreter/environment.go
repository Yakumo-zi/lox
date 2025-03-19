package interpreter

import (
	"fmt"
	"lox/errors"
	"lox/token"
)

type Environment struct {
	enclosing *Environment
	values    map[string]any
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		values:    make(map[string]any),
		enclosing: enclosing,
	}
}
func (e *Environment) define(name string, value any) {
	e.values[name] = value
}
func (e *Environment) get(name token.Token) (any, error) {
	if value, ok := e.values[name.Lexeme]; ok {
		return value, nil
	}
	if e.enclosing != nil {
		return e.enclosing.get(name)
	}
	errors.Error(&name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
	return nil, fmt.Errorf("Undefined variable '%s'.", name.Lexeme)
}
func (e *Environment) assign(name token.Token, value any) (any, error) {
	if _, ok := e.values[name.Lexeme]; !ok {
		errors.Error(&name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
		return nil, fmt.Errorf("Undefined variable '%s'.", name.Lexeme)
	}
	if e.enclosing != nil {
		return e.enclosing.assign(name, value)
	}
	e.values[name.Lexeme] = value
	return value, nil
}
