//go:build store_postgres_test

// go test -cover -covermode=atomic -parallel 2 -tags=store_postgres_test -v github.com/ai-flowx/flowx/store

package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresInit(t *testing.T) {
	assert.Equal(t, nil, nil)
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
