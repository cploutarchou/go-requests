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
	response, err := _client.Do(req)
	return response, err

}
