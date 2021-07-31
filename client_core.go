package main

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)

	response, err := client.Do(request)

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	return response, err

}
