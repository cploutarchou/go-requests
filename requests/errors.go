package requests

import "errors"

// ErrorContentType is the error type for content type errors
type ErrorContentType error

// UnsupportedContentType is the error type for unsupported content type errors. It is returned when the content type is not supported by the library.
func UnsupportedContentType() ErrorContentType { return errors.New("unsupported content type") }

// NoContentType is the error type for no content type errors. It is returned when the content type is not set. This is the default error type.
func NoContentType() ErrorContentType { return errors.New("no content type") }
