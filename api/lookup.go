package api

import (
	"encoding/json"
	"net/http"

	"github.com/mtklitgaard/go-urban/data"
)

// LookupWordDefinition looks up definition
func LookupWordDefinition(word string) (definition string) {
	resp, err := http.Get("http://api.urbandictionary.com/v0/define?term=" + word)
	if err != nil {
		definition = "Something went blegh!"
	}
	defer resp.Body.Close()

	var urbanResponse data.UrbanResponse

	if err := json.NewDecoder(resp.Body).Decode(&urbanResponse); err != nil {
		definition = "Something went wrong parsing shit"
	}
	definition = urbanResponse.List[0].Definition
	return
}
