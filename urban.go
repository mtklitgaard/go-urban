package main

import (
	"fmt"

	"flag"

	"github.com/mtklitgaard/go-urban/api"
)

func main() {
	wordPtr := flag.String("word", "foo", "Word you want to lookup")
	countPtr := flag.Int("count", 5, "Number of definitions you would like to see")
	flag.Parse()
	wordToLookup := *wordPtr
	numberOfResults := *countPtr
	definitions, _ := api.LookupWordDefinition(wordToLookup)
	if numberOfResults != 0 && len(definitions.List) != 0 {
		fmt.Printf("\n Results for: %s \n", wordToLookup)
		for index := 1; index <= numberOfResults; index++ {
			if len(definitions.List) < index {
				break
			}
			result := definitions.List[index-1]
			fmt.Printf("\n\nResult: %v \n", index)
			fmt.Println(result.Definition)
			fmt.Println("Examples:")
			fmt.Println(result.Example)
		}
	} else {
		fmt.Println("You got no results")
	}
}
