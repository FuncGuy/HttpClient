package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() HttpClient {
	client := New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	headers := make(http.Header)
	response, err := githubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))

}
