package http

import "net/http"

type Headers interface {
	SetHeaders(http.Header)
	GetHeaders() http.Header
}

type headersImpl struct {
	header http.Header
}

func newHeadersImpl(header http.Header) *headersImpl {
	return &headersImpl{header: make(http.Header)}
}

func (c headersImpl) SetHeaders(h http.Header) {
	c.header = h
}
func (c headersImpl) GetHeaders() http.Header {
	return c.header
}
