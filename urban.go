package main

import (
	"fmt"

	"flag"

	"github.com/mtklitgaard/go-urban/api"
)

func main() {
	countPtr := flag.Int("count", 5, "Number of definitions you would like to see")
	showExamplesPtr := flag.Bool("examples", false, "Show examples if you want them.")
	flag.Parse()
	args := flag.Args()
	wordFromCommand := args[0]
	numberOfResults := *countPtr
	showExamples := *showExamplesPtr

	definitions, _ := api.LookupWordDefinition(wordFromCommand)

	if numberOfResults != 0 && len(definitions.List) != 0 {
		fmt.Printf("\nResults for: %s \n", wordFromCommand)
		for index := 1; index <= numberOfResults; index++ {
			if len(definitions.List) < index {
				break
			}
			result := definitions.List[index-1]
			fmt.Printf("\n\nResult: %v \n", index)
			fmt.Println(result.Definition)
			printExamples(showExamples, result.Example)
		}
	} else {
		fmt.Println("You got no results")
	}
}

func printExamples(showExamples bool, example string) {
	if showExamples {
		fmt.Println("Examples:")
		fmt.Println(example)
	}
}
