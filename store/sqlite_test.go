//go:build store_sqlite_test

// go test -cover -covermode=atomic -parallel 2 -tags=store_sqlite_test -v github.com/ai-flowx/flowx/store

package store

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	pathTestSqlite = "/tmp/testSqlite.db"

	agentTestSqlite = "agentTestSqlite"
	nameTestSqlite  = "nameTestSqlite"
	textTestSqlite  = "textTestSqlite"
)

func initSqliteTest(_ context.Context) Sqlite {
	return Sqlite{
		Path: pathTestSqlite,
	}
}

func deinitSqliteTest(_ context.Context) {
	_ = os.Remove(pathTestSqlite)
}

func TestSqliteInit(t *testing.T) {
	ctx := context.Background()
	s := initSqliteTest(ctx)

	defer deinitSqliteTest(ctx)

	err := s.Init(ctx, nameTestSqlite)
	assert.Equal(t, nil, err)
}

func TestSqliteDeinit(t *testing.T) {
	ctx := context.Background()
	s := initSqliteTest(ctx)

	defer deinitSqliteTest(ctx)

	_ = s.Init(ctx, nameTestSqlite)

	err := s.Deinit(ctx)
	assert.Equal(t, nil, err)
}

func TestSqliteReset(t *testing.T) {
	ctx := context.Background()
	s := initSqliteTest(ctx)

	_ = s.Init(ctx, nameTestSqlite)

	defer func(s *Sqlite, ctx context.Context) {
		_ = s.Deinit(ctx)
		deinitSqliteTest(ctx)
	}(&s, ctx)

	err := s.Reset(ctx)
	assert.Equal(t, nil, err)
}

func TestSqliteSave(t *testing.T) {
	ctx := context.Background()
	s := initSqliteTest(ctx)

	_ = s.Init(ctx, nameTestSqlite)

	defer func(s *Sqlite, ctx context.Context) {
		_ = s.Reset(ctx)
		_ = s.Deinit(ctx)
		deinitSqliteTest(ctx)
	}(&s, ctx)

	text := textTestSqlite
	meta := map[string]interface{}{
		"key": "value",
	}
	agent := agentTestSqlite

	err := s.Save(ctx, text, meta, agent)
	assert.Equal(t, nil, err)
}

func TestSqliteSearch(t *testing.T) {
	ctx := context.Background()
	s := initSqliteTest(ctx)

	_ = s.Init(ctx, nameTestSqlite)

	defer func(s *Sqlite, ctx context.Context) {
		_ = s.Reset(ctx)
		_ = s.Deinit(ctx)
		deinitSqliteTest(ctx)
	}(&s, ctx)

	text := textTestSqlite
	meta := map[string]interface{}{
		"key": "value",
	}
	agent := agentTestSqlite

	_ = s.Save(ctx, text, meta, agent)

	query := textTestSqlite
	limit := 3
	threshold := 0.35

	_, err := s.Search(ctx, query, int32(limit), float32(threshold))
	assert.Equal(t, nil, err)
}
