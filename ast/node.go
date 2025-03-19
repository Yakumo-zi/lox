package ast

import (
	"lox/token"
)

type PrintStmt struct {
	Value any
}
type ExpressionStmt struct {
	Expression Expr
}
type VariableStmt struct {
	Name  token.Token
	Value Expr
}
type VariableNode struct {
	Name token.Token
}

type AssignNode struct {
	Name  token.Token
	Value Expr
}
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
type ConditionNode struct {
	Condition Expr
	Truth     Expr
	False     Expr
}
