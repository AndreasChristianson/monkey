package main

import (
	"fmt"
	"os"

	"github.com/AndreasChristianson/monkey/repl"
)

const version = "0.0.1"

func main() {
	fmt.Printf("Monkey [%s]\n", version)
	fmt.Printf("repl active; type Something.\n")
	repl.Repl(os.Stdin, os.Stdout)
}
