package ast

import (
	"lox/token"
)

type BinaryNode struct {
	Left  Expr
	Right Expr
	Op    token.Token
}
type UnaryNode struct {
	Op    token.Token
	Right Expr
}
type GroupNode struct {
	Expression Expr
}
type LiteralNode struct {
	Value any
}
