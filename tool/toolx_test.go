//go:build tool_toolx_test

// go test -cover -covermode=atomic -parallel 2 -tags=tool_toolx_test -v github.com/ai-flowx/flowx/tool

package tool

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nameToolXDecoratorTest = "decorator"
	nameToolXHelloTest     = "hello"
)

func initToolXTest(_ context.Context) tool {
	cfg := Config{
		Provider: []Provider{
			{
				Name: nameToolXDecoratorTest,
			},
			{
				Name: nameToolXHelloTest,
			},
		},
	}

	return tool{
		cfg: &cfg,
	}
}

func TestToolXInit(t *testing.T) {
	ctx := context.Background()
	_t := initToolXTest(ctx)

	err := _t.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestToolXDeinit(t *testing.T) {
	ctx := context.Background()
	_t := initToolXTest(ctx)

	_ = _t.Init(ctx)

	err := _t.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestToolXRun(t *testing.T) {
	ctx := context.Background()
	_t := initToolXTest(ctx)

	_ = _t.Init(ctx)

	defer func(_t *tool, ctx context.Context) {
		_ = _t.Deinit(ctx)
	}(&_t, ctx)

	buf, err := _t.Run(ctx, nameToolXDecoratorTest, "arg")
	assert.Equal(t, nil, err)

	fmt.Println(string(buf))

	buf, err = _t.Run(ctx, nameToolXHelloTest, "arg")
	assert.Equal(t, nil, err)

	fmt.Println(string(buf))
}
