package http

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type Method string

const (
	defaultMaxIdleConnectionsPerHost = 10
	defaultResponseTimeout           = 5 * time.Second
	defaultRequestTimeout            = 5 * time.Second
)

type goHTTPClient struct {
	Headers         http.Header
	client          *http.Client
	TimeoutSettings *TimeoutSettings
}

// TimeoutSettings is a struct that holds the configuration for the http client
//
//	RequestTimeout: the timeout for the request
//	ResponseTimeout: the timeout for the response
//	MaxIdleConnections: the maximum number of idle connections
type TimeoutSettings struct {
	ResponseTimeout    time.Duration
	RequestTimeout     time.Duration
	MaxIdleConnections int
}

// GoHTTPClient is an interface for http client
type GoHTTPClient interface {
	SetHeaders(http.Header)
	MakeHeaders() http.Header
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header, interface{}) (*http.Response, error)
	Head(string, http.Header, interface{}) (*http.Response, error)
	SetRequestTimeout(time.Duration)
	SetResponseTimeout(time.Duration)
	SetMaxIdleConnections(int)
	getRequestTimeout() time.Duration
	getResponseTimeout() time.Duration
	getMaxIdleConnections() int
	SetConfig(*TimeoutSettings)
}

func NewClient() GoHTTPClient {
	var client = &goHTTPClient{}
	return &goHTTPClient{
		client:          client.getClient(),
		TimeoutSettings: &TimeoutSettings{},
	}
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
func (c *goHTTPClient) SetHeaders(h http.Header) {
	c.Headers = h
}

func (c *goHTTPClient) MakeHeaders() http.Header {
	return make(http.Header)
}

func (c *goHTTPClient) Get(url string, headers http.Header) (*http.Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	response, err := c.do(http.MethodGet, url, headers, data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	response, err := c.do(http.MethodGet, url, headers, data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Delete(url string, headers http.Header, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	response, err := c.do(http.MethodGet, url, headers, data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	response, err := c.do(http.MethodGet, url, headers, data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Head(url string, headers http.Header, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	response, err := c.do(http.MethodGet, url, headers, data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) SetRequestTimeout(timeout time.Duration) {
	c.TimeoutSettings.RequestTimeout = timeout
}

func (c *goHTTPClient) SetResponseTimeout(timeout time.Duration) {
	c.TimeoutSettings.ResponseTimeout = timeout
}
func (c *goHTTPClient) SetMaxIdleConnections(maxConnections int) {
	c.TimeoutSettings.MaxIdleConnections = maxConnections
}

func (c *goHTTPClient) SetConfig(config *TimeoutSettings) {
	c.TimeoutSettings = config
}

// getRequestTimeout returns the request timeout
func (c *goHTTPClient) getRequestTimeout() time.Duration {
	if c.TimeoutSettings.RequestTimeout != defaultRequestTimeout {
		return c.TimeoutSettings.RequestTimeout
	}
	return defaultRequestTimeout
}

func (c *goHTTPClient) getResponseTimeout() time.Duration {
	if c.TimeoutSettings.ResponseTimeout != defaultResponseTimeout {
		return c.TimeoutSettings.ResponseTimeout
	}
	return defaultResponseTimeout
}

func (c *goHTTPClient) getMaxIdleConnections() int {
	if c.TimeoutSettings.MaxIdleConnections != defaultMaxIdleConnectionsPerHost {
		return c.TimeoutSettings.MaxIdleConnections
	}
	return defaultMaxIdleConnectionsPerHost
}

// getClient returns the *http.client if exist
// or creates a new one with the default settings
// and returns it.
// The default settings are:
//   - MaxIdleConnectionsPerHost: 10
//   - RequestTimeout: 5 seconds
//   - ResponseTimeout: 5 seconds
func (c *goHTTPClient) getClient() *http.Client {
	if c.client != nil {
		return c.client
	}
	c.TimeoutSettings = &TimeoutSettings{
		RequestTimeout:     defaultRequestTimeout,
		ResponseTimeout:    defaultResponseTimeout,
		MaxIdleConnections: defaultMaxIdleConnectionsPerHost,
	}
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
			ResponseHeaderTimeout: c.getResponseTimeout(),
			DialContext: (&net.Dialer{
				Timeout: c.getRequestTimeout(),
			}).DialContext,
		},
	}
	c.client = &client
	return c.client
}
