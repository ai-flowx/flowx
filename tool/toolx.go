package tool

import (
	"context"
)

type ToolX interface {
	Name() string
	Description() string
	Call(ctx context.Context, args ...interface{}) (string, error)
}
