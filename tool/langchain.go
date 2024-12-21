package tool

import (
	"context"
)

type LangChain interface {
	Name() string
	Description() string
	Call(ctx context.Context, args ...interface{}) (string, error)
}
