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
type Config struct {
	// Timeout is the timeout for HTTP response in nanoseconds
	ResponseTimeout time.Duration
	// Timeout is the timeout for HTTP requests in nanoseconds
	RequestTimeout time.Duration
	// MaxIdleConnections is the maximum number of connections in idle
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
func (c *goHTTPClient) SetMaxIdleConnections(maxIdleConnections int) {

	c.Config.MaxIdleConnections = maxIdleConnections
}

func (c *goHTTPClient) SetConfig(config *Config) {
	c.Config = config
}
