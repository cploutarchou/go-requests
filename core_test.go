package go_requests

import (
	"net/http"
	"testing"
)

func TestGetHeaders(t *testing.T) {
	// Initialization
	builder := NewBuilder()
	builder.Headers().
		SetContentType("application/json").
		SetAcceptEncoding("gzip").
		SetAcceptCharset("utf-8").
		SetCustom("User-Agent", "the-best-http-goHttpClient").
		SetCustom("X-Request-ID", "ABC123")
	headers := builder.Headers()
	// Validation
	if headers.Len() != 5 {
		t.Errorf("expected 5 Headers. Provided only %d", headers.Len())
	}
	if headers.Get("X-Request-ID") != "ABC123" {
		t.Errorf("expected 'X-Request-ID' header. Provided %s", headers.Get("X-Request-ID"))
	}
	if headers.Get("Content-Type") != "application/json" {
		t.Errorf("expected 'Content-Type' header. Provided %s", headers.Get("Content-Type"))
	}
	if headers.Get("User-Agent") != "the-best-http-goHttpClient" {
		t.Errorf("expected 'User-Agent' header. Provided %s", headers.Get("User-Agent"))
	}
	if headers.Get("Accept-Encoding") != "gzip" {
		t.Errorf("expected 'Accept-Encoding' header. Provided %s", headers.Get("Accept-Encoding"))
	}

	if headers.Get("Accept-Charset") != "utf-8" {
		t.Errorf("expected 'Accept-Charset' header. Provided %s", headers.Get("Accept-Charset"))
	}
}

func Test_goHTTPClient_getHeaders(t *testing.T) {
	type fields struct {
		client  *http.Client
		builder builderImpl
	}
	builder := builderImpl{
		header:  NewHeaders(),
		Timeout: newTimeouts(),
	}
	builderNoHeaders := builderImpl{
		header:  nil,
		Timeout: newTimeouts(),
	}
	builder.header = builder.Headers().SetContentType("application/json")
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		{name: "empty", fields: fields{builder: builderNoHeaders}, want: http.Header{}},
		{name: "set", fields: fields{builder: builder}, want: http.Header{"Content-Type": []string{"application/json"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goHTTPClient{
				client:  tt.fields.client,
				builder: &tt.fields.builder,
			}
			if tt.name == "set" {
				//check if the headers content type is set to application/json
				if c.builder.header.Get("Content-Type") != "application/json" {
					t.Errorf("invalid content type. Provided %s", c.builder.header.Get("Content-Type"))
				}
				if c.builder.header.Len() != 1 {
					t.Errorf("invalid header length. Provided %d", c.builder.header.Len())
				}
			}
			if tt.name == "empty" {
				//check if the headers content type is set to application/json
				if c.builder.header != nil {
					t.Errorf("invalid content type. Provided %s", c.builder.header.Get("Content-Type"))
				}
			}
		})
	}
}
