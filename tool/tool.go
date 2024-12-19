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
	Run(context.Context) error
}

type Config struct {
	Type string
}

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
		t.crewai = &CrewAi{
			Type: t.cfg.Type,
		}
		err = t.crewai.Init(ctx)
	} else if t.cfg.Type == typeFlowX {
		t.flowx = &FlowX{
			Type: t.cfg.Type,
		}
		err = t.flowx.Init(ctx)
	} else if t.cfg.Type == typeLangChain {
		t.langchain = &LangChain{
			Type: t.cfg.Type,
		}
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

func (t *tool) Run(ctx context.Context) error {
	var err error

	if t.cfg.Type == typeCrewAi {
		if t.crewai != nil {
			err = t.crewai.Run(ctx)
		}
	} else if t.cfg.Type == typeFlowX {
		if t.flowx != nil {
			err = t.flowx.Run(ctx)
		}
	} else if t.cfg.Type == typeLangChain {
		if t.langchain != nil {
			err = t.langchain.Run(ctx)
		}
	} else {
		err = errors.New("invalid tool type\n")
	}

	return err
}
