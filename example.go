package main

import (
	"fmt"
	"io"
	"time"

	"github.com/cploutarchou/go-http/http"
)

var client http.GoHTTPClient

func getGithubClientWithOutConfig() http.GoHTTPClient {
	_client := http.NewClient()
	commonHeaders := _client.MakeHeaders()
	commonHeaders.Add("Accept", "application/json")
	_client.SetHeaders(commonHeaders)
	return _client
}

func getGithubClientWithConfig() http.GoHTTPClient {
	_client := http.NewClient()
	_client.SetConfig(&http.Config{
		MaxIdleConnections: 10,
		ResponseTimeout:    50 * time.Second,
		RequestTimeout:     50 * time.Second,
	})
	return _client
}

func getGithubClientBySetters() http.GoHTTPClient {
	_client := http.NewClient()
	_client.SetRequestTimeout(50 * time.Second)
	_client.SetResponseTimeout(50 * time.Second)
	_client.SetMaxIdleConnections(10)
	return _client
}
func init() {
	client = getGithubClientWithOutConfig()
	client = getGithubClientWithConfig()
	client = getGithubClientBySetters()
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
