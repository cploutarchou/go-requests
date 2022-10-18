package main

import (
	"fmt"
	"io"

	"github.com/cploutarchou/go-http/http"
)

var client http.Client

func getGithubClient() http.Client {
	_client := http.NewClient()
	commonHeaders := _client.MakeHeaders()
	commonHeaders.Add("Accept", "application/json")
	_client.SetHeaders(commonHeaders)
	return _client
}

func init() {
	client = getGithubClient()
}

func main() {
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
