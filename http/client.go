package http

import (
	"encoding/json"
	"net"
	"net/http"
)

type Method string

// goHTTPClient is the default implementation of the Client interface
// it is used to make http requests
type goHTTPClient struct {
	client  *http.Client
	header  http.Header
	timeout Timeout
}

// Client is an interface for http client
type Client interface {
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header, interface{}) (*http.Response, error)
	Head(string, http.Header, interface{}) (*http.Response, error)
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
	client := http.Client{
		Timeout: c.timeout.GetRequestTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.timeout.GetMaxIdleConnections(),
			ResponseHeaderTimeout: c.timeout.GetResponseTimeout(),
			DialContext: (&net.Dialer{
				Timeout: c.timeout.GetRequestTimeout(),
			}).DialContext,
		},
	}
	c.client = &client
	return c.client
}
