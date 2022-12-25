package examples

import (
	"time"

	"github.com/cploutarchou/go-requests/http"
)

var (
	client  = getClient()
	baseURL = "https://63a8d253f4962215b588da7e.mockapi.io/api/v1/users/"
)

func getClient() http.Client {
	builder := http.NewBuilder()
	// set content type and accept to application/json
	builder.Headers().
		SetContentType("application/json").
		SetAccept("application/json")

	builder.SetRequestTimeout(50 * time.Second).
		SetResponseTimeout(50 * time.Second).
		SetMaxIdleConnections(10)
	return builder.Build()
}
