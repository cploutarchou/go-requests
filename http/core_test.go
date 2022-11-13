package http

import (
	"reflect"
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

func TestGetBody(t *testing.T) {
	// Initialization
	_client := goHTTPClient{}
	// Validation
	t.Run("noBody", func(t *testing.T) {
		// Execution
		requestBody, err := _client.getBody("", nil)
		// Validation
		if err != nil {
			t.Errorf("expected no error when passed a nil body. Provided %s", err.Error())
		}
		if requestBody != nil {
			t.Errorf("expected a nil body. Provided %s", requestBody)
		}
	})
	t.Run("jsonBody", func(t *testing.T) {
		// Execution
		body := []string{"Hello", "World"}
		requestBody, err := _client.getBody("application/json", body)
		// Validation
		if err != nil {
			t.Errorf("expected no error when passed a json body. Provided %s", err.Error())
		}

		if string(requestBody) != `["Hello","World"]` {
			t.Errorf("invalid request body. Provided %s", string(requestBody))
		}
	})
	t.Run("xmlBody", func(t *testing.T) {
		// Execution
		body := []string{"Hello", "World"}
		requestBody, err := _client.getBody("application/xml", body)
		// Validation
		if err != nil {
			t.Errorf("expected no error when passed a xml body. Provided %s", err.Error())
		}

		if string(requestBody) != `<string>Hello</string><string>World</string>` {
			t.Errorf("invalid request body. Provided %s", string(requestBody))
		}
	})
	t.Run("defaultJson", func(t *testing.T) {
		// Execution
		body := []string{"Hello", "World"}
		requestBody, err := _client.getBody("", body)
		// Validation
		if err != nil {
			t.Errorf("expected no error when passed a json body. Provided %s", err.Error())
		}
		if string(requestBody) != `["Hello","World"]` {
			t.Errorf("invalid request body. Provided %s", string(requestBody))
		}
	})
}

func Test_goHTTPClient_interfaceToJSONBytes(t *testing.T) {
	type args struct {
		data interface{}
	}
	var tests []struct {
		name    string
		c       *goHTTPClient
		args    args
		want    []byte
		wantErr bool
	}
	tests = append(tests, struct {
		name    string
		c       *goHTTPClient
		args    args
		want    []byte
		wantErr bool
	}{
		name: "string",
		c:    &goHTTPClient{},
		args: args{
			data: "Hello World",
		},
		want:    []byte(`"Hello World"`),
		wantErr: false,
	})
	tests = append(tests, struct {
		name string
		c    *goHTTPClient
		args args
		want []byte

		wantErr bool
	}{
		name: "int",
		c:    &goHTTPClient{},
		args: args{
			data: 123,
		},
		want:    []byte(`123`),
		wantErr: false,
	})
	tests = append(tests, struct {
		name string
		c    *goHTTPClient
		args args
		want []byte

		wantErr bool
	}{
		name: "float",
		c:    &goHTTPClient{},
		args: args{
			data: 123.45,
		},
		want:    []byte(`123.45`),
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.interfaceToJSONBytes(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("goHTTPClient.interfaceToJSONBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goHTTPClient.interfaceToJSONBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
