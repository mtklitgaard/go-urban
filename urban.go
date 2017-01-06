package main

import (
	"fmt"
	"os"
)

func main() {
	wordToLookup := os.Args[1]
	definition := LookupWordDefinition(wordToLookup)
	fmt.Println(definition)
}
