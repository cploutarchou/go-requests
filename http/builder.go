package http

import (
	"time"
)

type builderImpl struct {
	header  Headers
	Timeout Timeout
}

func (c builderImpl) SetMaxIdleConnections(maxConnections int) Timeout {
	c.Timeout.SetMaxIdleConnections(maxConnections)
	return c.Timeout
}

func (c builderImpl) Headers() Headers {
	return c.header
}

func (c builderImpl) GetMaxIdleConnections() int {
	return c.Timeout.GetMaxIdleConnections()
}

func (c builderImpl) SetRequestTimeout(timeout time.Duration) Timeout {
	c.Timeout.SetRequestTimeout(timeout)
	return c.Timeout
}

func (c builderImpl) SetResponseTimeout(timeout time.Duration) Timeout {
	c.Timeout.SetResponseTimeout(timeout)
	return c.Timeout
}

type Builder interface {
	SetRequestTimeout(timeout time.Duration) Timeout
	SetResponseTimeout(timeout time.Duration) Timeout
	SetMaxIdleConnections(maxConnections int) Timeout
	Headers() Headers
	Build() Client
}

func (c builderImpl) Build() Client {
	return &goHTTPClient{
		Timeout: c.Timeout,
		Headers: c.header,
		state:   make(chan string, 100),
	}
}

func NewBuilder() Builder {
	builder := &builderImpl{
		Timeout: newTimeouts(),
		header:  NewHeaders(),
	}
	return builder
}
