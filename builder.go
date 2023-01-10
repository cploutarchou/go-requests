package go_requests

import (
	"net/http"
	"time"
)

// builderImpl is the implementation of the Builder interface and is used to build a client with the desired configuration.
type builderImpl struct {
	header    Headers
	Timeout   Timeout
	State     chan string
	client    *goHTTPClient
	cstClient *http.Client
}

// Builder is the interface that wraps the basic Build method. The Build method returns a Client.
// The Client is used to make HTTP requests.
type Builder interface {
	//SetRequestTimeout sets the timeout for the HTTP request.
	SetRequestTimeout(timeout time.Duration) Timeout
	//SetResponseTimeout sets the timeout for the HTTP response.
	SetResponseTimeout(timeout time.Duration) Timeout
	//SetMaxIdleConnections sets the maximum number of idle (keep-alive) connections across all hosts.
	SetMaxIdleConnections(maxConnections int) Timeout
	//Headers returns the Headers object that is used to set the headers for the HTTP request.
	Headers() Headers
	//Build returns a Client that is used to make HTTP requests.
	Build() Client
	//SetHTTPClient sets the http client to be used for the request instead of the default one.
	SetHTTPClient(*http.Client)
}

// SetMaxIdleConnections sets the maximum number of idle (keep-alive) connections across all hosts.
// If zero, DefaultMaxIdleConnections per  host is used.
func (b *builderImpl) SetMaxIdleConnections(maxConnections int) Timeout {
	b.Timeout.SetMaxIdleConnections(maxConnections)
	return b.Timeout
}

// Headers returns the Headers object that is used to set the headers for the HTTP request.
// The Headers object is used to set the headers for the HTTP request.
func (b *builderImpl) Headers() Headers {
	return b.header
}

// GetMaxIdleConnections returns the maximum number of idle (keep-alive) connections across all hosts.
// If zero, DefaultMaxIdleConnection Per Host is used.
func (b *builderImpl) GetMaxIdleConnections() int {
	return b.Timeout.GetMaxIdleConnections()
}

// SetRequestTimeout sets the timeout for the HTTP request.
// If zero, no timeout exists.
// If negative, the request will not time out.
// If positive, the request will time out after the specified duration.
// The default is zero.
func (b *builderImpl) SetRequestTimeout(timeout time.Duration) Timeout {
	b.Timeout.SetRequestTimeout(timeout)
	return b.Timeout
}

// SetResponseTimeout sets the timeout for the HTTP response.
// If zero, no timeout exists.
// If negative, the response will not time out.
// If positive, the response will time out after the specified duration.
func (b *builderImpl) SetResponseTimeout(timeout time.Duration) Timeout {
	b.Timeout.SetResponseTimeout(timeout)
	return b.Timeout
}

// SetHTTPClient sets the http client to be used for the request instead of the default one.
func (b *builderImpl) SetHTTPClient(c *http.Client) {
	if c != nil {
		b.cstClient = c
	}
	return
}

// Build returns a Client that is used to make HTTP requests.
// The Client is used to make HTTP requests.
func (b *builderImpl) Build() Client {
	if b.client == nil {
		b.client = &goHTTPClient{
			builder: b,
		}
	}
	return b.client
}

// NewBuilder returns a new Builder.
// The Builder is used to build a client with the desired configuration.
func NewBuilder() Builder {
	builder := &builderImpl{
		Timeout: newTimeouts(),
		header:  NewHeaders(),
		State:   make(chan string, 100),
	}
	return builder
}
