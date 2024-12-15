package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ai-flowx/flowx/store"
)

const (
	agentTestShortTerm = "testAgent"
	nameTestShortTerm  = "testName"
	taskTestShortTerm  = "testTask"
	textTestShortTerm  = "testText"
)

func initShortTermTest(ctx context.Context) ShortTerm {
	c := store.DefaultConfig()
	c.Provider = store.ProviderChroma

	s := store.New(ctx, c)

	return ShortTerm{
		Store: s,
		Name:  nameTestShortTerm,
	}
}

func TestShortTermInit(t *testing.T) {
	ctx := context.Background()

	l := initShortTermTest(ctx)

	err := l.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestShortTermDeinit(t *testing.T) {
	ctx := context.Background()

	l := initShortTermTest(ctx)

	err := l.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestShortTermReset(t *testing.T) {
	ctx := context.Background()

	l := initShortTermTest(ctx)

	_ = l.Init(ctx)

	defer func(l *ShortTerm, ctx context.Context) {
		_ = l.Deinit(ctx)
	}(&l, ctx)

	err := l.Reset(ctx)
	assert.Equal(t, nil, err)
}

func TestShortTermSave(t *testing.T) {
	ctx := context.Background()

	l := initShortTermTest(ctx)

	_ = l.Init(ctx)

	defer func(l *ShortTerm, ctx context.Context) {
		_ = l.Deinit(ctx)
	}(&l, ctx)

	text := textTestShortTerm
	meta := map[string]interface{}{
		"task": taskTestShortTerm,
	}
	agent := agentTestShortTerm

	err := l.Save(ctx, text, meta, agent)
	assert.Equal(t, nil, err)
}

func TestShortTermSearch(t *testing.T) {
	ctx := context.Background()

	l := initShortTermTest(ctx)

	_ = l.Init(ctx)

	defer func(l *ShortTerm, ctx context.Context) {
		_ = l.Deinit(ctx)
	}(&l, ctx)

	query := textTestShortTerm
	limit := 3
	threshold := 0.35

	_, err := l.Search(ctx, query, int32(limit), float32(threshold))
	assert.Equal(t, nil, err)
}
