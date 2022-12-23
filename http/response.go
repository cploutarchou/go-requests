package http

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Response struct {
	statusCode int
	status     string
	header     http.Header
	body       []byte
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Header() http.Header {
	return r.header
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) UnmarshalJSON(v interface{}) error {
	return json.Unmarshal(r.body, &v)
}

func (r *Response) UnmarshalXML(v interface{}) error {
	return xml.Unmarshal(r.body, &v)
}

func (r *Response) UnmarshalYAML(v interface{}) error {
	return yaml.Unmarshal(r.body, &v)
}
