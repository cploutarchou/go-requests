package http

import (
	"time"
)

type builderImpl struct {
	header          Headers
	timeoutSettings Timeout
}

func (c builderImpl) SetMaxIdleConnections(maxConnections int) Timeout {
	c.timeoutSettings.SetMaxIdleConnections(maxConnections)
	return c.timeoutSettings
}

func (c builderImpl) Headers() Headers {
	return c.header
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
	SetMaxIdleConnections(maxConnections int) Timeout
	Headers() Headers
	Build() Client
}

func (c builderImpl) Build() Client {
	return &goHTTPClient{
		timeout: c.timeoutSettings,
		headers: c.header,
	}
}

func NewBuilder() Builder {
	builder := &builderImpl{
		timeoutSettings: newTimeouts(),
		header:          NewHeaders(),
	}
	return builder
}
