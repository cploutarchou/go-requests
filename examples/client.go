package examples

import (
	"time"

	"github.com/cploutarchou/go-requests/http"
)

var (
	client = getClient()
	baseURL = "https://go-requests.mocklab.io/"
)

func getClient() http.Client {
	builder := http.NewBuilder()
	builder.SetRequestTimeout(50 * time.Second).
		SetResponseTimeout(50 * time.Second).
		SetMaxIdleConnections(10)
	return builder.Build()
}
