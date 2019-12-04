package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func requestAPI(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func postAPI(url string, values url.Values, target interface{}) error {
	jsonBody, err := json.Marshal(values)
	r, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	/*if err != nil {
		return err
	}
	defer r.Body.Close()*/

	return json.NewDecoder(r.Body).Decode(target)
}