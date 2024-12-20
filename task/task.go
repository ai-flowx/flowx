package task

import (
	"context"
)

type Task interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct{}

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
