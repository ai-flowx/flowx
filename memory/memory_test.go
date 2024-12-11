package memory

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

type Memory interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Addr   string
	Logger hclog.Logger
}

type memory struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Memory {
	return &memory{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (m *memory) Init(_ context.Context) error {
	return nil
}

func (m *memory) Deinit(_ context.Context) error {
	return nil
}

func (m *memory) Run(ctx context.Context) error {
	return nil
}
