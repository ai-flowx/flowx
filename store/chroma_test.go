package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	urlTestChroma = "http://47.88.100.1:8082"

	agentTestChroma = "agentTestChroma"
	nameTestChroma  = "nameTestChroma"
	textTestChroma  = "textTestChroma"
)

func initChromaTest(_ context.Context) Chroma {
	return Chroma{
		Url: urlTestChroma,
	}
}

func TestChromaInit(t *testing.T) {
	ctx := context.Background()
	c := initChromaTest(ctx)

	err := c.Init(ctx, nameTestChroma)
	assert.Equal(t, nil, err)
}

func TestChromaDeinit(t *testing.T) {
	ctx := context.Background()
	c := initChromaTest(ctx)

	_ = c.Init(ctx, nameTestChroma)

	err := c.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestChromaReset(t *testing.T) {
	ctx := context.Background()
	c := initChromaTest(ctx)

	_ = c.Init(ctx, nameTestChroma)

	defer func(c *Chroma, ctx context.Context) {
		_ = c.Deinit(ctx)
	}(&c, ctx)

	err := c.Reset(ctx)
	assert.Equal(t, nil, err)
}

func TestChromaSave(t *testing.T) {
	ctx := context.Background()
	c := initChromaTest(ctx)

	_ = c.Init(ctx, nameTestChroma)

	defer func(c *Chroma, ctx context.Context) {
		_ = c.Reset(ctx)
		_ = c.Deinit(ctx)
	}(&c, ctx)

	text := textTestChroma
	meta := map[string]interface{}{
		"key": "value",
	}
	agent := agentTestChroma

	err := c.Save(ctx, text, meta, agent)
	assert.Equal(t, nil, err)
}

func TestChromaSearch(t *testing.T) {
	ctx := context.Background()
	c := initChromaTest(ctx)

	_ = c.Init(ctx, nameTestChroma)

	defer func(c *Chroma, ctx context.Context) {
		_ = c.Reset(ctx)
		_ = c.Deinit(ctx)
	}(&c, ctx)

	text := textTestChroma
	meta := map[string]interface{}{
		"key": "value",
	}
	agent := agentTestChroma

	_ = c.Save(ctx, text, meta, agent)

	query := textTestChroma
	limit := 3
	threshold := 0.35

	_, err := c.Search(ctx, query, int32(limit), float32(threshold))
	assert.Equal(t, nil, err)
}
