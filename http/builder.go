package http

import (
	"net/http"
	"time"
)

type builderImpl struct {
	header          Headers
	timeoutSettings Timeout
}

func (c builderImpl) SetHeaders(header http.Header) Headers {
	return c.header.SetHeaders(header)
}

func (c builderImpl) GetMaxIdleConnections() int {
	return c.timeoutSettings.GetMaxIdleConnections()
}

func (c builderImpl) SetRequestTimeout(timeout time.Duration) Timeout {
	c.timeoutSettings.SetRequestTimeout(timeout)
	return c.timeoutSettings
}

func (c builderImpl) SetResponseTimeout(timeout time.Duration) Timeout {
	c.timeoutSettings.SetResponseTimeout(timeout)
	return c.timeoutSettings
}

type Builder interface {
	SetRequestTimeout(timeout time.Duration) Timeout
	SetResponseTimeout(timeout time.Duration) Timeout
	SetHeaders(http.Header) Headers
	Build() Client
}

func (c timeoutImpl) SetRequestTimeout(timeout time.Duration) Timeout {
	c.RequestTimeout = timeout
	return c
}

func (c timeoutImpl) SetResponseTimeout(timeout time.Duration) Timeout {
	c.ResponseTimeout = timeout
	return c
}

func (c timeoutImpl) SetMaxIdleConnections(maxConnections int) Timeout {
	c.MaxIdleConnections = maxConnections
	return c
}

func (c builderImpl) Build() Client {
	return &goHTTPClient{
		timeout: c.timeoutSettings,
		header:  c.header,
	}
}

func NewBuilder() Builder {
	builder := &builderImpl{
		timeoutSettings: newTimeoutImpl(),
		header:          newHeadersImpl(),
	}
	return builder
}
