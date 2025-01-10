package flow

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ai-flowx/flowx/agent"
	"github.com/ai-flowx/flowx/gpt"
	"github.com/ai-flowx/flowx/memory"
	"github.com/ai-flowx/flowx/tool"
)

const (
	channelWeChat = "wechat"
)

type Flow interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Channel string
	Port    string
	Gpt     gpt.Gpt
	Memory  memory.Memory
	Tool    tool.Tool
	Agent   agent.Agent
}

type flow struct {
	cfg *Config
	fl  Flow
}

func New(_ context.Context, cfg *Config) Flow {
	var fl Flow

	if cfg.Channel == channelWeChat {
		fl = &WeChat{}
	} else {
		// BYPASS
	}

	return &flow{
		cfg: cfg,
		fl:  fl,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (f *flow) Init(ctx context.Context) error {
	if err := f.fl.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	return nil
}

func (f *flow) Deinit(ctx context.Context) error {
	if err := f.fl.Deinit(ctx); err != nil {
		return errors.Wrap(err, "failed to deinit\n")
	}

	return nil
}

func (f *flow) Run(ctx context.Context) error {
	if err := f.fl.Run(ctx); err != nil {
		return errors.Wrap(err, "failed to run\n")
	}

	return nil
}
