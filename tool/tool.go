package tool

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ai-flowx/flowx/gpt"
	"github.com/ai-flowx/toolx/decorator"
	"github.com/ai-flowx/toolx/gerrit"
	"github.com/ai-flowx/toolx/hello"
	"github.com/ai-flowx/toolx/structuredtool"
)

type Tool interface {
	Init(context.Context) error
	Deinit(context.Context) error
	List(context.Context) ([]Provider, error)
	Run(context.Context, string, ...interface{}) (string, error)
}

type Config struct {
	Gpt      gpt.Gpt
	Provider []Provider
}

type Provider struct {
	Name string
}

type tool struct {
	cfg      *Config
	callback func(context.Context, interface{}) (interface{}, error)
	toolx    []ToolX
}

func New(_ context.Context, cfg *Config) Tool {
	return &tool{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (t *tool) Init(ctx context.Context) error {
	var err error

	if err = t.initGpt(ctx); err != nil {
		return errors.Wrap(err, "failed to init gpt\n")
	}

	if err = t.initProvider(ctx); err != nil {
		return errors.Wrap(err, "failed to init provider\n")
	}

	for _, item := range t.toolx {
		if err = item.Init(ctx); err != nil {
			break
		}
	}

	return err
}

func (t *tool) Deinit(ctx context.Context) error {
	for _, item := range t.toolx {
		if err := item.Deinit(ctx); err != nil {
			return err
		}
	}

	if err := t.deinitProvider(ctx); err != nil {
		return err
	}

	if err := t.deinitGpt(ctx); err != nil {
		return err
	}

	return nil
}

func (t *tool) List(_ context.Context) ([]Provider, error) {
	return t.cfg.Provider, nil
}

func (t *tool) Run(ctx context.Context, name string, args ...interface{}) (string, error) {
	var found bool
	var res string
	var err error

	for _, item := range t.toolx {
		if item.Name(ctx) == name {
			found = true
			res, err = item.Call(ctx, t.callback, args...)
			break
		}
	}

	if !found {
		err = errors.New("invalid name\n")
	}

	return res, err
}

func (t *tool) initProvider(_ context.Context) error {
	t.toolx = append(t.toolx,
		&decorator.Decorator{},
		&gerrit.Gerrit{},
		&hello.Hello{},
		&structuredtool.StructuredTool{})

	return nil
}

func (t *tool) deinitProvider(_ context.Context) error {
	t.toolx = t.toolx[:0]

	return nil
}

func (t *tool) initGpt(ctx context.Context) error {
	if err := t.cfg.Gpt.Init(ctx); err != nil {
		return err
	}

	t.callback = t.wrapCallback(t.cfg.Gpt.Chat)

	return nil
}

func (t *tool) deinitGpt(ctx context.Context) error {
	t.callback = nil

	return t.cfg.Gpt.Deinit(ctx)
}

// nolint:lll
func (t *tool) wrapCallback(fn func(context.Context, *gpt.ChatRequest) (gpt.ChatResponse, error)) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*gpt.ChatRequest)
		if !ok {
			return nil, errors.New("invalid request\n")
		}

		res, err := fn(ctx, req)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
