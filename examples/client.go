package examples

import (
	"github.com/cploutarchou/go-requests/http"
	"time"
)

var (
	client = getClient()
)

func getClient() http.Client {
	builder := http.NewBuilder()
	builder.SetRequestTimeout(50 * time.Second).
		SetResponseTimeout(50 * time.Second).
		SetMaxIdleConnections(10)
	return builder.Build()
}
