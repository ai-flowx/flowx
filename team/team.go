package team

import (
	"context"
	"github.com/ai-flowx/flowx/agent"
	"github.com/ai-flowx/flowx/task"
)

type Team interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Tasks   []task.Task
	Agents  []agent.Agent
	Process string
	MaxRpm  int
	Memory  bool
	Cache   bool
	Plan    bool
}

type team struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Team {
	return &team{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (t *team) Init(_ context.Context) error {
	return nil
}

func (t *team) Deinit(_ context.Context) error {
	return nil
}

func (t *team) Run(ctx context.Context) error {
	return nil
}
