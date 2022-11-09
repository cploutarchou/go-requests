package http

import (
	"net/http"
	"time"
)

type builderImpl struct {
	header          http.Header
	timeoutSettings TimeoutSettings
}

type Builder interface {
	GetMaxIdleConnections() int
	SetConfig(TimeoutSettings) Builder

	SetHeaders(http.Header)
	MakeHeaders() http.Header
	Build() Client
}

// TimeoutSettings is a struct that holds the configuration for the http client
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
type Timeout interface {
	SetRequestTimeout(time.Duration) Builder
	SetResponseTimeout(time.Duration) Builder
	SetMaxIdleConnections(int) Builder
	GetRequestTimeout() time.Duration
	GetResponseTimeout() time.Duration
	Disable(bool) Builder
}

// SetHeaders sets the headers for the client
//
//	 headers: the headers to be set
//
//		Example:
//			headers := make(http.Header)
//			headers.Set("Content-Type", "application/json")
//			headers.Set("Authorization", "Bearer <token>")
//			client.SetHeaders(headers)
func (c builderImpl) SetHeaders(h http.Header) {
	c.header = h
}

// MakeHeaders makes the headers for the client
// returns the headers
// returns an error if the request fails
// Example:
//
//		headers := client.MakeHeaders()
//		headers.Set("Content-Type", "application/json")
//		headers.Set("Authorization", "Bearer <token>")
//		client.SetHeaders(headers)
//		response, err := client.Get("https://www.google.com", nil)
//	if err != nil {
//	log.Fatal(err)
//	}
//	defer response.Body.Close()
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//	log.Fatal(err)
//	}
//	fmt.Println(string(body))
func (c builderImpl) MakeHeaders() http.Header {
	return make(http.Header)
}

// SetRequestTimeout sets the request timeout
// requestTimeout: the request timeout
//
//	Example:
//		client.SetRequestTimeout(5 * time.Second)
func (c timeoutImpl) SetRequestTimeout(timeout time.Duration) Builder {
	c.timeoutSettings.RequestTimeout = timeout
	return c
}

// SetResponseTimeout sets the response timeout
// responseTimeout: the response timeout
//
//	Example:
//		client.SetResponseTimeout(5 * time.Second)
func (c timeoutImpl) SetResponseTimeout(timeout time.Duration) Builder {
	c.timeoutSettings.ResponseTimeout = timeout
	return c
}

// SetMaxIdleConnections sets the maximum number of idle connections
// maxIdleConnections: the maximum number of idle connections
//
//	Example:
//		client.SetMaxIdleConnections(10)
func (c timeoutImpl) SetMaxIdleConnections(maxConnections int) Builder {
	c.timeoutSettings.MaxIdleConnections = maxConnections
	return c
}

// SetConfig sets the configuration for the http client
// config: the configuration to be set
//
// Example:
//
//	config := &TimeoutSettings{
//		RequestTimeout:     5 * time.Second,
//		ResponseTimeout:    5 * time.Second,
//		MaxIdleConnections: 10,
//	}
//	client.SetConfig(config)
func (c builderImpl) SetConfig(config TimeoutSettings) Builder {
	c.timeoutSettings = config
	return c
}

// GetRequestTimeout returns the request timeout
// if the request timeout is not set, it returns the default request timeout.
func (c timeoutImpl) GetRequestTimeout() time.Duration {
	if c.timeoutSettings.RequestTimeout != defaultRequestTimeout {
		return c.timeoutSettings.RequestTimeout
	}
	if c.timeoutSettings.DisableTimeouts {
		return 0
	}
	return defaultRequestTimeout
}

// GetResponseTimeout returns the response timeout
// if the request timeout is not set, it returns the default response timeout
func (c timeoutImpl) GetResponseTimeout() time.Duration {
	if c.timeoutSettings.ResponseTimeout != defaultResponseTimeout {
		return c.timeoutSettings.ResponseTimeout
	}
	if c.timeoutSettings.DisableTimeouts {
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
	if c.timeoutSettings.MaxIdleConnections != defaultMaxIdleConnectionsPerHost {
		return c.timeoutSettings.MaxIdleConnections
	}
	return defaultMaxIdleConnectionsPerHost
}

// DisableTimeouts disables the timeouts for the client
// if the value is true, the timeouts will be disabled
// and the client will not time out.
//
//	Example:
//		client.DisableTimeouts(true)
func (c timeoutImpl) Disable(disable bool) Builder {
	c.timeoutSettings.DisableTimeouts = disable
	return c
}

func (c builderImpl) Build() Client {
	return &goHTTPClient{
		timeoutSettings: c.timeoutSettings,
		header:          c.header,
	}
}

// NewBuilder creates a new http client
// returns a new http client
//
// Example:
//
//		client := NewBuilder()
//		response, err := client.Get("https://www.google.com", nil)
//	if err != nil {
//	log.Fatal(err)
//	}
//	defer response.Body.Close()
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//	log.Fatal(err)
//	}
//	fmt.Println(string(body))
func NewBuilder() Builder {
	builder := &builderImpl{
		timeoutSettings: TimeoutSettings{
			RequestTimeout:     defaultRequestTimeout,
			ResponseTimeout:    defaultResponseTimeout,
			MaxIdleConnections: defaultMaxIdleConnectionsPerHost,
		},
	}
	return builder
}
