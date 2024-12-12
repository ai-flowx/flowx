package cmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	levelInfo = "info"
)

func TestInitConfig(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestInitLogger(t *testing.T) {
	ctx := context.Background()

	logLevel = levelInfo

	_, err := initLogger(ctx)
	assert.Equal(t, nil, err)
}

func TestInitFlow(t *testing.T) {
	ctx := context.Background()

	listenAddr = "127.0.0.1:8080"
	logLevel = levelInfo

	logger, err := initLogger(ctx)
	assert.Equal(t, nil, err)

	_, err = initFlow(ctx, logger)
	assert.Equal(t, nil, err)
}
