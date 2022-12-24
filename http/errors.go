package http

import "errors"

type ErrorContentType error

func UnsupportedContentType() ErrorContentType { return errors.New("unsupported content type") }
func NoContentType() ErrorContentType          { return errors.New("no content type") }
