package ast

import (
	"log"
	"lox/token"
	"testing"
)

func TestAstPrinter(t *testing.T) {
	var expr Expr
	expr = BinaryNode{
		Left: UnaryNode{
			Op: *token.NewToken(token.MINUS, "-", nil, 1),
			Right: LiteralNode{
				Value: 123,
			},
		},
		Op: *token.NewToken(token.STAR, "*", nil, 1),
		Right: GroupNode{
			LiteralNode{
				Value: 45.67,
			},
		},
	}
	str := AstPrinter(expr)
	log.Print(str)
}
