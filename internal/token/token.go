package token

type Token struct {
	TokenType string
	Literal   string
	Line      int
	Column    int
}

const (
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	LET       = "LET"
	IDENT     = "IDENT"
	DOUBLE_EQ = "DOUBLE_EQ"
	NUMBER    = "NUMBER"
	SEMICOLON = "SEMICOLON"
	COLON     = "COLON"
	FUNCTION  = "FUNCTION"
	COMMA     = "COMMA"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	SLASH     = "SLASH"
	ASTERISK  = "ASTERISK"
	BANG      = "BANG"
	LT        = "LT"
	GT        = "GT"
	RETURN    = "RETURN"
	EOF       = "EOF"
	IF        = "IF"
	ELSE      = "ELSE"
	EQ        = "EQ"
	NOT_EQ    = "NOT_EQ"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	INVALID   = "INVALID"
)

var Keywords = map[string]string{
	"let":    LET,
	"if":     IF,
	"true":   TRUE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"return": RETURN,
	"else":   ELSE,
}
