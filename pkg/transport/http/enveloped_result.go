package http

import "github.com/mojo-lang/core/go/pkg/mojo/core"

type EnvelopedResult struct {
	*core.Error
	Data interface{} `json:"data"`
}
