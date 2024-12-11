package cachex

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

type Cachex interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Addr   string
	Logger hclog.Logger
}

type cachex struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Cachex {
	return &cachex{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (c *cachex) Init(_ context.Context) error {
	return nil
}

func (c *cachex) Deinit(_ context.Context) error {
	return nil
}

func (c *cachex) Run(ctx context.Context) error {
	return nil
}
