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
		Api:      "https://ark.cn-beijing.volces.com/api/v3/chat/completions",
		Key:      "8429f8ab-*",
		Endpoint: "ep-*",
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

	r := ChatRequest{
		Messages: []ChatMessage{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: "Hello!",
			},
		},
	}

	_, err := c.Chat(ctx, &r)
	assert.Equal(t, nil, err)
}
