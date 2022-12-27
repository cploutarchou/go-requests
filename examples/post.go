package examples

import (
	"fmt"
	"github.com/cploutarchou/go-requests/http"
)

func PostCreateUser(user User) (*http.Response, error) {
	res, err := client.Post(baseURL, nil, user)
	fmt.Print(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
