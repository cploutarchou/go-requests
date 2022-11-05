package http

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type Method string

var (
	DefaultConfig *Config
)

type goHTTPClient struct {
	Headers http.Header
	client  *http.Client
}
type Config struct {
	// Timeout is the timeout for HTTP response in nanoseconds
	ResponseTimeout time.Duration
	// Timeout is the timeout for HTTP requests in nanoseconds
	RequestTimeout time.Duration
	// MaxIdleConnections is the maximum number of connections in idle
	MaxIdleConnections int
}
type GoHTTPClient interface {
	// SetHeaders sets the headers for the request
	SetHeaders(http.Header)
	// MakeHeaders Returns the headers for the request
	MakeHeaders() http.Header
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header, interface{}) (*http.Response, error)
	Head(string, http.Header, interface{}) (*http.Response, error)
}

func NewClient(config *Config) GoHTTPClient {
	if config == nil {
		config = &Config{
			RequestTimeout:     2,
			MaxIdleConnections: 5,
			ResponseTimeout:    5,
		}
	}
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   config.MaxIdleConnections,
			ResponseHeaderTimeout: config.RequestTimeout * time.Second,
			DialContext: (&net.Dialer{
				Timeout: config.ResponseTimeout * time.Second,
			}).DialContext,
		},
	}
	return &goHTTPClient{
		client: &client,
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
