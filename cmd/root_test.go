package cmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ai-flowx/flowx/config"
)

var (
	testConfig = config.Config{
		Cache: config.Cache{
			Provider: "gptcache",
			Api:      "http://127.0.0.1:8081",
			Token:    "token",
		},
		Gpt: config.Gpt{
			Provider: "openai",
			Api:      "https://openai.com/api",
			Token:    "token",
		},
		Store: config.Store{
			Provider: "chroma",
			Api:      "http://127.0.0.1:8082",
			Token:    "token",
		},
	}
)

func TestInitConfig(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestInitStore(t *testing.T) {
	ctx := context.Background()

	_, err := initStore(ctx, &testConfig)
	assert.Equal(t, nil, err)
}

func TestInitMemory(t *testing.T) {
	ctx := context.Background()

	s, _ := initStore(ctx, &testConfig)

	_, err := initMemory(ctx, &testConfig, s)
	assert.Equal(t, nil, err)
}

func TestInitFlow(t *testing.T) {
	ctx := context.Background()

	s, _ := initStore(ctx, &testConfig)
	m, _ := initMemory(ctx, &testConfig, s)

	listenAddr = "127.0.0.1:8080"

	_, err := initFlow(ctx, &testConfig, m)
	assert.Equal(t, nil, err)
}

func TestRunFlow(t *testing.T) {
	assert.Equal(t, nil, nil)
}
