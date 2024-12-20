package tool

import (
	"context"

	"github.com/pkg/errors"
)

const (
	typeCrewAi    = "crewai"
	typeFlowX     = "flowx"
	typeLangChain = "langchain"
)

type Tool interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, []*Invoke) error
}

type Config struct {
	Type string
}

type Invoke struct {
	Name        string
	Description string
	Path        string
	Func        Func
	Args        map[string]interface{}
	Result      string
}

type Func func(ctx context.Context, args ...interface{}) ([]byte, error)

type tool struct {
	cfg       *Config
	crewai    *CrewAi
	flowx     *FlowX
	langchain *LangChain
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

	if t.cfg.Type == typeCrewAi {
		t.crewai = &CrewAi{}
		err = t.crewai.Init(ctx)
	} else if t.cfg.Type == typeFlowX {
		t.flowx = &FlowX{}
		err = t.flowx.Init(ctx)
	} else if t.cfg.Type == typeLangChain {
		t.langchain = &LangChain{}
		err = t.langchain.Init(ctx)
	} else {
		err = errors.New("invalid tool type\n")
	}

	return err
}

func (t *tool) Deinit(ctx context.Context) error {
	var err error

	if t.cfg.Type == typeCrewAi {
		if t.crewai != nil {
			err = t.crewai.Deinit(ctx)
		}
	} else if t.cfg.Type == typeFlowX {
		if t.flowx != nil {
			err = t.flowx.Deinit(ctx)
		}
	} else if t.cfg.Type == typeLangChain {
		if t.langchain != nil {
			err = t.langchain.Deinit(ctx)
		}
	} else {
		err = errors.New("invalid tool type\n")
	}

	return err
}

func (t *tool) Run(ctx context.Context, invokes []*Invoke) error {
	var err error

	if t.cfg.Type == typeCrewAi {
		if t.crewai != nil {
			err = t.crewai.Run(ctx, invokes)
		}
	} else if t.cfg.Type == typeFlowX {
		if t.flowx != nil {
			err = t.flowx.Run(ctx, invokes)
		}
	} else if t.cfg.Type == typeLangChain {
		if t.langchain != nil {
			err = t.langchain.Run(ctx, invokes)
		}
	} else {
		err = errors.New("invalid tool type\n")
	}

	return err
}
