package http

import (
	"time"
)

// builderImpl is the implementation of the Builder interface and is used to build a client with the desired configuration.
type builderImpl struct {
	header  Headers
	Timeout Timeout
	builder Builder
	State   chan string
	client  *goHTTPClient
}

// Builder is the interface that wraps the basic Build method. The Build method returns a Client.
// The Client is used to make HTTP requests.
type Builder interface {
	SetRequestTimeout(timeout time.Duration) Timeout
	SetResponseTimeout(timeout time.Duration) Timeout
	SetMaxIdleConnections(maxConnections int) Timeout
	Headers() Headers
	Build() Client
}

// SetMaxIdleConnections sets the maximum number of idle (keep-alive) connections across all hosts.
// If zero, DefaultMaxIdleConnsPerHost is used.
func (c builderImpl) SetMaxIdleConnections(maxConnections int) Timeout {
	c.Timeout.SetMaxIdleConnections(maxConnections)
	return c.Timeout
}

// Headers returns the Headers object that is used to set the headers for the HTTP request.
// The Headers object is used to set the headers for the HTTP request.
func (c builderImpl) Headers() Headers {
	return c.header
}

// GetMaxIdleConnections returns the maximum number of idle (keep-alive) connections across all hosts.
// If zero, DefaultMaxIdleConnsPerHost is used.
func (c builderImpl) GetMaxIdleConnections() int {
	return c.Timeout.GetMaxIdleConnections()
}

// SetRequestTimeout sets the timeout for the HTTP request.
// If zero, no timeout exists.
// If negative, the request will not time out.
// If positive, the request will time out after the specified duration.
// The default is zero.
func (c builderImpl) SetRequestTimeout(timeout time.Duration) Timeout {
	c.Timeout.SetRequestTimeout(timeout)
	return c.Timeout
}

// SetResponseTimeout sets the timeout for the HTTP response.
// If zero, no timeout exists.
// If negative, the response will not time out.
// If positive, the response will time out after the specified duration.
func (c builderImpl) SetResponseTimeout(timeout time.Duration) Timeout {
	c.Timeout.SetResponseTimeout(timeout)
	return c.Timeout
}

// Build returns a Client that is used to make HTTP requests.
// The Client is used to make HTTP requests.
func (c builderImpl) Build() Client {
	if c.client == nil {
		c.client = &goHTTPClient{
			builder: &c,
		}
	}
	return c.client
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
