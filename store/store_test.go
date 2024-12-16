//go:build store_test

// go test -cover -covermode=atomic -parallel 2 -tags=store_test -v github.com/ai-flowx/flowx/store

package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	assert.Equal(t, nil, nil)
}
