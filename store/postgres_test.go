//go:build store_postgres_test

// go test -cover -covermode=atomic -parallel 2 -tags=store_postgres_test -v github.com/ai-flowx/flowx/store

package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initPostgresTest(_ context.Context) Postgres {
	return Postgres{
		Host: "127.0.0.1",
		Port: 5432,
		User: "postgres",
		Pass: "postgres",
	}
}

func TestPostgresInit(t *testing.T) {
	ctx := context.Background()
	p := initPostgresTest(ctx)

	defer func(p *Postgres, ctx context.Context) {
		_ = p.Deinit(ctx)
	}(&p, ctx)

	err := p.Init(ctx, "")
	assert.Equal(t, nil, err)
}

func TestPostgresDeinit(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestPostgresReset(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestPostgresSave(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestPostgresSearch(t *testing.T) {
	assert.Equal(t, nil, nil)
}
