package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]
	lex, err := NewLexer(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, t := range lex.Tokens {
		fmt.Println("[" + strconv.Itoa(t.Row+1) + "," + strconv.Itoa(t.Col+1) + "] " + t.Value)
	}
}
