//go:build gpt_doubao_test

// go test -cover -covermode=atomic -parallel 2 -tags=gpt_doubao_test -v github.com/ai-flowx/flowx/gpt

package gpt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initDoubaoChatTest(_ context.Context) DoubaoChat {
	return DoubaoChat{
		Api:   "https://ark.cn-beijing.volces.com/api/v3/chat/completions",
		Model: "ep-*",
		Key:   "8429f8ab-*",
	}
}

func initDoubaoVisionTest(_ context.Context) DoubaoVision {
	return DoubaoVision{
		Api:   "https://ark.cn-beijing.volces.com/api/v3/chat/completions",
		Model: "ep-*",
		Key:   "8429f8ab-*",
	}
}

func TestDoubaoChatInit(t *testing.T) {
	ctx := context.Background()
	c := initDoubaoChatTest(ctx)

	err := c.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestDoubaoChatDeinit(t *testing.T) {
	ctx := context.Background()
	c := initDoubaoChatTest(ctx)

	_ = c.Init(ctx)

	err := c.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestDoubaoChatRun(t *testing.T) {
	ctx := context.Background()
	c := initDoubaoChatTest(ctx)

	_ = c.Init(ctx)

	defer func(c *DoubaoChat, ctx context.Context) {
		_ = c.Deinit(ctx)
	}(&c, ctx)

	_, err := c.Run(ctx, "hello world")
	assert.Equal(t, nil, err)
}

func TestDoubaoVisionInit(t *testing.T) {
	ctx := context.Background()
	v := initDoubaoVisionTest(ctx)

	err := v.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestDoubaoVisionDeinit(t *testing.T) {
	ctx := context.Background()
	v := initDoubaoVisionTest(ctx)

	_ = v.Init(ctx)

	err := v.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestDoubaoVisionRun(t *testing.T) {
	ctx := context.Background()
	v := initDoubaoVisionTest(ctx)

	_ = v.Init(ctx)

	defer func(v *DoubaoVision, ctx context.Context) {
		_ = v.Deinit(ctx)
	}(&v, ctx)

	_, err := v.Run(ctx, "png-base64")
	assert.Equal(t, nil, err)
}
