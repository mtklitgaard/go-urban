package api

import (
	"encoding/json"
	"net/http"

	"github.com/mtklitgaard/go-urban/data"
)

// LookupWordDefinition looks up definition
func LookupWordDefinition(word string) (urbanResponse data.UrbanResponse, err error) {
	resp, err := http.Get("http://api.urbandictionary.com/v0/define?term=" + word)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if errorInResponse := json.NewDecoder(resp.Body).Decode(&urbanResponse); errorInResponse != nil {
		err = errorInResponse
		return
	}
	return
}
