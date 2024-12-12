package tool

import (
	"context"
)

type Tool interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct{}

type tool struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Tool {
	return &tool{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (t *tool) Init(_ context.Context) error {
	return nil
}

func (t *tool) Deinit(_ context.Context) error {
	return nil
}

func (t *tool) Run(ctx context.Context) error {
	return nil
}
