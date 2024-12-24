//go:build tool_test

// go test -cover -covermode=atomic -parallel 2 -tags=tool_test -v github.com/ai-flowx/flowx/tool

package tool

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nameToolTest = "hello"
)

func initToolTest(_ context.Context) tool {
	cfg := Config{
		Provider: []Provider{
			{
				Name: nameToolTest,
			},
		},
	}

	return tool{
		cfg: &cfg,
	}
}

func TestToolInit(t *testing.T) {
	ctx := context.Background()
	_t := initToolTest(ctx)

	err := _t.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestToolDeinit(t *testing.T) {
	ctx := context.Background()
	_t := initToolTest(ctx)

	_ = _t.Init(ctx)

	err := _t.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestToolList(t *testing.T) {
	ctx := context.Background()
	_t := initToolTest(ctx)

	_ = _t.Init(ctx)

	defer func(_t *tool, ctx context.Context) {
		_ = _t.Deinit(ctx)
	}(&_t, ctx)

	_, err := _t.List(ctx)
	assert.Equal(t, nil, err)
}

func TestToolRun(t *testing.T) {
	ctx := context.Background()
	_t := initToolTest(ctx)

	_ = _t.Init(ctx)

	defer func(_t *tool, ctx context.Context) {
		_ = _t.Deinit(ctx)
	}(&_t, ctx)

	_, err := _t.Run(ctx, nameToolTest, "arg")
	assert.Equal(t, nil, err)
}
