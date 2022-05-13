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
	ASSIGN    = "ASSIGN"
	INT       = "INT"
	SEMICOLON = "SEMICOLON"
	FUNCTION  = "FUNCTION"
	COMMA     = "COMMA"
	PLUS      = "PLUS"
	BANG      = "BANG"
	MINUS     = "MINUS"
	ASTERISK  = "ASTERISK"
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
	SLASH     = "SLASH"
)
