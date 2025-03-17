package ast

import (
	"fmt"
	"lox/util"
)

func AstPrinter(expr Expr) string {
	switch expr := expr.(type) {
	case BinaryNode:
		return bianryPrinter(expr)
	case UnaryNode:
		return unaryPrinter(expr)
	case GroupNode:
		return groupingPrinter(expr)
	case LiteralNode:
		return literalPrinter(expr)
	}
	return ""
}
func bianryPrinter(expr BinaryNode) string {
	left := AstPrinter(expr.Left)
	op := expr.Op.Lexeme
	right := AstPrinter(expr.Right)
	return fmt.Sprintf("(%s %s %s)", op, left, right)
}
func unaryPrinter(expr UnaryNode) string {
	op := expr.Op.Lexeme
	right := AstPrinter(expr.Right)
	return fmt.Sprintf("(%s %s)", op, right)
}
func groupingPrinter(expr GroupNode) string {
	str := AstPrinter(expr.Expression)
	return fmt.Sprintf("(%s)", str)
}
func literalPrinter(expr LiteralNode) string {
	return util.When(expr.Value == nil, "nil", fmt.Sprint(expr.Value))
}
