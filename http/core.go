package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

func (c *client) getBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return c.interfaceToJsonBytes(body)
	case "application/xml":
		data, err := xml.Marshal(body)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return c.interfaceToJsonBytes(body)
	}
}
func (c *client) do(method Method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	_client := &http.Client{}
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

	// Set all set headers to the http request

	req.Header = availableHeaders
	return _client.Do(req)
}

func (c *client) getHeaders(headers http.Header) http.Header {
	res := make(http.Header)
	// Set common headers to the request
	for header, value := range c.Headers {
		if len(value) > 0 {
			res.Set(header, value[0])
		}
	}
	// Set headers to the request
	for header, value := range headers {
		if len(value) > 0 {
			res.Set(header, value[0])
		}
	}
	return res
}

func (c *client) interfaceToJsonBytes(data interface{}) ([]byte, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
