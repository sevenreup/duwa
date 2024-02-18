package main

import (
	"fmt"

	"github.com/sevenreup/chewa/src/scanner"
)

func main() {
	lex := scanner.NewScanner("./examples/main.ny")
	defer lex.Close()

	tokens := lex.AccumTokens()

	for _, v := range tokens {
		fmt.Println(v)
	}
}
