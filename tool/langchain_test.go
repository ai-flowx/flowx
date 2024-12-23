//go:build tool_langchain_test

// go test -cover -covermode=atomic -parallel 2 -tags=tool_langchain_test -v github.com/ai-flowx/flowx/tool

package tool

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nameLangChainTest = "langchain"
)

func initLangChainTest(_ context.Context) tool {
	cfg := Config{
		Provider: []Provider{
			{
				Type: typeLangChain,
				Name: nameLangChainTest,
			},
		},
	}

	return tool{
		cfg: &cfg,
	}
}

func TestLangChainInit(t *testing.T) {
	ctx := context.Background()
	_t := initLangChainTest(ctx)

	err := _t.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestLangChainDeinit(t *testing.T) {
	ctx := context.Background()
	_t := initLangChainTest(ctx)

	_ = _t.Init(ctx)

	err := _t.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestLangChainRun(t *testing.T) {
	ctx := context.Background()
	_t := initLangChainTest(ctx)

	_ = _t.Init(ctx)

	defer func(_t *tool, ctx context.Context) {
		_ = _t.Deinit(ctx)
	}(&_t, ctx)

	_, err := _t.Run(ctx, typeLangChain, nameLangChainTest, "arg")
	assert.Equal(t, nil, err)
}
