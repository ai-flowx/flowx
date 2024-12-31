package gpt

import (
	"context"

	"github.com/pkg/errors"
)

const (
	providerDoubaoChat   = "doubao-chat"
	providerDoubaoVision = "doubao-vision"
)

type Gpt interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, string) ([]string, error)
}

type Config struct {
	Provider string
	Api      string
	Model    string
	Key      string
}

type gpt struct {
	cfg *Config
	gt  Gpt
}

func New(_ context.Context, cfg *Config) Gpt {
	var gt Gpt

	if cfg.Provider == providerDoubaoChat {
		gt = &DoubaoChat{
			Api:   cfg.Api,
			Model: cfg.Model,
			Key:   cfg.Key,
		}
	} else if cfg.Provider == providerDoubaoVision {
		gt = &DoubaoVision{
			Api:   cfg.Api,
			Model: cfg.Model,
			Key:   cfg.Key,
		}
	} else {
		// BYPASS
	}

	return &gpt{
		cfg: cfg,
		gt:  gt,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (g *gpt) Init(ctx context.Context) error {
	if err := g.gt.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	return nil
}

func (g *gpt) Deinit(ctx context.Context) error {
	if err := g.gt.Deinit(ctx); err != nil {
		return errors.Wrap(err, "failed to deinit\n")
	}

	return nil
}

func (g *gpt) Run(ctx context.Context, content string) ([]string, error) {
	buf, err := g.gt.Run(ctx, content)
	if err != nil {
		return nil, errors.Wrap(err, "failed to run\n")
	}

	return buf, nil
}
