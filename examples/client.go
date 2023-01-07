package examples

import (
	"net"
	"net/http"
	"time"

	"github.com/cploutarchou/go-requests/requests"
)

var (
	//jsonContentClient = getClient("application/json")
	xmlContentClient  = getClient("application/xml")
	jsonContentClient = getCustomClient()
	baseURL           = "https://go-requests.wiremockapi.cloud"
)

func getCustomClient() requests.Client {
	// create a custom http client that allows you to set your own http client set headers to accept and content type to application/json
	customClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		},
	}
	// set content type and accept to application/json
	builder := requests.NewBuilder()
	builder.Headers().
		SetContentType("application/json").
		SetAccept("application/json").SetUserAgent("go-requests")
	builder.SetHTTPClient(customClient)
	return builder.Build()
}

func getClient(contentType string) requests.Client {
	builder := requests.NewBuilder()
	// set content type and accept to application/json
	builder.Headers().
		SetContentType(contentType).
		SetAccept(contentType).
		SetUserAgent("go-requests")
	builder.SetRequestTimeout(10 * time.Second).
		SetResponseTimeout(10 * time.Second).
		SetMaxIdleConnections(10)

	return builder.Build()
}
