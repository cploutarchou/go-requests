package http

import "time"

const (
	// defaultMaxIdleConnectionsPerHost is the default value for MaxIdleConnectionsPerHost
	defaultMaxIdleConnectionsPerHost = 10
	// defaultResponseTimeout is the default value for ResponseTimeout
	defaultResponseTimeout = 5 * time.Second
	// defaultRequestTimeout is the default value for RequestTimeout
	defaultRequestTimeout = 5 * time.Second
)

// timeoutImpl is a struct that holds the configuration for the http client
//
//	RequestTimeout: the timeout for the request
//	ResponseTimeout: the timeout for the response
//	MaxIdleConnections: the maximum number of idle connections
type timeoutImpl struct {
	ResponseTimeout    time.Duration
	RequestTimeout     time.Duration
	MaxIdleConnections int
	DisableTimeouts    bool
}

func newTimeoutImpl() *timeoutImpl {
	return &timeoutImpl{
		ResponseTimeout:    defaultResponseTimeout,
		RequestTimeout:     defaultRequestTimeout,
		MaxIdleConnections: defaultMaxIdleConnectionsPerHost,
		DisableTimeouts:    false,
	}
}

type Timeout interface {
	SetRequestTimeout(time.Duration) Timeout
	SetResponseTimeout(time.Duration) Timeout
	SetMaxIdleConnections(int) Timeout
	GetMaxIdleConnections() int
	GetRequestTimeout() time.Duration
	GetResponseTimeout() time.Duration
	Disable(bool) Timeout
	Build()
}

// GetRequestTimeout returns the request timeout
// if the request timeout is not set, it returns the default request timeout.
func (c timeoutImpl) GetRequestTimeout() time.Duration {
	if c.RequestTimeout != defaultRequestTimeout {
		return c.RequestTimeout
	}
	if c.DisableTimeouts {
		return 0
	}
	return defaultRequestTimeout
}

// GetResponseTimeout returns the response timeout
// if the request timeout is not set, it returns the default response timeout
func (c timeoutImpl) GetResponseTimeout() time.Duration {
	if c.ResponseTimeout != defaultResponseTimeout {
		return c.ResponseTimeout
	}
	if c.DisableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

// GetMaxIdleConnections
//
// Returns the maximum number of idle connections
// if the value is not set, it returns the default value.
//
// default value is 10
func (c timeoutImpl) GetMaxIdleConnections() int {
	if c.MaxIdleConnections != defaultMaxIdleConnectionsPerHost {
		return c.MaxIdleConnections
	}
	return defaultMaxIdleConnectionsPerHost
}

// Disable disables the timeouts for the client
// if the value is true, the timeouts will be disabled
// and the client will not time out.
//
//	Example:
//		client.DisableTimeouts(true)
func (c timeoutImpl) Disable(disable bool) Timeout {
	c.DisableTimeouts = disable
	return c
}
