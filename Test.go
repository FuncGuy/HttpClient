package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func mainn() {

	httpMethod := "GET"

	url := "https://api.github.com"

	client := http.Client{}

	request, err := http.NewRequest(httpMethod, url, nil)

	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))

}
