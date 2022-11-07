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
	Headers http.Header
	client  *http.Client
	Config  *Config
}

// Config is a struct that holds the configuration for the http client
//
//	RequestTimeout: the timeout for the request
//	ResponseTimeout: the timeout for the response
//	MaxIdleConnections: the maximum number of idle connections
type Config struct {
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
	SetConfig(*Config)
}

func NewClient() GoHTTPClient {
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   defaultMaxIdleConnectionsPerHost,
			ResponseHeaderTimeout: defaultResponseTimeout,
			DialContext: (&net.Dialer{
				Timeout: defaultRequestTimeout,
			}).DialContext,
		},
	}
	return &goHTTPClient{
		client: &client,
		Config: &Config{},
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
	c.Config.RequestTimeout = timeout
}

func (c *goHTTPClient) SetResponseTimeout(timeout time.Duration) {
	c.Config.ResponseTimeout = timeout
}
func (c *goHTTPClient) SetMaxIdleConnections(maxConnections int) {
	c.Config.MaxIdleConnections = maxConnections
}

func (c *goHTTPClient) SetConfig(config *Config) {
	c.Config = config
}

func (c *goHTTPClient) getRequestTimeout() time.Duration {
	if c.Config.RequestTimeout == 0 {
		return defaultRequestTimeout
	}
	return c.Config.RequestTimeout
}

func (c *goHTTPClient) getResponseTimeout() time.Duration {
	if c.Config.ResponseTimeout == 0 {
		return defaultResponseTimeout
	}
	return c.Config.ResponseTimeout
}

func (c *goHTTPClient) getMaxIdleConnections() int {
	if c.Config.MaxIdleConnections == 0 {
		return defaultMaxIdleConnectionsPerHost
	}
	return c.Config.MaxIdleConnections
}
