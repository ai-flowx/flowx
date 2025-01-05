package task

import (
	"context"

	"github.com/ai-flowx/flowx/agent"
	"github.com/ai-flowx/flowx/tool"
)

type Task interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Description    string
	ExpectedOutput string
	Name           string
	Agent          agent.Agent
	Tool           []tool.Tool
	AsyncExecution bool
}

type task struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Task {
	return &task{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (t *task) Init(_ context.Context) error {
	return nil
}

func (t *task) Deinit(_ context.Context) error {
	return nil
}

func (t *task) Run(ctx context.Context) error {
	return nil
}
