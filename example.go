package main

import (
	"fmt"
	"github.com/cploutarchou/go-http/http"
	"io"
)

var client http.GoHttpClient

func getGithubClient() http.GoHttpClient {
	_client := http.NewClient(http.DefaultConfig)
	commonHeaders := _client.MakeHeaders()
	commonHeaders.Add("Accept", "application/json")
	_client.SetHeaders(commonHeaders)
	return _client
}

func init() {
	client = getGithubClient()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func main() {
	//GetExample()
	user := User{
		FirstName: "Christos",
		LastName:  "Ploutarchou",
		Username:  "username",
	}
	PostExample(user)
	GetExample()
	GetExample()
	GetExample()
	GetExample()
	GetExample()
	GetExample()
}

func GetExample() {
	response, err := client.Get("https://api.github.com", nil)
	client.MakeHeaders()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func PostExample(u User) {
	response, err := client.Post("https://api.github.com", nil, u)
	client.MakeHeaders()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//bytes, err := io.ReadAll(response.Body)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(response.StatusCode)
}
