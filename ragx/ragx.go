package ragx

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

type Ragx interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Addr   string
	Logger hclog.Logger
}

type ragx struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Ragx {
	return &ragx{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (r *ragx) Init(_ context.Context) error {
	return nil
}

func (r *ragx) Deinit(_ context.Context) error {
	return nil
}

func (r *ragx) Run(ctx context.Context) error {
	return nil
}
