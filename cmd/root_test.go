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
			Provider: "cachex",
			Host:     "127.0.0.1",
			Port:     8081,
			User:     "user",
			Pass:     "pass",
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
			Provider: "vecx",
			Host:     "127.0.0.1",
			Port:     8082,
			Path:     "/path/to/file",
			User:     "user",
			Pass:     "pass",
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

func TestInitPrompt(t *testing.T) {
	ctx := context.Background()

	_, err := initPrompt(ctx, &testConfig)
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

	g, _ := initGpt(ctx, &testConfig)

	_, err := initTool(ctx, &testConfig, g)
	assert.Equal(t, nil, err)
}

func TestInitAgent(t *testing.T) {
	ctx := context.Background()

	g, _ := initGpt(ctx, &testConfig)
	p, _ := initPrompt(ctx, &testConfig)
	_t, _ := initTool(ctx, &testConfig, g)

	_, err := initAgent(ctx, &testConfig, g, p, _t)
	assert.Equal(t, nil, err)
}

func TestInitFlow(t *testing.T) {
	ctx := context.Background()

	listenPort = ":8080"

	g, _ := initGpt(ctx, &testConfig)
	p, _ := initPrompt(ctx, &testConfig)
	s, _ := initStore(ctx, &testConfig)
	m, _ := initMemory(ctx, &testConfig, s)
	_t, _ := initTool(ctx, &testConfig, g)
	a, _ := initAgent(ctx, &testConfig, g, p, _t)

	_, err := initFlow(ctx, &testConfig, g, m, _t, a)
	assert.Equal(t, nil, err)
}

func TestRunFlow(t *testing.T) {
	assert.Equal(t, nil, nil)
}
