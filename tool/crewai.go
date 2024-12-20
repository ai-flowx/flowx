package tool

import (
	"context"
)

type CrewAi interface {
	Name() string
	Description() string
	Call(ctx context.Context, args ...interface{}) (string, error)
}
