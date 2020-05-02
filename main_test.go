package main

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {

	is := assert.New(t)
	resp, err := search("Akon")
	is.NoError(err)
	is.Equal("ok", resp.Status)
	is.True(len(resp.Songs) > 1)
}

func TestParseResponse(t *testing.T) {

	is := assert.New(t)
	htmlStr, err := parseResponse([]Song{
		{
			Title:  "Some Title 1",
			Artist: "Some Artist 1",
		},
		{
			Title:  "Some Title 2",
			Artist: "Some Artist 2",
		},
	})

	is.NoError(err)

	_, err = html.Parse(strings.NewReader(htmlStr))

	is.NoError(err)

}
