package data

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
