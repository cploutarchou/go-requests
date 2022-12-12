package http

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type Method string

// goHTTPClient is the default implementation of the Client interface
// it is used to make http requests
type goHTTPClient struct {
	client  *http.Client
	Headers Headers
	Timeout Timeout
}

// Client is an interface for http client
type Client interface {
	DisableTimeouts()
	EnableTimeouts()
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header, interface{}) (*http.Response, error)
	Head(string, http.Header, interface{}) (*http.Response, error)
}

// Get sends a GET request to the specified URL
// url: the url to send the request to
// Headers: the Headers to be sent with the request
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
//	Headers: the Headers to be sent with the request
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
// Headers: the Headers to be sent with the request
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
// Headers: the Headers to be sent with the request
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
// Headers: the Headers to be sent with the request
//
//	Example:
//		Headers := make(http.Header)
//		Headers.Set("Content-Type", "application/json")
//		Headers.Set("Authorization", "Bearer <token>")
//		client.Headers(Headers)
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
//			response, err := client.Patch("https://example.com", Headers, user)
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
// Headers: the Headers to be sent with the request
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

// DisableTimeouts disables the timeouts for the client requests
// Example:
//
//	client.DisableTimeouts()
//	response, err := client.Get("https://www.google.com", nil, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
func (c *goHTTPClient) DisableTimeouts() {
	c.Timeout = c.Timeout.Disable()
}

// EnableTimeouts enables the timeouts for the client requests
// Example:
//
//	client.EnableTimeouts()
//	response, err := client.Get("https://www.google.com", nil, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
func (c *goHTTPClient) EnableTimeouts() {
	c.Timeout = c.Timeout.Enable()
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
	fmt.Println("request timeout: ", c.Timeout.GetRequestTimeout())
	client := http.Client{
		Timeout: c.Timeout.GetRequestTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.Timeout.GetMaxIdleConnections(),
			ResponseHeaderTimeout: c.Timeout.GetResponseTimeout(),
			DialContext: (&net.Dialer{
				Timeout: c.Timeout.GetRequestTimeout(),
			}).DialContext,
		},
	}
	c.client = &client
	return c.client
}
