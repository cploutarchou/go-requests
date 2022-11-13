package http

import (
	"bytes"
	"fmt"
	"strconv"
)

type HeaderType string

const (
	HeaderTypeAccept             HeaderType = "Accept"
	HeaderTypeAcceptCharset      HeaderType = "Accept-Charset"
	HeaderTypeAcceptEncoding     HeaderType = "Accept-Encoding"
	HeaderTypeAcceptLanguage     HeaderType = "Accept-Language"
	HeaderTypeAcceptRanges       HeaderType = "Accept-Ranges"
	HeaderTypeAge                HeaderType = "Age"
	HeaderTypeAllow              HeaderType = "Allow"
	HeaderTypeContentDisposition HeaderType = "Content-Disposition"
	HeaderTypeContentEncoding    HeaderType = "Content-Encoding"
	HeaderTypeContentLanguage    HeaderType = "Content-Language"
	HeaderTypeContentLength      HeaderType = "Content-Length"
	HeaderTypeContentLocation    HeaderType = "Content-Location"
	HeaderTypeContentMD5         HeaderType = "Content-MD5"
	HeaderTypeContentRange       HeaderType = "Content-Range"
	HeaderTypeContentType        HeaderType = "Content-Type"
	HeaderTypeCookie             HeaderType = "Cookie"
	HeaderTypeDate               HeaderType = "Date"
	HeaderTypeETag               HeaderType = "ETag"
	HeaderTypeExpires            HeaderType = "Expires"
)

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
}

type headerImpl struct {
	values map[string]string
}

func (h *headerImpl) Set(key, value string) Headers {
	h.values[key] = value
	return h
}

func (h *headerImpl) SetContentType(contentType string) Headers {
	h.values[string(HeaderTypeContentType)] = contentType
	return h
}

func (h *headerImpl) SetContentLength(contentLength int) Headers {
	h.values[string(HeaderTypeContentLength)] = strconv.Itoa(contentLength)
	return h
}

func (h *headerImpl) SetContentDisposition(contentDisposition string) Headers {
	h.values[string(HeaderTypeContentDisposition)] = contentDisposition
	return h
}

func (h *headerImpl) SetContentEncoding(contentEncoding string) Headers {
	h.values[string(HeaderTypeContentEncoding)] = contentEncoding
	return h
}

func (h *headerImpl) SetContentLanguage(contentLanguage string) Headers {
	h.values[string(HeaderTypeContentLanguage)] = contentLanguage
	return h
}

func (h *headerImpl) SetContentLocation(contentLocation string) Headers {
	h.values[string(HeaderTypeContentLocation)] = contentLocation
	return h
}

func (h *headerImpl) SetContentMD5(contentMD5 string) Headers {
	h.values[string(HeaderTypeContentMD5)] = contentMD5
	return h
}

func (h *headerImpl) SetContentRange(contentRange string) Headers {
	h.values[string(HeaderTypeContentRange)] = contentRange
	return h
}

func (h *headerImpl) SetCookie(cookie string) Headers {
	h.values[string(HeaderTypeCookie)] = cookie
	return h
}

func (h *headerImpl) SetDate(date string) Headers {
	h.values[string(HeaderTypeDate)] = date
	return h
}

func (h *headerImpl) SetETag(etag string) Headers {
	h.values[string(HeaderTypeETag)] = etag
	return h
}

func (h *headerImpl) SetExpires(expires string) Headers {
	h.values[string(HeaderTypeExpires)] = expires
	return h
}

func (h *headerImpl) SetAccept(accept string) Headers {
	h.values[string(HeaderTypeAccept)] = accept
	return h
}

func (h *headerImpl) SetAcceptCharset(acceptCharset string) Headers {
	h.values[string(HeaderTypeAcceptCharset)] = acceptCharset
	return h
}

func (h *headerImpl) SetAcceptEncoding(acceptEncoding string) Headers {
	h.values[string(HeaderTypeAcceptEncoding)] = acceptEncoding
	return h
}

func (h *headerImpl) SetAcceptLanguage(acceptLanguage string) Headers {
	h.values[string(HeaderTypeAcceptLanguage)] = acceptLanguage
	return h
}

func (h *headerImpl) SetAcceptRanges(acceptRanges string) Headers {
	h.values[string(HeaderTypeAcceptRanges)] = acceptRanges
	return h
}

func (h *headerImpl) SetAge(age string) Headers {
	h.values[string(HeaderTypeAge)] = age
	return h
}

func (h *headerImpl) SetAllow(allow string) Headers {
	h.values[string(HeaderTypeAllow)] = allow
	return h
}

func (h *headerImpl) SetCustom(key, value string) Headers {
	h.values[key] = value
	return h
}
func (h *headerImpl) Get(key string) string {
	return h.values[key]
}

func (h *headerImpl) Del(key string) Headers {
	delete(h.values, key)
	return h
}

func (h *headerImpl) Clone() Headers {
	clone := make(map[string]string)
	for k, v := range h.values {
		clone[k] = v
	}
	return &headerImpl{values: clone}
}

func (h *headerImpl) IsEmpty() bool {
	return len(h.values) == 0
}

func (h *headerImpl) IsSet() bool {
	return !h.IsEmpty()
}

func (h *headerImpl) String() string {
	var buffer bytes.Buffer
	for k, v := range h.values {
		buffer.WriteString(fmt.Sprintf("%s: %s", k, v))
	}
	return buffer.String()
}

func (h *headerImpl) Values() map[string]string {
	return h.values
}

func (h *headerImpl) Keys() []string {
	var keys []string
	for k := range h.values {
		keys = append(keys, k)
	}
	return keys
}

func (h *headerImpl) Len() int {
	return len(h.values)
}

func (h *headerImpl) GetAll() map[string][]string {
	all := make(map[string][]string)
	for k, v := range h.values {
		all[k] = []string{v}
	}
	return all
}
func NewHeaders() Headers {
	return &headerImpl{values: make(map[string]string)}
}
