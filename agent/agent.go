package agent

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/ai-flowx/flowx/gpt"
	"github.com/ai-flowx/flowx/prompt"
	"github.com/ai-flowx/flowx/tool"
)

const (
	routineNum = -1
)

type Agent interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, string) (string, error)
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

type Action struct {
	Thought string
	Answer  string
	Tool    Tool
}

type Tool struct {
	Name  string
	Input string
}

type agent struct {
	cfg  *Config
	iter int
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

func (a *agent) Run(ctx context.Context, content string) (string, error) {
	var message string
	var answer string
	var action Action
	var result string
	var err error

	ctx, cancel := context.WithTimeout(ctx, a.cfg.MaxExecutionTime)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(routineNum)

	g.Go(func() error {
		if a.maxIterExceeded(ctx) {
			return errors.New("max iteration exceeded\n")
		}
		if a.maxRpmExceeded(ctx) {
			return errors.New("max rpm exceeded\n")
		}
		if message, err = a.formatMessage(ctx, content); err != nil {
		}
		if answer, err = a.invokeGpt(ctx, message); err != nil {
		}
		if action, err = a.parseAnswer(ctx, answer); err != nil {
			if action.Tool.Name != "" {
				if result, err = a.invokeTool(ctx, action.Tool); err != nil {
				}
			} else {
				result = action.Answer
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return "", err
	}

	if err = a.createMemory(ctx, result); err != nil {
		return "", errors.Wrap(err, "failed to create memory\n")
	}

	return result, nil
}

func (a *agent) maxIterExceeded(_ context.Context) bool {
	return a.iter >= a.cfg.MaxIter
}

func (a *agent) maxRpmExceeded(_ context.Context) bool {
	// TBD: FIXME
	return true
}

func (a *agent) formatMessage(_ context.Context, content string) (string, error) {
	// TBD: FIXME
	return "", nil
}

func (a *agent) invokeGpt(_ context.Context, content string) (string, error) {
	var answer string

	a.iter += 1
	// TBD: FIXME

	if a.cfg.RespectContextWindow {
		// TBD: FIXME
	}

	return answer, nil
}

func (a *agent) parseAnswer(ctx context.Context, content string) (Action, error) {
	p := parser{}

	act, err := p.parse(ctx, content)
	if err != nil {
		return Action{}, errors.Wrap(err, "failed to parse answer\n")
	}

	return act, nil
}

func (a *agent) invokeTool(_ context.Context, _tool Tool) (string, error) {
	// TBD: FIXME
	return "", nil
}

func (a *agent) createMemory(_ context.Context, content string) error {
	// TBD: FIXME
	return nil
}
