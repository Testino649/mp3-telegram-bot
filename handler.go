package main

import "log"

func getReply(q string) string {

	var err error

	resp, err := search(q)

	if err != nil {
		return apiErrorMessage
	}

	if len(resp.Songs) == 0 || resp.Status != "ok" {
		return apiQueryMessage
	}

	// Fetching only top 10 results
	top15 := 15
	if top15 > len(resp.Songs) {
		top15 = len(resp.Songs)
	}

	message, err := parseResponse(resp.Songs[:top15])

	if err != nil {
		log.Print(err)
		return apiParseErrMessage
	}

	return message
}
