package scanner

import (
	"lox/token"
	"testing"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		source   string
		expected []token.TokenType
	}{
		{
			source:   `print "Hello, world!";`,
			expected: []token.TokenType{token.PRINT, token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `true;`,
			expected: []token.TokenType{token.TRUE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `false;`,
			expected: []token.TokenType{token.FALSE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `1234;`,
			expected: []token.TokenType{token.NUMBER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `12.34;`,
			expected: []token.TokenType{token.NUMBER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `"I am a string";`,
			expected: []token.TokenType{token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `"";`,
			expected: []token.TokenType{token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `"123";`,
			expected: []token.TokenType{token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `add + me;`,
			expected: []token.TokenType{token.IDENTIFIER, token.PLUS, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `subtract - me;`,
			expected: []token.TokenType{token.IDENTIFIER, token.MINUS, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `multiply * me;`,
			expected: []token.TokenType{token.IDENTIFIER, token.STAR, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `divide / me;`,
			expected: []token.TokenType{token.IDENTIFIER, token.SLASH, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `-negateMe;`,
			expected: []token.TokenType{token.MINUS, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `less < than;`,
			expected: []token.TokenType{token.IDENTIFIER, token.LESS, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `lessThan <= orEqual;`,
			expected: []token.TokenType{token.IDENTIFIER, token.LESS_EQUAL, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `greater > than;`,
			expected: []token.TokenType{token.IDENTIFIER, token.GREATER, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `greaterThan >= orEqual;`,
			expected: []token.TokenType{token.IDENTIFIER, token.GREATER_EQUAL, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `1 == 2;`,
			expected: []token.TokenType{token.NUMBER, token.EQUAL_EQUAL, token.NUMBER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `"cat" != "dog";`,
			expected: []token.TokenType{token.STRING, token.BANG_EQUAL, token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `314 == "pi";`,
			expected: []token.TokenType{token.NUMBER, token.EQUAL_EQUAL, token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `123 == "123";`,
			expected: []token.TokenType{token.NUMBER, token.EQUAL_EQUAL, token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `!true;`,
			expected: []token.TokenType{token.BANG, token.TRUE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `!false;`,
			expected: []token.TokenType{token.BANG, token.FALSE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `true and false;`,
			expected: []token.TokenType{token.TRUE, token.AND, token.FALSE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `true and true;`,
			expected: []token.TokenType{token.TRUE, token.AND, token.TRUE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `false or false;`,
			expected: []token.TokenType{token.FALSE, token.OR, token.FALSE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `true or false;`,
			expected: []token.TokenType{token.TRUE, token.OR, token.FALSE, token.SEMICOLON, token.EOF},
		},
		{
			source:   `var average = (min + max) / 2;`,
			expected: []token.TokenType{token.VAR, token.IDENTIFIER, token.EQUAL, token.LEFT_PAREN, token.IDENTIFIER, token.PLUS, token.IDENTIFIER, token.RIGHT_PAREN, token.SLASH, token.NUMBER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `print "Hello, world!";`,
			expected: []token.TokenType{token.PRINT, token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `{ print "One statement."; print "Two statements."; }`,
			expected: []token.TokenType{token.LEFT_BRACE, token.PRINT, token.STRING, token.SEMICOLON, token.PRINT, token.STRING, token.SEMICOLON, token.RIGHT_BRACE, token.EOF},
		},
		{
			source:   `var breakfast = "bagels";`,
			expected: []token.TokenType{token.VAR, token.IDENTIFIER, token.EQUAL, token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `print breakfast;`,
			expected: []token.TokenType{token.PRINT, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `breakfast = "beignets";`,
			expected: []token.TokenType{token.IDENTIFIER, token.EQUAL, token.STRING, token.SEMICOLON, token.EOF},
		},
		{
			source:   `print breakfast;`,
			expected: []token.TokenType{token.PRINT, token.IDENTIFIER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `if (condition) { print "yes"; } else { print "no"; }`,
			expected: []token.TokenType{token.IF, token.LEFT_PAREN, token.IDENTIFIER, token.RIGHT_PAREN, token.LEFT_BRACE, token.PRINT, token.STRING, token.SEMICOLON, token.RIGHT_BRACE, token.ELSE, token.LEFT_BRACE, token.PRINT, token.STRING, token.SEMICOLON, token.RIGHT_BRACE, token.EOF},
		},
		{
			source:   `for (var a = 1; a < 10; a = a + 1) { print a; }`,
			expected: []token.TokenType{token.FOR, token.LEFT_PAREN, token.VAR, token.IDENTIFIER, token.EQUAL, token.NUMBER, token.SEMICOLON, token.IDENTIFIER, token.LESS, token.NUMBER, token.SEMICOLON, token.IDENTIFIER, token.EQUAL, token.IDENTIFIER, token.PLUS, token.NUMBER, token.RIGHT_PAREN, token.LEFT_BRACE, token.PRINT, token.IDENTIFIER, token.SEMICOLON, token.RIGHT_BRACE, token.EOF},
		},
		{
			source:   `var a = 1;`,
			expected: []token.TokenType{token.VAR, token.IDENTIFIER, token.EQUAL, token.NUMBER, token.SEMICOLON, token.EOF},
		},
		{
			source:   `while (a < 10) { print a; a = a + 1; }`,
			expected: []token.TokenType{token.WHILE, token.LEFT_PAREN, token.IDENTIFIER, token.LESS, token.NUMBER, token.RIGHT_PAREN, token.LEFT_BRACE, token.PRINT, token.IDENTIFIER, token.SEMICOLON, token.IDENTIFIER, token.EQUAL, token.IDENTIFIER, token.PLUS, token.NUMBER, token.SEMICOLON, token.RIGHT_BRACE, token.EOF},
		},
		{
			source:   `fun printSum(a, b) { print a + b; }`,
			expected: []token.TokenType{token.FUN, token.IDENTIFIER, token.LEFT_PAREN, token.IDENTIFIER, token.COMMA, token.IDENTIFIER, token.RIGHT_PAREN, token.LEFT_BRACE, token.PRINT, token.IDENTIFIER, token.PLUS, token.IDENTIFIER, token.SEMICOLON, token.RIGHT_BRACE, token.EOF},
		},
		{
			source:   `makeBreakfast(bacon, eggs, toast);`,
			expected: []token.TokenType{token.IDENTIFIER, token.LEFT_PAREN, token.IDENTIFIER, token.COMMA, token.IDENTIFIER, token.COMMA, token.IDENTIFIER, token.RIGHT_PAREN, token.SEMICOLON, token.EOF},
		},
	}

	for _, test := range tests {
		scanner := NewSacnner(test.source)
		tokens := scanner.ScanTokens()

		if len(tokens) != len(test.expected) {
			t.Fatalf("expected %d tokens, got %d", len(test.expected), len(tokens))
		}

		for i, token := range tokens {
			if token.Typ != test.expected[i] {
				t.Errorf("expected token type %s, got %s", test.expected[i], token.Typ)
			}
		}
	}
}
