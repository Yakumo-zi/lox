package interpreter

import (
	"fmt"
	"lox/ast"
	"lox/errors"
	"lox/token"
	"lox/util"
	"strings"
)

type Interpreter struct {
	env *Environment
}

func NewInterpreter(env *Environment) *Interpreter {
	return &Interpreter{
		env: env,
	}
}
func (i *Interpreter) Run(stmts []ast.Stmt) (ret any, err error) {
	for _, stmt := range stmts {
		ret, err = i.evalStatement(stmt)
		if err != nil {
			return
		}
	}
	return
}

func (i *Interpreter) evalStatement(stmt ast.Stmt) (any, error) {
	switch stmt := stmt.(type) {
	case *ast.ExpressionStmt:
		return i.eval(stmt.Expression)
	case *ast.PrintStmt:
		val, err := i.eval(stmt.Value)
		if err != nil {
			return nil, err
		}
		if val == nil {
			fmt.Printf("Error, <nil> pointer deref!\n")
		}
		fmt.Printf("%#v\n", val)
		return nil, nil
	case *ast.VariableStmt:
		var val any
		if stmt.Value != nil {
			var err error
			val, err = i.eval(stmt.Value)
			if err != nil {
				return nil, err
			}
		}
		i.env.define(stmt.Name.Lexeme, val)
		return nil, nil
	case *ast.BlockStmt:
		return i.evalBlock(stmt.Stmts, NewEnvironment(i.env))
	case *ast.IfStmt:
		cond, err := i.eval(stmt.Cond)
		if err != nil {
			return nil, err
		}
		if isTruthy(cond) {
			return i.evalStatement(stmt.Then)
		} else if stmt.Else != nil {
			return i.evalStatement(stmt.Else)
		}
		return nil, nil
	default:
		return nil, nil
	}
}
func (i *Interpreter) evalBlock(stmts []ast.Stmt, env *Environment) (ret any, err error) {
	previous := i.env
	defer func() {
		i.env = previous
	}()
	i.env = env
	for _, stmts := range stmts {
		ret, err = i.evalStatement(stmts)
		if err != nil {
			return nil, err
		}
	}
	return
}
func (i *Interpreter) eval(expr ast.Expr) (any, error) {
	switch expr := expr.(type) {
	case *ast.LiteralNode:
		return i.evalLiteral(expr), nil
	case *ast.BinaryNode:
		return i.evalBinary(expr)
	case *ast.UnaryNode:
		return i.evalUnary(expr)
	case *ast.GroupNode:
		return i.eval(expr.Expression)
	case *ast.ConditionNode:
		return i.evalCondition(expr)
	case *ast.VariableNode:
		return i.env.get(expr.Name)
	case *ast.AssignNode:
		val, err := i.eval(expr.Value)
		if err != nil {
			return nil, err
		}
		v, _ := i.env.assign(expr.Name, val)
		return v, nil
	}
	return nil, nil
}

func (i *Interpreter) evalLiteral(expr *ast.LiteralNode) any {
	return expr.Value
}
func (i *Interpreter) evalUnary(expr *ast.UnaryNode) (any, error) {
	right, err := i.eval(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Op.Typ {
	case token.BANG:
		return !isTruthy(right), nil
	case token.MINUS:
		if ok := checkOp(expr.Op, right); !ok {
			return nil, fmt.Errorf("%+v Operand must be a number", expr.Op.Lexeme)
		}
		return -right.(float64), nil
	}
	return nil, nil
}
func (i *Interpreter) evalBinary(expr *ast.BinaryNode) (any, error) {
	left, err := i.eval(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := i.eval(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Op.Typ {
	case token.MINUS:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil, fmt.Errorf("%+v Operands must be numbers or strings", expr.Op)
		}
		return left.(float64) - right.(float64), nil
	case token.PLUS:
		switch left.(type) {
		case string:
			if right, ok := right.(float64); ok {
				return fmt.Sprintf("%s%+v", left, right), nil
			}
			if _, ok := right.(string); ok {
				return left.(string) + right.(string), nil
			}
			errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be  numbers or strings.", expr.Op))
			return nil, fmt.Errorf("%+v Operands must be  numbers or strings.", expr.Op)

		case float64:
			if right, ok := right.(string); ok {
				return fmt.Sprintf("%+v%s", left, right), nil
			}
			if _, ok := right.(float64); ok {

				return left.(float64) + right.(float64), nil
			}
			errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be  numbers or  strings.", expr.Op))
		}
		errors.Error(&expr.Op, fmt.Sprintf("%+v Operands must be  numbers or  strings.", expr.Op))
		return nil, fmt.Errorf("%+v Operands must be  numbers or  strings.", expr.Op)
	case token.SLASH:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil, fmt.Errorf("%+v Operands must be numbers", expr.Op)
		}
		return left.(float64) / right.(float64), nil
	case token.STAR:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil, fmt.Errorf("%+v Operands must be numbers", expr.Op)
		}
		return left.(float64) * right.(float64), nil
	case token.GREATER:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil, fmt.Errorf("%+v Operands must be numbers or strings", expr.Op)
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) > 0, nil
		}
		return left.(float64) > right.(float64), nil
	case token.GREATER_EQUAL:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil, fmt.Errorf("%+v Operands must be numbers or strings", expr.Op)
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) >= 0, nil
		}
		return left.(float64) >= right.(float64), nil
	case token.LESS:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil, fmt.Errorf("%+v Operands must be numbers or strings", expr.Op)
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) < 0, nil
		}
		return left.(float64) < right.(float64), nil
	case token.LESS_EQUAL:
		if ok := checkOps(expr.Op, left, right); !ok {
			return nil, fmt.Errorf("%+v Operands must be numbers or strings", expr.Op)
		}
		if left, ok := left.(string); ok {
			return strings.Compare(left, right.(string)) <= 0, nil
		}
		return left.(float64) <= right.(float64), nil
	case token.EQUAL_EQUAL:
		return isEqual(left, right), nil
	case token.BANG_EQUAL:
		return !isEqual(left, right), nil
	default:
		return nil, fmt.Errorf("not supported operator %#v", expr.Op.Lexeme)
	}
}
func (i *Interpreter) evalCondition(expr *ast.ConditionNode) (any, error) {
	cond, err := i.eval(expr.Condition)
	if err != nil {
		return nil, err
	}
	t, err := i.eval(expr.Truth)
	if err != nil {
		return nil, err
	}
	f, err := i.eval(expr.False)
	if err != nil {
		return nil, err
	}
	return util.When(cond.(bool), t, f), nil
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
		if tok.Typ == token.STAR || tok.Typ == token.SLASH || tok.Typ == token.MINUS {
			errors.Error(&tok, fmt.Sprintf("%+v %+v %+v , Operands must be two numbers ", left, tok.Lexeme, right))
			return false
		}
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
