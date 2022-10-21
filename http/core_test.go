package http

import (
	"testing"
)

func TestGetHeaders(t *testing.T) {
	// Initialization
	client := client{}
	commonHeaders := client.MakeHeaders()
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "the-best-http-client")
	client.Headers = commonHeaders
	// Execution
	requestHeaders := client.MakeHeaders()
	requestHeaders.Set("X-Request-ID", "ABC123")
	headers := client.getHeaders(requestHeaders)
	//Validation
	if len(headers) != 3 {
		t.Errorf("expected 3 headers. Provided only %d", len(headers))
	}
}
