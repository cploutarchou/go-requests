package gohttp

import (
	"github.com/cploutarchou/go-http/http"
)

func exampleUsage() {
	client := http.NewClient()
	client.Get()
}
