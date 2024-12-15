package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ai-flowx/flowx/store"
)

const (
	urlTestStore = "http://127.0.0.1:8082/"

	agentTestMemory = "testAgent"
	taskTestMemory  = "testTask"
	textTestMemory  = "testText"
)

func initMemoryTest(ctx context.Context) memory {
	c := store.DefaultConfig()
	c.Provider = store.ProviderChroma
	c.Url = urlTestStore

	cfg := Config{}
	cfg.Store = store.New(ctx, c)
	cfg.Type = typeShortTerm

	return memory{
		cfg: &cfg,
	}
}

func TestMemoryInit(t *testing.T) {
	ctx := context.Background()

	m := initMemoryTest(ctx)

	err := m.Init(ctx)
	assert.Equal(t, nil, err)
}

func TestMemoryDeinit(t *testing.T) {
	ctx := context.Background()

	m := initMemoryTest(ctx)

	err := m.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestMemoryReset(t *testing.T) {
	ctx := context.Background()

	m := initMemoryTest(ctx)

	_ = m.Init(ctx)

	defer func(m *memory, ctx context.Context) {
		_ = m.Deinit(ctx)
	}(&m, ctx)

	err := m.Reset(ctx)
	assert.Equal(t, nil, err)
}

func TestMemorySave(t *testing.T) {
	ctx := context.Background()

	m := initMemoryTest(ctx)

	_ = m.Init(ctx)

	defer func(m *memory, ctx context.Context) {
		_ = m.Deinit(ctx)
	}(&m, ctx)

	text := textTestMemory
	meta := map[string]interface{}{
		"task":    taskTestMemory,
		"quality": 0.5,
	}
	agent := agentTestMemory

	err := m.Save(ctx, text, meta, agent)
	assert.Equal(t, nil, err)
}

func TestMemorySearch(t *testing.T) {
	ctx := context.Background()

	m := initMemoryTest(ctx)

	_ = m.Init(ctx)

	defer func(m *memory, ctx context.Context) {
		_ = m.Deinit(ctx)
	}(&m, ctx)

	query := textTestMemory
	limit := 3
	threshold := 0.35

	_, err := m.Search(ctx, query, int32(limit), float32(threshold))
	assert.Equal(t, nil, err)
}
