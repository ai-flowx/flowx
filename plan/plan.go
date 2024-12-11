package plan

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

type Plan interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Addr   string
	Logger hclog.Logger
}

type plan struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Plan {
	return &plan{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (p *plan) Init(_ context.Context) error {
	return nil
}

func (p *plan) Deinit(_ context.Context) error {
	return nil
}

func (p *plan) Run(ctx context.Context) error {
	return nil
}