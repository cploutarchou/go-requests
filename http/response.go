package http

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Response is the struct that holds the response of the HTTP request.
//
//   - It contains the HTTP status code, HTTP header, HTTP status, and response body.
//   - The response body is in []byte format.
//   - The response body can be unmarshaled into a struct using the UnmarshalJSON, UnmarshalXML, or UnmarshalYAML methods.
//   - The response body can be converted into a string using the String method.
//   - The response body can be converted into a []byte using the Body method.
//   - The HTTP status code can be retrieved using the StatusCode method.
//   - The HTTP header can be retrieved using the Header method.
//   - The HTTP status can be retrieved using the Status method.
type Response struct {
	statusCode int
	status     string
	header     http.Header
	body       []byte
}

// StatusCode returns the HTTP status code of the response.
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Header returns the HTTP header of the response.
func (r *Response) Header() http.Header {
	return r.header
}

// Body returns a response body in []byte format.
func (r *Response) Body() []byte {
	return r.body
}

// Status returns the HTTP status of the response.
func (r *Response) Status() string {
	return r.status
}

// String returns a response body in string format.
func (r *Response) String() string {
	return string(r.body)
}

// UnmarshalJSON unmarshal the response body into the given interface.
func (r *Response) UnmarshalJSON(v interface{}) error {
	return json.Unmarshal(r.body, &v)
}

// UnmarshalXML unmarshal the response body into the given interface.
func (r *Response) UnmarshalXML(v interface{}) error {
	return xml.Unmarshal(r.body, &v)
}

// UnmarshalYAML unmarshal the response body into the given interface.
func (r *Response) UnmarshalYAML(v interface{}) error {
	return yaml.Unmarshal(r.body, &v)
}
