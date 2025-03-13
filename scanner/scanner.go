package scanner

import (
	"lox/errors"
	"lox/token"
	"lox/util"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []*token.Token
	start   int
	line    int
	current int
}

func NewSacnner(source string) *Scanner {
	return &Scanner{
		source: source,
		tokens: make([]*token.Token, 0, 10),
	}
}

func (s *Scanner) ScanTokens() []*token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()

	}
	s.tokens = append(s.tokens, token.NewToken(token.EOF, "", nil, s.line))
	return s.tokens
}
func (s *Scanner) scanToken() {
	ch := s.advance()
	switch ch {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		s.addToken(util.When(s.match('='), token.BANG_EQUAL, token.BANG))
	case '=':
		s.addToken(util.When(s.match('='), token.EQUAL_EQUAL, token.EQUAL))
	case '<':
		s.addToken(util.When(s.match('='), token.LESS_EQUAL, token.LESS))
	case '>':
		s.addToken(util.When(s.match('='), token.GREATER_EQUAL, token.GREATER))
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':

	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigital(ch) {
			s.number()
		} else if s.isAlpha(ch) {
			s.identifier()
		} else {
			errors.Report(s.line, string(ch), "Unexpected character.")
		}
	}
}
func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	typ, ok := token.KeyWords[text]
	if !ok {
		typ = token.IDENTIFIER
	}
	s.addToken(typ)
}
func (s *Scanner) number() {
	for s.isDigital(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && s.isDigital(s.peekNext()) {
		s.advance()
		for s.isDigital(s.peek()) {
			s.advance()
		}
	}
	f, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addTokenLiteral(token.NUMBER, f)
}
func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		errors.Report(s.line, "", "Unterminated string.")
		return
	}
	s.advance()

	value := string(s.source[s.start+1 : s.current-1])
	s.addTokenLiteral(token.STRING, value)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}
func (s *Scanner) isAlphaNumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigital(c)
}
func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) isDigital(c byte) bool {
	return c >= '0' && c <= '9'
}
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}
func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}
func (s *Scanner) advance() byte {
	b := s.source[s.current]
	s.current++
	return b
}

func (s *Scanner) addToken(typ token.TokenType) {
	s.addTokenLiteral(typ, nil)
}
func (s *Scanner) addTokenLiteral(typ token.TokenType, literal any) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, token.NewToken(typ, text, literal, s.line))
}
