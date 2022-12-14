package go_requests

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// do is the main method to make the request
// It returns the response and an error if something goes wrong
// It is private because it is only used by the public methods
//
//	func (c *goHTTPClient) do(method Method, url string, Headers http.Header, body interface{}) (*http.Response, error) {
//		var err error
//		var req *http.Request
//		availableHeaders := c.getHeaders(Headers)
//		requestBody, err := c.getBody(availableHeaders.Get("Content-Type"), body)
//		if err != nil {
//			return nil, err
//		}
//		if body != nil {
//			reader := bytes.NewReader(requestBody)
//			req, err = http.NewRequest(string(method), url, reader)
//		} else {
//			req, err = http.NewRequest(string(method), url, nil)
//		}
//		if err != nil {
//			return nil, errors.New("unable to create request")
//		}
//		// Set all set Headers to the http request
//		req.Header = availableHeaders
//		// Return the response
//		return c.client.Do(req)
//	}
func (c *goHTTPClient) do(method Method, url string, headers http.Header, body []byte) (*Response, error) {
	var req *http.Request
	var err error
	availableHeaders := c.getHeaders(headers)
	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(string(method), url, reader)
	} else {
		req, err = http.NewRequest(string(method), url, nil)
	}
	if c.QueryParams().Len() > 0 {
		q := req.URL.Query()
		for key, value := range c.QueryParams().Values() {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
		c.QueryParams().Reset()
	}

	if err != nil {
		return nil, errors.New("unable to create request")
	}
	// Set all set Headers to the http request
	req.Header = availableHeaders
	// Return the response
	c.client = c.getClient()
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("unable to read response body. Error: " + err.Error())
	}
	finalResponse := Response{
		statusCode:  response.StatusCode,
		header:      response.Header,
		body:        responseBody,
		status:      response.Status,
		contentType: response.Header.Get("Content-Type"),
	}
	return &finalResponse, nil
}

// getHeaders returns the Headers that are set by the user and the default Headers that are set by the client
func (c *goHTTPClient) getHeaders(headers http.Header) http.Header {
	res := make(http.Header)
	// Set common Headers to the request
	for header, value := range c.builder.Headers().GetAll() {
		if len(value) > 0 {
			res.Set(header, value[0])
		}
	}
	// Set Headers to the request
	for header, value := range headers {
		if len(value) > 0 {
			res.Set(header, value[0])
		}
	}
	return res
}
