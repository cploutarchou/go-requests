package go_requests

import (
	"encoding"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

// ContentType of the response to determine the unmarshal method.
type ContentType string

const (
	// jsonContentType is the content type for json.
	jsonContentType ContentType = "application/json"
	// xmlContentType is the content type for xml.
	xmlContentType ContentType = "application/xml"
	// yamlContentType is the content type for yaml.
	yamlContentType ContentType = "application/yaml"
	// textContentType is the content type for text.
	textContentType ContentType = "text/plain"
	// noneContentType is the content type for none.
	noneContentType ContentType = ""
)

// Response is the struct that holds the response of the HTTP request.
//
//   - It contains the HTTP status code, HTTP header, HTTP status, and response body.
//   - The response body is in []byte format.
//   - The response body can be unmarshaled into a struct using the unmarshalJSON, unmarshalXML, or UnmarshalYAML methods.
//   - The response body can be converted into a string using the String method.
//   - The response body can be converted into a []byte using the Body method.
//   - The HTTP status code can be retrieved using the StatusCode method.
//   - The HTTP header can be retrieved using the Header method.
//   - The HTTP status can be retrieved using the Status method.
type Response struct {
	statusCode  int
	status      string
	header      http.Header
	body        []byte
	contentType string
}

// StatusCode returns the HTTP status code of the response.
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Header returns the HTTP header of the response.
func (r *Response) Header() http.Header {
	return r.header
}

// Bytes returns the response body in []byte format.
func (r *Response) Bytes() []byte {
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

// getContentType returns the content-type of the response. It returns an empty string if the content-type is not set.
func (r *Response) getContentType() ContentType {
	if strings.Contains(r.contentType, "application/json") {
		return jsonContentType
	}
	if strings.Contains(r.contentType, "application/xml") {
		return xmlContentType
	}
	if strings.Contains(r.contentType, "application/yaml") {
		return yamlContentType
	}
	if strings.Contains(r.contentType, "text/plain") {
		return textContentType
	}
	return noneContentType
}

// unmarshalJSON unmarshal the response body into the given interface.
func (r *Response) unmarshalJSON(v interface{}) error {
	return json.Unmarshal(r.body, &v)
}

// unmarshalXML unmarshal the response body into the given interface.
func (r *Response) unmarshalXML(v interface{}) error {
	return xml.Unmarshal(r.body, &v)
}

// UnmarshalYAML unmarshal the response body into the given interface.
func (r *Response) UnmarshalYAML(v interface{}) error {
	return yaml.Unmarshal(r.body, &v)
}

// unmarshalText unmarshal the response body into the given interface.
func (r *Response) unmarshalText(v interface{}) error {
	if p, ok := v.(encoding.TextUnmarshaler); ok {
		return p.UnmarshalText(r.body)
	}
	return UnsupportedContentType()
}

// Unmarshal the response body into the given interface.
//   - It uses the content-type of the response to determine the unmarshal method.
//   - It supports json, xml, and yaml.
//   - It returns an error if the content-type is not supported.
//   - It returns an error if the unmarshal method fails.
//   - It returns an error if the given interface is not a pointer.
func (r *Response) Unmarshal(v interface{}) ErrorContentType {
	switch r.getContentType() {
	case jsonContentType:
		return r.unmarshalJSON(v)
	case xmlContentType:
		return r.unmarshalXML(v)
	case yamlContentType:
		return r.UnmarshalYAML(v)
	case textContentType:
		return r.unmarshalText(v)
	case noneContentType:
		return NoContentType()
	default:
		return UnsupportedContentType()
	}

}

func (r *Response) ContentType() ContentType {
	return r.getContentType()
}
