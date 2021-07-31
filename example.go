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

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))

}
