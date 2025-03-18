package interpreter

import (
	"fmt"
	"lox/ast"
	"lox/errors"
	"lox/token"
	"lox/util"
	"strings"
)

func Eval(expr ast.Expr) any {
	switch expr := expr.(type) {
	case *ast.LiteralNode:
		return evalLiteral(expr)
	case *ast.BinaryNode:
		return evalBinary(expr)
	case *ast.UnaryNode:
		return evalUnary(expr)
	case *ast.GroupNode:
		return Eval(expr.Expression)
	case *ast.ConditionNode:
		return evalCondition(expr)
	}
	return nil
}

func evalLiteral(expr *ast.LiteralNode) any {
	return expr.Value
}
func evalUnary(expr *ast.UnaryNode) any {
	right := Eval(expr.Right)
	switch expr.Op.Typ {
	case token.BANG:
		return !isTruthy(right)
	case token.MINUS:
		checkOp(expr.Op, right)
		return -right.(float64)
	}
	return nil
}
func evalBinary(expr *ast.BinaryNode) any {
	left := Eval(expr.Left)
	right := Eval(expr.Right)
	switch expr.Op.Typ {
	case token.MINUS:
		checkOps(expr.Op, left, right)
		return left.(float64) - right.(float64)
	case token.PLUS:
		switch left.(type) {
		case string:
			if _, ok := right.(string); !ok {
				errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be two numbers or two strings.", expr.Op))
				return nil
			}
			return left.(string) + right.(string)
		case float64:
			if _, ok := right.(float64); !ok {
				errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be two numbers or two strings.", expr.Op))
				return nil
			}
			return left.(float64) + right.(float64)
		}
		errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be two numbers or two strings.", expr.Op))
		return nil
	case token.SLASH:
		checkOps(expr.Op, left, right)
		return left.(float64) / right.(float64)
	case token.STAR:
		checkOps(expr.Op, left, right)
		return left.(float64) * right.(float64)
	case token.GREATER:
		checkOps(expr.Op, left, right)
		return left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		checkOps(expr.Op, left, right)
		return left.(float64) >= right.(float64)
	case token.LESS:
		checkOps(expr.Op, left, right)
		return left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		checkOps(expr.Op, left, right)
		return left.(float64) <= right.(float64)
	case token.EQUAL_EQUAL:
		return isEqual(left, right)
	case token.BANG_EQUAL:
		return !isEqual(left, right)
	default:
		return nil
	}
}
func evalCondition(expr *ast.ConditionNode) any {
	cond := Eval(expr.Condition).(bool)
	t := Eval(expr.Truth)
	f := Eval(expr.False)
	return util.When(cond, t, f)
}
func isEqual(left, right any) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	switch left := left.(type) {
	case string:
		return strings.EqualFold(left, right.(string))
	case float64:
		return left == right.(float64)
	case bool:
		return left == right.(bool)
	default:
		return false
	}
}
func checkOp(tok token.Token, operand any) {
	if _, ok := operand.(float64); ok {
		return
	}
	errors.Error(&tok, fmt.Sprintf("%+v Operand must be a number", tok))
}
func checkOps(tok token.Token, left, right any) {
	if _, ok := left.(float64); !ok {
		errors.Error(&tok, fmt.Sprintf("%+v %+v %+v , Operands must be a number", left, tok, right))
		return
	}
	if _, ok := right.(float64); !ok {
		errors.Error(&tok, fmt.Sprintf("%+v %+v %+v , Operands must be a number", left, tok, right))
		return
	}

}
func isTruthy(val any) bool {
	if val == nil {
		return false
	}
	switch v := val.(type) {
	case bool:
		return v
	default:
		return true
	}
}
