package parser

import (
	"lox/ast"
	"lox/errors"
	"lox/token"
)

type Parser struct {
	tokens  []*token.Token
	current int
}

func NewParser(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}
func (p *Parser) Parse() ast.Expr {
	return p.expression()
}
func (p *Parser) expression() ast.Expr {
	return p.equality()
}
func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		op := p.previous()
		right := p.comparison()
		expr = &ast.BinaryNode{
			Left:  expr,
			Op:    *op,
			Right: right,
		}
	}
	return expr
}
func (p *Parser) comparison() ast.Expr {
	expr := p.term()
	for p.match(token.LESS, token.LESS_EQUAL, token.GREATER, token.GREATER_EQUAL) {
		op := p.previous()
		right := p.term()
		expr = &ast.BinaryNode{
			Left:  expr,
			Op:    *op,
			Right: right,
		}
	}
	return expr
}
func (p *Parser) term() ast.Expr {
	expr := p.factor()
	for p.match(token.MINUS, token.PLUS) {
		op := p.previous()
		right := p.factor()
		expr = &ast.BinaryNode{
			Left:  expr,
			Right: right,
			Op:    *op,
		}
	}
	return expr
}
func (p *Parser) factor() ast.Expr {
	expr := p.unary()
	for p.match(token.SLASH, token.STAR) {
		op := p.previous()
		right := p.unary()
		expr = &ast.BinaryNode{
			Left:  expr,
			Op:    *op,
			Right: right,
		}
	}
	return expr
}
func (p *Parser) unary() ast.Expr {
	if p.match(token.MINUS, token.BANG) {
		op := p.previous()
		expr := p.unary()
		return &ast.UnaryNode{
			Op:    *op,
			Right: expr,
		}
	}
	return p.primary()
}
func (p *Parser) primary() ast.Expr {
	if p.match(token.FALSE) {
		return &ast.LiteralNode{
			Value: false,
		}
	}
	if p.match(token.TRUE) {
		return &ast.LiteralNode{
			Value: true,
		}
	}
	if p.match(token.NIL) {
		return &ast.LiteralNode{
			Value: nil,
		}
	}
	if p.match(token.NUMBER, token.STRING) {
		return &ast.LiteralNode{
			Value: p.previous().Literal,
		}
	}
	if p.match(token.LEFT_PAREN) {
		expr := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return &ast.GroupNode{
			Expression: expr,
		}

	}
	errors.Error(p.peek(), "Expect expression.")
	return nil
}
func (p *Parser) sync() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Typ == token.SEMICOLON {
			return
		}
		switch p.peek().Typ {
		case token.CLASS:
			return
		case token.FUN:
			return
		case token.VAR:
			return
		case token.FOR:
			return
		case token.WHILE:
			return
		case token.IF:
			return
		case token.PRINT:
			return
		case token.RETURN:
			return
		}
		p.advance()
	}
}
func (p *Parser) consume(typ token.TokenType, msg string) *token.Token {
	if p.check(typ) {
		return p.advance()
	}
	errors.Error(p.peek(), msg)
	return nil
}
func (p *Parser) match(types ...token.TokenType) bool {
	for _, typ := range types {
		if p.check(typ) {
			p.advance()
			return true
		}
	}
	return false
}
func (p *Parser) check(typ token.TokenType) bool {
	if p.isAtEnd() {
		p.advance()
		return false
	}
	return p.peek().Typ == typ
}
func (p *Parser) advance() *token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Typ == token.EOF
}
func (p *Parser) peek() *token.Token {
	return p.tokens[p.current]
}
func (p *Parser) previous() *token.Token {
	return p.tokens[p.current-1]
}
