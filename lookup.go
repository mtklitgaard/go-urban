package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// UrbanResponse response from urban dictionary
type UrbanResponse struct {
	Tags       []string `json:"tags"`
	ResultType string   `json:"result_type"`
	List       []struct {
		Definition  string `json:"definition"`
		Permalink   string `json:"permalink"`
		ThumbsUp    int    `json:"thumbs_up"`
		Author      string `json:"author"`
		Word        string `json:"word"`
		Defid       int    `json:"defid"`
		CurrentVote string `json:"current_vote"`
		Example     string `json:"example"`
		ThumbsDown  int    `json:"thumbs_down"`
	} `json:"list"`
	Sounds []string `json:"sounds"`
}

// LookupWordDefinition looks up definition
func LookupWordDefinition(word string) (definition string) {
	resp, err := http.Get("http://api.urbandictionary.com/v0/define?term=word")
	if err != nil {
		definition = "Something went blegh!"
	}
	defer resp.Body.Close()

	var urbanResponse UrbanResponse

	if err := json.NewDecoder(resp.Body).Decode(&urbanResponse); err != nil {
		log.Println(err)
	}
	definition = urbanResponse.List[0].Definition
	return
}
