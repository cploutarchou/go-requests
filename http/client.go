package http

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type Method string

const (
	// defaultMaxIdleConnectionsPerHost is the default value for MaxIdleConnectionsPerHost
	defaultMaxIdleConnectionsPerHost = 10
	// defaultResponseTimeout is the default value for ResponseTimeout
	defaultResponseTimeout = 5 * time.Second
	// defaultRequestTimeout is the default value for RequestTimeout
	defaultRequestTimeout = 5 * time.Second
)

// goHTTPClient is the default implementation of the GoHTTPClient interface
// it is used to make http requests
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
	DisableTimeouts    bool
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
	DisableTimeouts(bool)
}

// NewClient creates a new http client
// returns a new http client
//
// Example:
//
//		client := NewClient()
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
func (c *goHTTPClient) MakeHeaders() http.Header {
	return make(http.Header)
}

// Get sends a GET request to the specified URL
// url: the url to send the request to
// headers: the headers to be sent with the request
// returns the response and an error if there is one
// returns an error if the request fails
//
//	Example:
//		response, err := client.Get("https://www.google.com", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer response.Body.Close()
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//	log.Fatal(err)
//	}
//	fmt.Println(string(body))
func (c *goHTTPClient) Get(url string, headers http.Header) (*http.Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Post sends a POST request to the specified URL
// url: the url to send the request to
//
//	headers: the headers to be sent with the request
//	body: the body to be sent with the request
//	returns the response and an error if there is one
//	returns an error if the request fails
//	Example:
//		response, err := client.Post("https://www.google.com", nil, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer response.Body.Close()
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(body))
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

// Put sends a PUT request to the specified URL
// url: the url to send the request to
// headers: the headers to be sent with the request
// body: the body to be sent with the request
// returns the response and an error if there is one
// returns an error if the request fails
//
//	Example:
//		response, err := client.Put("https://www.google.com", nil, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer response.Body.Close()
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//	log.Fatal(err)
//	}
//	fmt.Println(string(body))
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

// Delete sends a DELETE request to the specified URL
// url: the url to send the request to
// headers: the headers to be sent with the request
//
//	Example:
//	response, err := client.Delete("https://www.google.com", nil, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer response.Body.Close()
//		body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(body))
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

// Patch sends a PATCH request to the specified URL
// url: the url to send the request to
// headers: the headers to be sent with the request
//
//	Example:
//		headers := make(http.Header)
//		headers.Set("Content-Type", "application/json")
//		headers.Set("Authorization", "Bearer <token>")
//		client.SetHeaders(headers)
//
// body: the body to be sent with the request
// returns the response and an error if there is one
// returns an error if the request fails
//
//		Example:
//			type User struct {
//				FirstName string `json:"first_name"`
//				LastName  string `json:"last_name"`
//			}
//			user := User{
//				FirstName: "Christos",
//				LastName: "Ploutarchou",
//			}
//			response, err := client.Patch("https://example.com", headers, user)
//			if err != nil {
//				log.Fatal(err)
//			}
//			defer response.Body.Close()
//		body, err := ioutil.ReadAll(response.Body)
//		if err != nil {
//			log.Fatal(err)
//	}
//		fmt.Println(string(body))
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

// Head sends a HEAD request to the specified URL
// url: the url to send the request to
// headers: the headers to be sent with the request
// body: the body to be sent with the request
// returns the response and an error if there is one
// returns an error if the request fails
//
//	Example:
//		response, err := client.Head("https://www.google.com", nil, nil)
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

// SetRequestTimeout sets the request timeout
// requestTimeout: the request timeout
//
//	Example:
//		client.SetRequestTimeout(5 * time.Second)
func (c *goHTTPClient) SetRequestTimeout(timeout time.Duration) {
	c.TimeoutSettings.RequestTimeout = timeout
}

// SetResponseTimeout sets the response timeout
// responseTimeout: the response timeout
//
//	Example:
//		client.SetResponseTimeout(5 * time.Second)
func (c *goHTTPClient) SetResponseTimeout(timeout time.Duration) {
	c.TimeoutSettings.ResponseTimeout = timeout
}

// SetMaxIdleConnections sets the maximum number of idle connections
// maxIdleConnections: the maximum number of idle connections
//
//	Example:
//		client.SetMaxIdleConnections(10)
func (c *goHTTPClient) SetMaxIdleConnections(maxConnections int) {
	c.TimeoutSettings.MaxIdleConnections = maxConnections
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
func (c *goHTTPClient) SetConfig(config *TimeoutSettings) {
	c.TimeoutSettings = config
}

// getRequestTimeout returns the request timeout
// if the request timeout is not set, it returns the default request timeout.
func (c *goHTTPClient) getRequestTimeout() time.Duration {
	if c.TimeoutSettings.RequestTimeout != defaultRequestTimeout {
		return c.TimeoutSettings.RequestTimeout
	}
	if c.TimeoutSettings.DisableTimeouts {
		return 0
	}
	return defaultRequestTimeout
}

// getResponseTimeout returns the response timeout
// if the request timeout is not set, it returns the default response timeout
func (c *goHTTPClient) getResponseTimeout() time.Duration {
	if c.TimeoutSettings.ResponseTimeout != defaultResponseTimeout {
		return c.TimeoutSettings.ResponseTimeout
	}
	if c.TimeoutSettings.DisableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

// getMaxIdleConnections()
//
// Returns the maximum number of idle connections
// if the value is not set, it returns the default value.
//
// default value is 10
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
		Timeout: c.getRequestTimeout(),
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

// DisableTimeouts disables the timeouts for the client
// if the value is true, the timeouts will be disabled
// and the client will not time out.
//
//	Example:
//		client.DisableTimeouts(true)
func (c *goHTTPClient) DisableTimeouts(disable bool) {
	c.TimeoutSettings.DisableTimeouts = disable
}
