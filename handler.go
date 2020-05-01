package main

import "log"

func getReply(q string) string {

	var err error

	resp, err := search(q)

	if err != nil {
		return apiErrorMessage
	}

	if len(resp.Songs) == 0 {
		return apiQueryMessage
	}

	// Fetching only top 10 results
	top10 := 11
	if top10 > len(resp.Songs) {
		top10 = len(resp.Songs)
	}
	message, err := parseResponse(resp.Songs[:top10-1])

	if err != nil {
		log.Print(err)
		return apiParseErrMessage
	}

	return message
}
