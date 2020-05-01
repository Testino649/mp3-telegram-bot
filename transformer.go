package main

// Contains hardcoded messages
const (
	apiErrorMessage    = "Unable to perform search, please try later."
	apiQueryMessage    = "No result found, please try with different query."
	apiParseErrMessage = "Opps! Something went wrong, please try later."
	hello              = "Hello %s, please try <b>/help</b> for help."
	help               = `<b>/search your_song_or_artist</b> for searching your track. 
			Example : 
			/search <i>The Weeknd</i>	`

	allReply = "Made with ♥, try <b>/search<b>"
)

type Response struct {
	Status string `json:"status"`
	Songs  []Song `json:"data"`
}
type Song struct {
	SourceID string `json:"source_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Date     int    `json:"date"`
	GenreID  int    `json:"genre_id"`
	Download string `json:"download"`
	Stream   string `json:"stream"`
	Cover    string `json:"cover"`
}
