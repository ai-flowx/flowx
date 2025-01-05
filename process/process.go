package process

import (
	"context"
)

const (
	Hierarchical = "hierarchical"
	Sequential   = "sequential"
)

type Process interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct{}

type process struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Process {
	return &process{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (p *process) Init(_ context.Context) error {
	return nil
}

func (p *process) Deinit(_ context.Context) error {
	return nil
}

func (p *process) Run(ctx context.Context) error {
	return nil
}
