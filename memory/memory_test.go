package memory

import (
	"context"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"

	"github.com/ai-flowx/flowx/store"
)

const (
	valueTestMemory = "testMemory"
)

func initMemoryTest(ctx context.Context) memory {
	cfg := Config{}

	cfg.Logger = hclog.New(&hclog.LoggerOptions{
		Name:  "memory",
		Level: hclog.LevelFromString("info")})

	c := store.DefaultConfig()
	c.Logger = hclog.New(&hclog.LoggerOptions{
		Name:  "store",
		Level: hclog.LevelFromString("info")})
	c.Provider = store.ProviderChroma
	cfg.Store = store.New(ctx, c)

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

	value := valueTestMemory
	meta := map[string]interface{}{
		"task":    "testTask",
		"quality": 0.5,
	}
	agent := "testAgent"

	err := m.Save(ctx, value, meta, agent)
	assert.Equal(t, nil, err)
}

func TestMemorySearch(t *testing.T) {
	ctx := context.Background()

	m := initMemoryTest(ctx)

	_ = m.Init(ctx)

	defer func(m *memory, ctx context.Context) {
		_ = m.Deinit(ctx)
	}(&m, ctx)

	query := valueTestMemory
	limit := 3
	threshold := 0.35

	_, err := m.Search(ctx, query, limit, threshold)
	assert.Equal(t, nil, err)
}
