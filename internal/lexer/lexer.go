package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/AndreasChristianson/monkey/internal/token"
)

const (
	nl = '\n'
)

type Lexer struct {
	line   int
	column int
	text   []byte
}

func New(text []byte) *Lexer {
	return &Lexer{
		text:   text,
		column: 0,
		line:   0,
	}
}

func (lexer *Lexer) NextToken() token.Token {
	lexer.gulp()
	tok := token.Token{
		Line:   lexer.line,
		Column: lexer.column,
	}

	r := lexer.readRune()
	switch r {
	case 0:
		tok.TokenType = token.EOF
	case '(':
		tok.Literal = string(r)
		tok.TokenType = token.LPAREN
	case ')':
		tok.Literal = string(r)
		tok.TokenType = token.RPAREN
	case '{':
		tok.Literal = string(r)
		tok.TokenType = token.LBRACE
	case '}':
		tok.Literal = string(r)
		tok.TokenType = token.RBRACE
	case '=':
		peek, _ := lexer.peekRune()
		if peek == '=' {
			next := lexer.readRune()
			tok.Literal = string(string(r) + string(next))
			tok.TokenType = token.EQ
		} else {
			tok.Literal = string(r)
			tok.TokenType = token.ASSIGN
		}
	}
	return tok
}

func (lexer *Lexer) readRune() rune {
	if len(lexer.text) == 0 {
		return 0
	} else {
		r, size := lexer.peekRune()
		lexer.text = lexer.text[size:]
		if r == nl {
			lexer.line++
			lexer.column = 0
		} else {
			lexer.column++
		}
		return r
	}
}

func (lexer *Lexer) peekRune() (rune, int) {
	if len(lexer.text) == 0 {
		return 0, 0
	} else {
		return utf8.DecodeRune(lexer.text)
	}
}
func (lexer *Lexer) gulp() {
	for r, _ := lexer.peekRune(); unicode.IsSpace(r); r, _ = lexer.peekRune() {
		lexer.readRune()
	}
}
