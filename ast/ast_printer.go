package ast

import (
	"fmt"
	"lox/util"
)

func AstPrinter(expr Expr) string {
	switch expr := expr.(type) {
	case *BinaryNode:
		return bianryPrinter(expr)
	case *UnaryNode:
		return unaryPrinter(expr)
	case *GroupNode:
		return groupingPrinter(expr)
	case *LiteralNode:
		return literalPrinter(expr)
	case *ConditionNode:
		return conditionPrinter(expr)
	case *VariableNode:
		return fmt.Sprintf("var %s;", expr.Name.Lexeme)
	default:
		return fmt.Sprintf("not a valid node, %+#v", expr)
	}
}
func conditionPrinter(expr *ConditionNode) string {
	condition := AstPrinter(expr.Condition)
	trueVal := AstPrinter(expr.Truth)
	falseVal := AstPrinter(expr.False)
	return fmt.Sprintf("(%s ? %s : %s)", condition, trueVal, falseVal)
}
func bianryPrinter(expr *BinaryNode) string {
	left := AstPrinter(expr.Left)
	op := expr.Op.Lexeme
	right := AstPrinter(expr.Right)
	return fmt.Sprintf("(%s %s %s)", op, left, right)
}
func unaryPrinter(expr *UnaryNode) string {
	op := expr.Op.Lexeme
	right := AstPrinter(expr.Right)
	return fmt.Sprintf("(%s %s)", op, right)
}
func groupingPrinter(expr *GroupNode) string {
	str := AstPrinter(expr.Expression)
	return fmt.Sprintf("(%s)", str)
}
func literalPrinter(expr *LiteralNode) string {
	return util.When(expr.Value == nil, "nil", fmt.Sprint(expr.Value))
}
