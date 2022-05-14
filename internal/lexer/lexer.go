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
	tok.Literal = string(r)
	switch r {
	case 0:
		tok.TokenType = token.EOF
	case '(':
		tok.TokenType = token.LPAREN
	case ')':
		tok.TokenType = token.RPAREN
	case '{':
		tok.TokenType = token.LBRACE
	case '}':
		tok.TokenType = token.RBRACE
	case ';':
		tok.TokenType = token.SEMICOLON
	case ':':
		tok.TokenType = token.COLON
	case ',':
		tok.TokenType = token.COMMA
	case '+':
		tok.TokenType = token.PLUS
	case '-':
		tok.TokenType = token.MINUS
	case '/':
		tok.TokenType = token.SLASH
	case '*':
		tok.TokenType = token.ASTERISK
	case '<':
		tok.TokenType = token.LT
	case '>':
		tok.TokenType = token.GT
	case '=':
		peek, _ := lexer.peekRune()
		if peek == '=' {
			next := lexer.readRune()
			tok.Literal += string(next)
			tok.TokenType = token.EQ
		} else {
			tok.TokenType = token.DOUBLE_EQ
		}
	case '!':
		peek, _ := lexer.peekRune()
		if peek == '=' {
			next := lexer.readRune()
			tok.Literal += string(next)
			tok.TokenType = token.NOT_EQ
		} else {
			tok.TokenType = token.BANG
		}
	default:
		if unicode.IsDigit(r) {
			digits := lexer.readNumber()
			tok.TokenType = token.NUMBER
			tok.Literal += digits
		} else if unicode.IsPunct(r) {
			tok.TokenType = token.INVALID
		} else {
			identifier := string(r) + lexer.readIdentifier()
			if keywordType, exists := token.Keywords[identifier]; exists {
				tok.Literal = identifier
				tok.TokenType = keywordType
			} else {
				tok.Literal = identifier
				tok.TokenType = token.IDENT
			}
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
	lexer.readWhile(unicode.IsSpace)
}

func (lexer *Lexer) readIdentifier() string {
	return lexer.readWhile(func(r rune) bool { return r == '_' || (!unicode.IsSpace(r) && !unicode.IsPunct(r)) })
}

func (lexer *Lexer) readNumber() string {
	return lexer.readWhile(func(r rune) bool { return unicode.IsDigit(r) || r == '.' })
}

func (lexer *Lexer) readWhile(predicate func(rune) bool) string {
	result := ""
	for r, _ := lexer.peekRune(); predicate(r); r, _ = lexer.peekRune() {
		result += string(lexer.readRune())
	}
	return result
}
