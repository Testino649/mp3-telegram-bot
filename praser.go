package main

import (
	"bytes"
	"html/template"
)

func parseResponse(songs []Song) (string, error) {

	tpl, err := template.ParseFiles("response.gohtml")

	if err != nil {
		return "", err
	}

	var tmplBytes bytes.Buffer

	err = tpl.Execute(&tmplBytes, songs)
	if err != nil {
		return "", err
	}

	return tmplBytes.String(), nil
}
