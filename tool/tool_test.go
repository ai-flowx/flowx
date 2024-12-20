//go:build tool_test

// go test -cover -covermode=atomic -parallel 2 -tags=tool_test -v github.com/ai-flowx/flowx/tool

package tool

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initToolTest(_ context.Context) tool {
	cfg := Config{
		Type: typeFlowX,
	}

	flowx := FlowX{}

	return tool{
		cfg:   &cfg,
		flowx: &flowx,
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

func TestToolRun(t *testing.T) {
	ctx := context.Background()
	_t := initToolTest(ctx)

	_ = _t.Init(ctx)

	defer func(_t *tool, ctx context.Context) {
		_ = _t.Deinit(ctx)
	}(&_t, ctx)

	invokes := []*Invoke{
		{
			Name:        "flowx.sh",
			Description: "flowx bash",
			Path:        "../test/tool/flowx.sh",
			Func:        nil,
			Args:        nil,
		},
	}

	err := _t.Run(ctx, invokes)
	assert.Equal(t, nil, err)
}
