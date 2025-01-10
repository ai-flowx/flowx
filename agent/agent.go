package agent

import (
	"context"
	"time"

	"github.com/ai-flowx/flowx/gpt"
	"github.com/ai-flowx/flowx/prompt"
	"github.com/ai-flowx/flowx/tool"
)

type Agent interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, string, string) (string, error)
}

type Config struct {
	Role                 string
	Goal                 string
	Backstory            string
	Gpt                  gpt.Gpt
	Tool                 tool.Tool
	MaxIter              int
	MaxRpm               int
	MaxExecutionTime     time.Duration
	Memory               bool
	Cache                bool
	PromptTemplate       string
	ResponseTemplate     string
	RespectContextWindow bool
	Prompt               prompt.Prompt
}

type agent struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Agent {
	return &agent{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (a *agent) Init(ctx context.Context) error {
	return nil
}

func (a *agent) Deinit(ctx context.Context) error {
	return nil
}

func (a *agent) Run(ctx context.Context, _prompt, _context string) (string, error) {
	return "", nil
}
