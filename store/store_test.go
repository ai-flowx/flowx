package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nameTestStore = "testName"
	textTestStore = "testText"
	urlTestStore  = "http://127.0.0.1:8082/"
)

func initStoreTest(_ context.Context) store {
	cfg := Config{}
	cfg.Provider = ProviderChroma
	cfg.Url = urlTestStore

	return store{
		cfg: &cfg,
		st:  &Chroma{Url: cfg.Url},
	}
}

func TestStoreInit(t *testing.T) {
	ctx := context.Background()

	s := initStoreTest(ctx)

	err := s.Init(ctx, nameTestStore)
	assert.Equal(t, nil, err)
}

func TestStoreDeinit(t *testing.T) {
	ctx := context.Background()

	s := initStoreTest(ctx)

	err := s.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestStoreReset(t *testing.T) {
	ctx := context.Background()

	s := initStoreTest(ctx)

	_ = s.Init(ctx, nameTestStore)

	defer func(s *store, ctx context.Context) {
		_ = s.Deinit(ctx)
	}(&s, ctx)

	err := s.Reset(ctx)
	assert.Equal(t, nil, err)
}

func TestStoreSave(t *testing.T) {
	ctx := context.Background()

	s := initStoreTest(ctx)

	_ = s.Init(ctx, nameTestStore)

	defer func(s *store, ctx context.Context) {
		_ = s.Deinit(ctx)
	}(&s, ctx)

	text := textTestStore
	meta := map[string]interface{}{
		"task":    "testTask",
		"quality": 0.5,
	}
	agent := "testAgent"

	err := s.Save(ctx, text, meta, agent)
	assert.Equal(t, nil, err)
}

func TestStoreSearch(t *testing.T) {
	ctx := context.Background()

	s := initStoreTest(ctx)

	_ = s.Init(ctx, nameTestStore)

	defer func(s *store, ctx context.Context) {
		_ = s.Deinit(ctx)
	}(&s, ctx)

	query := textTestStore
	limit := 3
	threshold := 0.35

	_, err := s.Search(ctx, query, int32(limit), float32(threshold))
	assert.Equal(t, nil, err)
}
