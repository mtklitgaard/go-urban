package main

import (
	"fmt"
	"os"

	"github.com/mtklitgaard/go-urban/api"
)

func main() {
	wordToLookup := os.Args[1]
	definition := api.LookupWordDefinition(wordToLookup)
	fmt.Println(definition)
}
