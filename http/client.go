package http

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
)

type Method string

// goHTTPClient is the default implementation of the Client interface
// it is used to make http requests
type goHTTPClient struct {
	builder    *builderImpl
	client     *http.Client
	clientOnce sync.Once
	params     *QueryParams
}

func (c *goHTTPClient) QueryParams() QueryParams {
	return *c.params
}

// Client is an interface for http client
type Client interface {
	QueryParams() QueryParams
	DisableTimeouts()
	EnableTimeouts()
	Get(string, http.Header) (*Response, error)
	Post(string, http.Header, interface{}) (*Response, error)
	Put(string, http.Header, interface{}) (*Response, error)
	Patch(string, http.Header, interface{}) (*Response, error)
	Delete(string, http.Header, interface{}) (*Response, error)
	Head(string, http.Header, interface{}) (*Response, error)
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
//	body, err := util.ReadAll(response.Body)
//	if err != nil {
//	log.Fatal(err)
//	}
//	fmt.Println(string(body))
func (c *goHTTPClient) Get(url string, headers http.Header) (*Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	// restore timeout state to default in case it was disabled
	if c.builder.Timeout.GetRequestTimeout() == 0 {
		c.builder.Timeout = c.builder.Timeout.Enable()
	}
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
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(body))
func (c *goHTTPClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
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
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//	log.Fatal(err)
//	}
//	fmt.Println(string(body))
func (c *goHTTPClient) Put(url string, headers http.Header, body interface{}) (*Response, error) {
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
//		body, err := io.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(body))
func (c *goHTTPClient) Delete(url string, headers http.Header, body interface{}) (*Response, error) {
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
//		body, err := util.ReadAll(response.Body)
//		if err != nil {
//			log.Fatal(err)
//	}
//		fmt.Println(string(body))
func (c *goHTTPClient) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
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
func (c *goHTTPClient) Head(url string, headers http.Header, body interface{}) (*Response, error) {
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
	c.builder.Timeout = c.builder.Timeout.Disable()
}

func (c *goHTTPClient) EnableTimeouts() {
	c.builder.Timeout = c.builder.Timeout.Enable()
}

// getClient returns the *http.client if exists or creates a new one and returns it.
func (c *goHTTPClient) getClient() *http.Client {
	c.clientOnce.Do(func() {
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
