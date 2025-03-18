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
		if ok := checkOp(expr.Op, right); !ok {
			return nil
		}
		return -right.(float64)
	}
	return nil
}
func evalBinary(expr *ast.BinaryNode) any {
	left := Eval(expr.Left)
	right := Eval(expr.Right)
	switch expr.Op.Typ {
	case token.MINUS:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil
		}
		return left.(float64) - right.(float64)
	case token.PLUS:
		switch left.(type) {
		case string:
			if right, ok := right.(float64); ok {
				return fmt.Sprintf("%s%+v", left, right)
			}
			if _, ok := right.(string); ok {
				return left.(string) + right.(string)
			}
			errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be  numbers or strings.", expr.Op))

		case float64:
			if right, ok := right.(string); ok {
				return fmt.Sprintf("%+v%s", left, right)
			}
			if _, ok := right.(float64); ok {

				return left.(float64) + right.(float64)
			}
			errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be  numbers or  strings.", expr.Op))
		}
		errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be  numbers or  strings.", expr.Op))
		return nil
	case token.SLASH:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil
		}
		return left.(float64) / right.(float64)
	case token.STAR:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil
		}
		return left.(float64) * right.(float64)
	case token.GREATER:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) > 0
		}
		return left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) >= 0
		}
		return left.(float64) >= right.(float64)
	case token.LESS:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) < 0
		}
		return left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) <= 0
		}
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
func checkOp(tok token.Token, operand any) bool {
	if _, ok := operand.(float64); ok {
		return true
	}
	errors.Error(&tok, fmt.Sprintf("%+v Operand must be a number", tok.Lexeme))
	return false
}
func checkOps(tok token.Token, left, right any) bool {

	if _, ok := left.(float64); ok {
		if right, ok := right.(float64); !ok {
			errors.Error(&tok, fmt.Sprintf("%+v %+v %+v , Operands must be two numbers or two strings", left, tok.Lexeme, right))
			return false
		} else if right == 0 {
			errors.Error(&tok, fmt.Sprintf("%+v %+v %+v , Devide zero!", left, tok.Lexeme, right))
			return false
		}
	} else if _, ok = left.(string); ok {
		if right, ok := right.(string); !ok {
			errors.Error(&tok, fmt.Sprintf("%+v %+v %+v , Operands must be two numbers or two strings", left, tok.Lexeme, right))
			return false
		}
	} else {
		errors.Error(&tok, fmt.Sprintf("%+v %+v %+v , Operands must be  two numbers or  two strings", left, tok.Lexeme, right))
		return false
	}
	return true
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
