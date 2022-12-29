package examples

import (
	"time"
	
	"github.com/cploutarchou/go-requests/http"
)

var (
	jsonContentClient = getClient("application/json")
	xmlContentClient  = getClient("application/xml")
	baseURL           = "https://go-requests.wiremockapi.cloud"
)

func getClient(contentType string) http.Client {
	builder := http.NewBuilder()
	// set content type and accept to application/json
	builder.Headers().
		SetContentType(contentType).
		SetAccept(contentType)
	
	builder.SetRequestTimeout(10 * time.Second).
		SetResponseTimeout(10 * time.Second).
		SetMaxIdleConnections(10)
	return builder.Build()
}
