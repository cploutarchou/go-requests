package http

import (
	"bytes"
	"errors"
	"net/http"
)

func (c *client) do(method Method, url string, headers http.Header, body []byte) (*http.Response, error) {
	_client := &http.Client{}
	var err error
	var req *http.Request
	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(string(method), url, reader)
	} else {
		req, err = http.NewRequest(string(method), url, nil)
	}
	if err != nil {
		return nil, errors.New("unable to create request")
	}

	// Set all set headers to the http request
	availableHeaders := c.getHeaders(headers)
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
