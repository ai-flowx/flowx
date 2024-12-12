package proto

import (
	"context"
)

type Proto interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct{}

type proto struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Proto {
	return &proto{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (p *proto) Init(_ context.Context) error {
	return nil
}

func (p *proto) Deinit(_ context.Context) error {
	return nil
}

func (p *proto) Run(ctx context.Context) error {
	return nil
}
