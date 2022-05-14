package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/AndreasChristianson/monkey/internal/lexer"
	"github.com/AndreasChristianson/monkey/internal/token"
)

const PROMPT = "🐒 »"

func Repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		lexer := lexer.New(scanner.Bytes())

		for {
			tok := lexer.NextToken()
			fmt.Fprintf(out, "%+v\n", tok)
			if tok.TokenType == token.EOF {
				break
			}
		}
	}
}
