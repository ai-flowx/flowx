package flow

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
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
	Addr   string
	Logger hclog.Logger
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
	s := make(chan string, 1)

	g, _ := errgroup.WithContext(ctx)
	g.SetLimit(routineNum)

	g.Go(func() error {
		<-s
		return nil
	})

	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "failed to wait\n")
	}

	return nil
}
