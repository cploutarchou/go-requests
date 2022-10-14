package main

import (
	"fmt"
	"github.com/cploutarchou/go-http/http"
	"io"
)

func main() {
	client := http.NewClient()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
