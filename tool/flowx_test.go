//go:build tool_flowx_test

// go test -cover -covermode=atomic -parallel 2 -tags=tool_flowx_test -v github.com/ai-flowx/flowx/tool

package tool

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initFlowXTest(_ context.Context) FlowX {
	return FlowX{}
}

func TestFlowXInit(t *testing.T) {
	ctx := context.Background()
	f := initFlowXTest(ctx)

	err := f.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestFlowXDeinit(t *testing.T) {
	ctx := context.Background()
	f := initFlowXTest(ctx)

	_ = f.Init(ctx)

	err := f.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestFlowXRun(t *testing.T) {
	ctx := context.Background()
	f := initFlowXTest(ctx)

	_ = f.Init(ctx)

	defer func(f *FlowX, ctx context.Context) {
		_ = f.Deinit(ctx)
	}(&f, ctx)

	invokes := []*Invoke{
		{
			Name:        "crewai.py",
			Description: "crewai python",
			Path:        "../test/tool/crewai.py",
			Func:        nil,
			Args:        nil,
		},
		{
			Name:        "flowx.go",
			Description: "flowx golang",
			Path:        "../test/tool/flowx.go",
			Func:        nil,
			Args:        nil,
		},
		{
			Name:        "flowx.sh",
			Description: "flowx bash",
			Path:        "../test/tool/flowx.sh",
			Func:        nil,
			Args:        nil,
		},
		{
			Name:        "langchain.py",
			Description: "langchain python",
			Path:        "../test/tool/langchain.py",
			Func:        nil,
			Args:        nil,
		},
	}

	err := f.Run(ctx, invokes)
	assert.Equal(t, nil, err)
}
