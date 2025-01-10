package prompt

import (
	"context"
)

type Prompt interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct{}

type prompt struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Prompt {
	return &prompt{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (p *prompt) Init(ctx context.Context) error {
	return nil
}

func (p *prompt) Deinit(ctx context.Context) error {
	return nil
}

func (p *prompt) Run(ctx context.Context) error {
	return nil
}
