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
			Url:      "http://127.0.0.1:8081",
		},
		Flow: config.Flow{
			Channel: "wechat",
		},
		Gpt: config.Gpt{
			Provider: "doubao-chat",
			Api:      "https://ark.cn-beijing.volces.com/api/v3/chat/completions",
			Model:    "ep-*",
			Key:      "8429f8ab-*",
		},
		Memory: config.Memory{
			Type: "shortterm",
		},
		Store: config.Store{
			Provider: "chroma",
			Url:      "http://127.0.0.1:8082",
			Path:     "",
		},
		Tool: []config.Tool{
			{
				Name: "hello",
			},
		},
	}
)

func TestInitConfig(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestInitGpt(t *testing.T) {
	ctx := context.Background()

	_, err := initGpt(ctx, &testConfig)
	assert.Equal(t, nil, err)
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

func TestInitTool(t *testing.T) {
	ctx := context.Background()

	_, err := initTool(ctx, &testConfig)
	assert.Equal(t, nil, err)
}

func TestInitFlow(t *testing.T) {
	ctx := context.Background()

	listenPort = ":8080"

	g, _ := initGpt(ctx, &testConfig)
	s, _ := initStore(ctx, &testConfig)
	m, _ := initMemory(ctx, &testConfig, s)
	_t, _ := initTool(ctx, &testConfig)

	_, err := initFlow(ctx, &testConfig, g, m, _t)
	assert.Equal(t, nil, err)
}

func TestRunFlow(t *testing.T) {
	assert.Equal(t, nil, nil)
}
