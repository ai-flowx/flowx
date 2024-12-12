package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ai-flowx/flowx/store"
)

const (
	agentTestLongTerm = "testAgent"
	taskTestLongTerm  = "testTask"
	valueTestLongTerm = "testValue"
)

func initLongTermTest(ctx context.Context) LongTerm {
	c := store.DefaultConfig()
	c.Provider = store.ProviderChroma
	s := store.New(ctx, c)

	return LongTerm{
		Store: s,
	}
}

func TestLongTermInit(t *testing.T) {
	ctx := context.Background()

	l := initLongTermTest(ctx)

	err := l.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestLongTermDeinit(t *testing.T) {
	ctx := context.Background()

	l := initLongTermTest(ctx)

	err := l.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestLongTermReset(t *testing.T) {
	ctx := context.Background()

	l := initLongTermTest(ctx)

	_ = l.Init(ctx)

	defer func(l *LongTerm, ctx context.Context) {
		_ = l.Deinit(ctx)
	}(&l, ctx)

	err := l.Reset(ctx)
	assert.Equal(t, nil, err)
}

func TestLongTermSave(t *testing.T) {
	ctx := context.Background()

	l := initLongTermTest(ctx)

	_ = l.Init(ctx)

	defer func(l *LongTerm, ctx context.Context) {
		_ = l.Deinit(ctx)
	}(&l, ctx)

	value := valueTestLongTerm
	meta := map[string]interface{}{
		"task":    taskTestLongTerm,
		"quality": 0.5,
	}
	agent := agentTestLongTerm

	err := l.Save(ctx, value, meta, agent)
	assert.Equal(t, nil, err)
}

func TestLongTermSearch(t *testing.T) {
	ctx := context.Background()

	l := initLongTermTest(ctx)

	_ = l.Init(ctx)

	defer func(l *LongTerm, ctx context.Context) {
		_ = l.Deinit(ctx)
	}(&l, ctx)

	query := valueTestLongTerm
	limit := 3
	threshold := 0.35

	_, err := l.Search(ctx, query, limit, threshold)
	assert.Equal(t, nil, err)
}
