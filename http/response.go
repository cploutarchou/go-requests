package http

import (
	"net/http"
)

type Response struct {
	statusCode int
	status     string
	header     http.Header
	body       []byte
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Header() http.Header {
	return r.header
}

func (r *Response) Bytes() []byte {
	return r.body
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) String() string {
	return string(r.body)
}
