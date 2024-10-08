// Package registry uses the go-micro registry for selection
package registry

import (
	"go-micro.dev/v5/selector"
	"go-micro.dev/v5/cmd"
)

func init() {
	cmd.DefaultSelectors["registry"] = NewSelector
}

// NewSelector returns a new registry selector.
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
