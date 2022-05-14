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
	if r == 0 {
		tok.TokenType = token.EOF
		return tok
	}
	peek, _ := lexer.peekRune()
	twoCharPunct := string(r) + string(peek)

	if punctuationType, exists := token.DoubleCharPunctuation[twoCharPunct]; exists {
		lexer.readRune()
		tok.TokenType = punctuationType
		tok.Literal = twoCharPunct
		return tok
	}
	if punctuationType, exists := token.SingleCharPunctuation[r]; exists {
		tok.TokenType = punctuationType
		return tok
	}
	if unicode.IsDigit(r) {
		digits := lexer.readNumber()
		tok.TokenType = token.NUMBER
		tok.Literal += digits
		return tok
	}
	if unicode.IsPunct(r) || unicode.IsControl(r) {
		tok.TokenType = token.INVALID
		return tok
	}

	phrase := string(r) + lexer.readIdentifier()
	if keywordType, exists := token.Keywords[phrase]; exists {
		tok.Literal = phrase
		tok.TokenType = keywordType
	} else {
		tok.Literal = phrase
		tok.TokenType = token.IDENT

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
	return lexer.readWhile(func(r rune) bool {
		return r == '_' || (!unicode.IsSpace(r) && !unicode.IsPunct(r) && r != 0)
	})
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
