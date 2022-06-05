package main

import (
	"io"
	"net/http"
)

func testhttp() {

	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if body != nil {

	}

}
