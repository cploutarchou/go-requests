package requests

import (
	"net"
	"net/http"
	"sync"
)

type Method string

// goHTTPClient is the default implementation of the Client interface
// it is used to make http requests
type goHTTPClient struct {
	builder     *builderImpl
	client      *http.Client
	clientOnce  sync.Once
	queryParams QueryParams
}

func (c *goHTTPClient) QueryParams() QueryParams {
	if c.queryParams == nil {
		c.queryParams = NewQueryParams()
	}
	return c.queryParams
}

// Client is an interface for http client
type Client interface {
	QueryParams() QueryParams
	DisableTimeouts()
	EnableTimeouts()
	Headers() Headers

	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body []byte, headers ...http.Header) (*Response, error)
	Put(url string, body []byte, headers ...http.Header) (*Response, error)
	Patch(url string, body []byte, headers ...http.Header) (*Response, error)
	Delete(url string, body []byte, headers ...http.Header) (*Response, error)
	Head(url string, body []byte, headers ...http.Header) (*Response, error)
}

func (c *goHTTPClient) Get(url string, headers ...http.Header) (*Response, error) {
	response, err := c.do(http.MethodGet, url, getHeader(headers...), nil)
	// restore timeout state to default in case it was disabled
	if c.builder.Timeout.GetRequestTimeout() == 0 {
		c.builder.Timeout = c.builder.Timeout.Enable()
	}
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Post(url string, body []byte, headers ...http.Header) (*Response, error) {
	response, err := c.do(http.MethodPost, url, getHeader(headers...), body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Put(url string, body []byte, headers ...http.Header) (*Response, error) {
	response, err := c.do(http.MethodPut, url, getHeader(headers...), body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *goHTTPClient) Delete(url string, body []byte, headers ...http.Header) (*Response, error) {
	response, err := c.do(http.MethodDelete, url, getHeader(headers...), body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Patch(url string, body []byte, headers ...http.Header) (*Response, error) {
	response, err := c.do(http.MethodPatch, url, getHeader(headers...), body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *goHTTPClient) Head(url string, body []byte, headers ...http.Header) (*Response, error) {
	response, err := c.do(http.MethodHead, url, getHeader(headers...), body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// DisableTimeouts disables the timeouts for the client requests
// Example:
//
//	client.DisableTimeouts()
//	response, err := client.Get("https://www.google.com", nil, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
func (c *goHTTPClient) DisableTimeouts() {
	c.builder.Timeout = c.builder.Timeout.Disable()
}

func (c *goHTTPClient) EnableTimeouts() {
	c.builder.Timeout = c.builder.Timeout.Enable()
}

// getClient returns the *http.client if exists or creates a new one and returns it.
func (c *goHTTPClient) getClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.builder.cstClient != nil {
			c.client = c.builder.cstClient
			return
		}
		c.client = &http.Client{
			Timeout: c.builder.Timeout.GetRequestTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.builder.Timeout.GetMaxIdleConnections(),
				ResponseHeaderTimeout: c.builder.Timeout.GetResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.builder.Timeout.GetRequestTimeout(),
				}).DialContext,
			},
		}
	})
	return c.client

}

// Headers sets the headers for the client
func (c *goHTTPClient) Headers() Headers {
	return c.builder.Headers()
}
