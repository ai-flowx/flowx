package tool

import (
	"context"
)

type ToolX interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Name(context.Context) string
	Description(context.Context) string
	Call(context.Context, ...interface{}) (string, error)
}
