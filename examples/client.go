package examples

import (
	"github.com/cploutarchou/go-requests"
	"net"
	"net/http"
	"time"
)

var (
	//jsonContentClient = getClient("application/json")
	xmlContentClient  = getClient("application/xml")
	jsonContentClient = getCustomClient()
	baseURL           = "https://go-requests.wiremockapi.cloud"
)

func getCustomClient() go_requests.Client {
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
	builder := go_requests.NewBuilder()
	builder.Headers().
		SetContentType("application/json").
		SetAccept("application/json").SetUserAgent("go-requests")
	builder.SetHTTPClient(customClient)
	return builder.Build()
}

func getClient(contentType string) go_requests.Client {
	builder := go_requests.NewBuilder()
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
