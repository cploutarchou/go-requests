package requests

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

type HeaderType string

const (
	//HeaderTypeAccept is the Accept header
	HeaderTypeAccept HeaderType = "Accept"
	//HeaderTypeAcceptCharset is the Accept-Charset header
	HeaderTypeAcceptCharset HeaderType = "Accept-Charset"
	//HeaderTypeAcceptEncoding is the Accept-Encoding header
	HeaderTypeAcceptEncoding HeaderType = "Accept-Encoding"
	//HeaderTypeAcceptLanguage is the Accept-Language header
	HeaderTypeAcceptLanguage HeaderType = "Accept-Language"
	//HeaderTypeAcceptRanges is the Accept-Ranges header
	HeaderTypeAcceptRanges HeaderType = "Accept-Ranges"
	//HeaderTypeAge is the Age header
	HeaderTypeAge HeaderType = "Age"
	//HeaderTypeAllow is the Allow header
	HeaderTypeAllow HeaderType = "Allow"
	//HeaderTypeContentDisposition is the Content-Disposition header
	HeaderTypeContentDisposition HeaderType = "Content-Disposition"
	//HeaderTypeContentEncoding is the Content-Encoding header
	HeaderTypeContentEncoding HeaderType = "Content-Encoding"
	//HeaderTypeContentLanguage is the Content-Language header
	HeaderTypeContentLanguage HeaderType = "Content-Language"
	//HeaderTypeContentLength is the Content-Length header
	HeaderTypeContentLength HeaderType = "Content-Length"
	//HeaderTypeContentLocation is the Content-Location header
	HeaderTypeContentLocation HeaderType = "Content-Location"
	//HeaderTypeContentMD5 is the Content-MD5 header
	HeaderTypeContentMD5 HeaderType = "Content-MD5"
	//HeaderTypeContentRange is the Content-Range header
	HeaderTypeContentRange HeaderType = "Content-Range"
	//HeaderTypeContentType is the Content-Type header
	HeaderTypeContentType HeaderType = "Content-Type"
	//HeaderTypeCookie is the Cookie header
	HeaderTypeCookie HeaderType = "Cookie"
	//HeaderTypeDate is the Date header
	HeaderTypeDate HeaderType = "Date"
	//HeaderTypeETag is the ETag header
	HeaderTypeETag HeaderType = "ETag"
	//HeaderTypeExpires is the Expires header
	HeaderTypeExpires HeaderType = "Expires"
)

// Headers is the interface for the http headers object of the http package of the standard library of Go (golang)
type Headers interface {
	// Set sets a header to the header
	Set(key, value string) Headers
	// SetContentType sets the content type to the header
	SetContentType(contentType string) Headers
	// SetContentLength sets the content length to the header
	SetContentLength(contentLength int) Headers
	// SetContentDisposition sets the content disposition to the header
	SetContentDisposition(contentDisposition string) Headers
	// SetContentEncoding sets the content encoding to the header
	SetContentEncoding(contentEncoding string) Headers
	// SetContentLanguage sets the content language to the header
	SetContentLanguage(contentLanguage string) Headers
	// SetContentLocation sets the content location to the header
	SetContentLocation(contentLocation string) Headers
	// SetContentMD5 sets the content md5 to the header
	SetContentMD5(contentMD5 string) Headers
	// SetContentRange sets the content range to the header
	SetContentRange(contentRange string) Headers
	// SetCookie sets the cookie to the header
	SetCookie(cookie string) Headers
	// SetDate sets the date to the header
	SetDate(date string) Headers
	// SetETag sets the etag to the header
	SetETag(etag string) Headers
	// SetExpires sets the expires to the header
	SetExpires(expires string) Headers
	// SetAccept sets to accept to the header
	SetAccept(accept string) Headers
	// SetAcceptCharset sets to accept charset to the header
	SetAcceptCharset(acceptCharset string) Headers
	// SetAcceptEncoding sets to accept encoding to the header
	SetAcceptEncoding(acceptEncoding string) Headers
	// SetAcceptLanguage sets to accept language to the header
	SetAcceptLanguage(acceptLanguage string) Headers
	// SetAcceptRanges sets to accept ranges to the header
	SetAcceptRanges(acceptRanges string) Headers
	// SetAge sets the age to the header
	SetAge(age string) Headers
	// SetAllow sets to allow to the header
	SetAllow(allow string) Headers
	// SetCustom sets a custom header to the header
	SetCustom(key, value string) Headers
	// Get returns the value of the header
	Get(key string) string
	// Del deletes a header from the header
	Del(key string) Headers
	// Clone returns a clone of the header
	Clone() Headers
	// IsEmpty returns true if the header is empty
	IsEmpty() bool
	// IsSet returns true if the header is set
	IsSet() bool
	// String returns the string representation of the header
	String() string
	// Values returns the values of the header
	Values() map[string]string
	// Keys returns the keys of the header
	Keys() []string
	// Len returns the length of the header
	Len() int
	// GetAll returns all Headers with key and value
	GetAll() map[string][]string
	//GetAllHttpHeaders returns all http headers as http.Header object
	GetAllHttpHeaders() http.Header
}

// headerImpl is the implementation of the Headers interface
type headerImpl struct {
	values map[string]string
}

// Set sets a header to the header object
// key is the key of the header
// value is the value of the header
// returns the header object
//
// Example:
//
//	header := NewHeaders()
//	header.Set("Content-Type", "application/json")
//	header.Set("Content-Length", "100")
func (h *headerImpl) Set(key, value string) Headers {
	h.values[key] = value
	return h
}

// SetContentType sets the content type to the header object
func (h *headerImpl) SetContentType(contentType string) Headers {
	h.values[string(HeaderTypeContentType)] = contentType
	return h
}

// SetContentLength sets the content length to the header object
func (h *headerImpl) SetContentLength(contentLength int) Headers {
	h.values[string(HeaderTypeContentLength)] = strconv.Itoa(contentLength)
	return h
}

// SetContentDisposition sets the content disposition to the header object
func (h *headerImpl) SetContentDisposition(contentDisposition string) Headers {
	h.values[string(HeaderTypeContentDisposition)] = contentDisposition
	return h
}

// SetContentEncoding sets the content encoding to the header object
func (h *headerImpl) SetContentEncoding(contentEncoding string) Headers {
	h.values[string(HeaderTypeContentEncoding)] = contentEncoding
	return h
}

// SetContentLanguage sets the content language to the header object
func (h *headerImpl) SetContentLanguage(contentLanguage string) Headers {
	h.values[string(HeaderTypeContentLanguage)] = contentLanguage
	return h
}

// SetContentLocation sets the content location to the header object
func (h *headerImpl) SetContentLocation(contentLocation string) Headers {
	h.values[string(HeaderTypeContentLocation)] = contentLocation
	return h
}

// SetContentMD5 sets the content md5 to the header object
func (h *headerImpl) SetContentMD5(contentMD5 string) Headers {
	h.values[string(HeaderTypeContentMD5)] = contentMD5
	return h
}

func (h *headerImpl) SetContentRange(contentRange string) Headers {
	h.values[string(HeaderTypeContentRange)] = contentRange
	return h
}

// SetCookie sets the cookie to the header object
func (h *headerImpl) SetCookie(cookie string) Headers {
	h.values[string(HeaderTypeCookie)] = cookie
	return h
}

// SetDate sets the date to the header object
func (h *headerImpl) SetDate(date string) Headers {
	h.values[string(HeaderTypeDate)] = date
	return h
}

// SetETag sets the etag to the header object
func (h *headerImpl) SetETag(etag string) Headers {
	h.values[string(HeaderTypeETag)] = etag
	return h
}

// SetExpires sets the expires to the header object
func (h *headerImpl) SetExpires(expires string) Headers {
	h.values[string(HeaderTypeExpires)] = expires
	return h
}

// SetAccept sets to accept to the header object
func (h *headerImpl) SetAccept(accept string) Headers {
	h.values[string(HeaderTypeAccept)] = accept
	return h
}

// SetAcceptCharset sets to accept charset to the header object
func (h *headerImpl) SetAcceptCharset(acceptCharset string) Headers {
	h.values[string(HeaderTypeAcceptCharset)] = acceptCharset
	return h
}

// SetAcceptEncoding sets to accept encoding to the header object
func (h *headerImpl) SetAcceptEncoding(acceptEncoding string) Headers {
	h.values[string(HeaderTypeAcceptEncoding)] = acceptEncoding
	return h
}

// SetAcceptLanguage sets to accept language to the header object
func (h *headerImpl) SetAcceptLanguage(acceptLanguage string) Headers {
	h.values[string(HeaderTypeAcceptLanguage)] = acceptLanguage
	return h
}

// SetAcceptRanges sets to accept ranges to the header object
func (h *headerImpl) SetAcceptRanges(acceptRanges string) Headers {
	h.values[string(HeaderTypeAcceptRanges)] = acceptRanges
	return h
}

// SetAge sets the age to the header object
func (h *headerImpl) SetAge(age string) Headers {
	h.values[string(HeaderTypeAge)] = age
	return h
}

// SetAllow sets to allow to the header object
func (h *headerImpl) SetAllow(allow string) Headers {
	h.values[string(HeaderTypeAllow)] = allow
	return h
}

// SetCustom sets a custom header to the header object
func (h *headerImpl) SetCustom(key, value string) Headers {
	h.values[key] = value
	return h
}

// Get returns the value of the header
func (h *headerImpl) Get(key string) string {
	return h.values[key]
}

// GetAllHttpHeaders returns all http headers as http.Header object
func (h *headerImpl) GetAllHttpHeaders() http.Header {
	headers := http.Header{}
	for k, v := range h.values {
		headers.Set(k, v)
	}
	return headers
}

// Del deletes a header from the header object
func (h *headerImpl) Del(key string) Headers {
	delete(h.values, key)
	return h
}

// Clone clones the header object
func (h *headerImpl) Clone() Headers {
	clone := make(map[string]string)
	for k, v := range h.values {
		clone[k] = v
	}
	return &headerImpl{values: clone}
}

// IsEmpty checks if the header object is empty
// returns true if the header object is empty, otherwise false
func (h *headerImpl) IsEmpty() bool {
	return len(h.values) == 0
}

// IsSet checks if the header object has a header with the given key
func (h *headerImpl) IsSet() bool {
	return !h.IsEmpty()
}

// GetAllHttpHeaders returns all http headers as http.Header object
func (h *headerImpl) String() string {
	var buffer bytes.Buffer
	for k, v := range h.values {
		buffer.WriteString(fmt.Sprintf("%s: %s", k, v))
	}
	return buffer.String()
}

// Values returns the values of the header object
func (h *headerImpl) Values() map[string]string {
	return h.values
}

// Keys returns the keys of the header object
func (h *headerImpl) Keys() []string {
	var keys []string
	for k := range h.values {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the length of the header object
func (h *headerImpl) Len() int {
	return len(h.values)
}

// GetAll returns all headers as map of string and string array object (key, value) pairs of the header object as map
func (h *headerImpl) GetAll() map[string][]string {
	all := make(map[string][]string)
	for k, v := range h.values {
		all[k] = []string{v}
	}
	return all
}

// NewHeaders returns a new header object
func NewHeaders() Headers {
	return &headerImpl{values: make(map[string]string)}
}

// getHeader returns the header object of the  ...http.Header object ad http.Header object.
// If the header object is nil, it returns a new header object
// If the header object is not nil, it returns the header object
func getHeader(headers ...http.Header) http.Header {
	if len(headers) == 0 {
		return http.Header{}
	}
	return headers[0]
}
