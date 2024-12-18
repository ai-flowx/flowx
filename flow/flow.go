package flow

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/ai-flowx/flowx/memory"
)

const (
	routineNum = -1
)

type Flow interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Memory memory.Memory
	Port   string
}

type flow struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Flow {
	return &flow{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (f *flow) Init(_ context.Context) error {
	return nil
}

func (f *flow) Deinit(_ context.Context) error {
	return nil
}

func (f *flow) Run(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)
	g.SetLimit(routineNum)

	g.Go(func() error {
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
