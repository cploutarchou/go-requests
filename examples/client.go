package examples

import (
	"time"

	"github.com/cploutarchou/go-requests/http"
)

var (
	client  = getClient()
	baseURL = "https://go-requests.wiremockapi.cloud"
)

func getClient() http.Client {
	builder := http.NewBuilder()
	// set content type and accept to application/json
	builder.Headers().
		SetContentType("application/json").
		SetAccept("application/json")

	builder.SetRequestTimeout(10 * time.Second).
		SetResponseTimeout(10 * time.Second).
		SetMaxIdleConnections(10)
	return builder.Build()
}
