package lexer

import (
	"testing"

	"github.com/AndreasChristianson/monkey/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestNextToken1(t *testing.T) {
	text := `let five = 5;
	let 🐒 = 10;

	let add = fn(x, y) {
		x + y;
	};

	let lett_uce123 = add(five, 🐒);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10.02) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	`

	lexer := New([]byte(text))

	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.EQ, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "🐒"},
		{token.EQ, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.EQ, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "lett_uce123"},
		{token.EQ, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "🐒"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.GT, ">"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10.02"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.NUMBER, "10"},
		{token.DOUBLE_EQ, "=="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "10"},
		{token.NOT_EQ, "!="},
		{token.NUMBER, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, "\x00"},
	}
	for _, tt := range tests {
		tok := lexer.NextToken()

		assert.Equal(t, tt.expectedLiteral, tok.Literal)
		assert.Equal(t, tt.expectedType, tok.TokenType)
	}
}

func TestNextToken(t *testing.T) {
	text := `
	()=
	===  my🐒_monkey
	`

	lexer := New([]byte(text))

	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.EQ, "="},
		{token.DOUBLE_EQ, "=="},
		{token.EQ, "="},
		{token.IDENT, "my🐒_monkey"},
		{token.EOF, "\x00"},
	}
	for _, tt := range tests {
		tok := lexer.NextToken()

		assert.Equal(t, tt.expectedLiteral, tok.Literal, tok)
		assert.Equal(t, tt.expectedType, tok.TokenType, tok)
	}
}
