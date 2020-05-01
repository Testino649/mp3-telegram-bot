package main

import (
	resty "github.com/go-resty/resty/v2"
)

func search(searchQuery string) (*Response, error) {

	var datmusic Response
	r := resty.New()
	_, err := r.SetDebug(true).R().
		EnableTrace().
		SetHeader("user-agent", "FZ5YEZm0WEXUcBFECkD").
		SetHeader("host", "api.datmusic.xyz").
		SetQueryParams(map[string]string{
			"q": searchQuery,
			//"page": "0",
		}).
		SetResult(&datmusic).
		Get("https://api.datmusic.xyz/search")

	if err != nil {
		return nil, err
	}

	return &datmusic, nil
}
