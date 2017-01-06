package main

import (
	"fmt"
	"os"
)

func main() {
	wordToLookup := os.Args[1]
	fmt.Println(wordToLookup)
}
