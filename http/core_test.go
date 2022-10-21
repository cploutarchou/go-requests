package http

import (
	"testing"
)

func TestGetHeaders(t *testing.T) {
	// Initialization
	_client := client{}
	commonHeaders := _client.MakeHeaders()
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "the-best-http-client")
	_client.Headers = commonHeaders
	// Execution
	requestHeaders := _client.MakeHeaders()
	requestHeaders.Set("X-Request-ID", "ABC123")
	headers := _client.getHeaders(requestHeaders)
	//Validation
	if len(headers) != 3 {
		t.Errorf("expected 3 headers. Provided only %d", len(headers))
	}
	if headers.Get("X-Request-ID") != "ABC123" {
		t.Errorf("expected 'X-Request-ID' header. Provided %s", headers.Get("X-Request-ID"))
	}
	if headers.Get("Content-Type") != "application/json" {
		t.Errorf("expected 'Content-Type' header. Provided %s", headers.Get("Content-Type"))
	}

	if headers.Get("User-Agent") != "the-best-http-client" {
		t.Errorf("expected 'User-Agent' header. Provided %s", headers.Get("User-Agent"))
	}
}

func TestGetBody(t *testing.T) {
	// Initialization
	_client := client{}
	t.Run("noBody", func(t *testing.T) {
		// Execution
		requestBody, err := _client.getBody("", nil)
		//Validation
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
		//Validation
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
		//Validation
		if err != nil {
			t.Errorf("expected no error when passed a json body. Provided %s", err.Error())
		}

		if string(requestBody) != `["Hello","World"]` {
			t.Errorf("invalid request body. Provided %s", string(requestBody))
		}
	})

}
