package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sevenreup/chewa/src/lexer"
)

func main() {
	file, err := os.ReadFile("./examples/main.ny")
	if err != nil {
		log.Fatal(err)
	}
	lex := lexer.NewLexer(file)

	tokens := lex.AccumTokens()

	for _, v := range tokens {
		fmt.Println(v)
	}
}
