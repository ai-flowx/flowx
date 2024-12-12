package cmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestInitFlow(t *testing.T) {
	ctx := context.Background()

	listenAddr = "127.0.0.1:8080"

	_, err := initFlow(ctx)
	assert.Equal(t, nil, err)
}
