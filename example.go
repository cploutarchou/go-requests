package main

import (
	"fmt"
	"time"

	"github.com/cploutarchou/go-requests/http"
)

var client http.Client

func getGithubClientWithOutConfig() http.Client {
	builder := http.NewBuilder()
	builder.SetRequestTimeout(50 * time.Second).
		SetResponseTimeout(50 * time.Second).
		SetMaxIdleConnections(10)
	return builder.Build()
}

func getGithubClientBySetters() http.Client {
	_client := http.NewBuilder()

	_client.
		SetRequestTimeout(50 * time.Second).
		SetResponseTimeout(50 * time.Second).
		SetMaxIdleConnections(10)

	_client.Headers().
		SetAcceptCharset("utf-8").
		SetAccept("application/json")
	return _client.Build()
}

func init() {
	// client = getGithubClientWithOutConfig()
	client = getGithubClientBySetters()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func main() {
	user := User{
		FirstName: "Christos",
		LastName:  "Ploutarchou",
		Username:  "username",
	}
	PostExample(user)
	GetExample()
	// Test concurrency safety
	for i := 0; i < 100; i++ {
		go GetExample()
	}
	time.Sleep(10 * time.Second)
}

func GetExample() {
	client.DisableTimeouts()
	client.QueryParams().Set("q", "go-requests")
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.String())
	fmt.Println(response.StatusCode())
	fmt.Println(response.Status())
}

func PostExample(u User) {
	response, err := client.Post("https://api.github.com", nil, u)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.String())
	fmt.Println(response.StatusCode())
	fmt.Println(response.Status())
}
