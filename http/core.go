package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strings"
)

// getBody returns the body of the request.
func (c *goHTTPClient) getBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return c.interfaceToJSONBytes(body)
	case "application/xml":
		return c.interfaceToXMLBytes(body)
	default:
		return c.interfaceToJSONBytes(body)
	}
}

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
func (c *goHTTPClient) do(method Method, url string, headers http.Header, body interface{}) (*Response, error) {
	var err error
	var req *http.Request
	availableHeaders := c.getHeaders(headers)
	requestBody, err := c.getBody(availableHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		reader := bytes.NewReader(requestBody)
		req, err = http.NewRequest(string(method), url, reader)
	} else {
		req, err = http.NewRequest(string(method), url, nil)
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
		statusCode: response.StatusCode,
		header:     response.Header,
		body:       responseBody,
		status:     response.Status,
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

// interfaceToJSONBytes converts an interface to a JSON byte array and returns it and an error if something goes wrong
func (c *goHTTPClient) interfaceToJSONBytes(data interface{}) ([]byte, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// interfaceToXMLBytes converts an interface to an XML byte array and returns it and an error if something goes wrong
func (c *goHTTPClient) interfaceToXMLBytes(data interface{}) ([]byte, error) {
	res, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
