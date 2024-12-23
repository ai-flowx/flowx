package tool

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ai-flowx/toolx/hello"
	"github.com/ai-flowx/toolx/langchain"
)

const (
	typeCrewAi    = "crewai"
	typeLangChain = "langchain"
	typeToolX     = "toolx"
)

type Tool interface {
	Init(context.Context) error
	Deinit(context.Context) error
	List(context.Context) ([]Provider, error)
	Run(context.Context, string, string, ...interface{}) (string, error)
}

type Config struct {
	Provider []Provider
}

type Provider struct {
	Type string
	Name string
}

type tool struct {
	cfg       *Config
	crewai    []CrewAi
	langchain []LangChain
	toolx     []ToolX
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

	if err = t.initProvider(ctx); err != nil {
		return errors.Wrap(err, "failed to init provider\n")
	}

	for _, item := range t.crewai {
		if err = item.Init(ctx); err != nil {
			break
		}
	}

	for _, item := range t.langchain {
		if err = item.Init(ctx); err != nil {
			break
		}
	}

	for _, item := range t.toolx {
		if err = item.Init(ctx); err != nil {
			break
		}
	}

	return err
}

func (t *tool) Deinit(ctx context.Context) error {
	var err error

	for _, item := range t.toolx {
		if err = item.Deinit(ctx); err != nil {
			break
		}
	}

	for _, item := range t.langchain {
		if err = item.Deinit(ctx); err != nil {
			break
		}
	}

	for _, item := range t.crewai {
		if err = item.Deinit(ctx); err != nil {
			break
		}
	}

	return t.deinitProvider(ctx)
}

func (t *tool) List(_ context.Context) ([]Provider, error) {
	return t.cfg.Provider, nil
}

func (t *tool) Run(ctx context.Context, _type, name string, args ...interface{}) (string, error) {
	var found bool
	var res string
	var err error

	if _type == typeCrewAi {
		err = errors.New("TBD: FIXME\n")
	} else if _type == typeLangChain {
		for _, item := range t.langchain {
			if item.Name(ctx) == name {
				found = true
				res, err = item.Call(ctx, args)
				break
			}
		}
	} else if _type == typeToolX {
		for _, item := range t.toolx {
			if item.Name(ctx) == name {
				found = true
				res, err = item.Call(ctx, args)
				break
			}
		}
	} else {
		err = errors.New("invalid type\n")
	}

	if !found {
		err = errors.New("invalid name\n")
	}

	return res, err
}

func (t *tool) initProvider(_ context.Context) error {
	var err error

	for _, item := range t.cfg.Provider {
		if item.Type == typeCrewAi {
			err = errors.New("TBD: FIXME\n")
		} else if item.Type == typeLangChain {
			t.langchain = append(t.langchain, langchain.LangChain{})
		} else if item.Type == typeToolX {
			t.toolx = append(t.toolx, hello.Hello{})
		} else {
			err = errors.New("invalid type\n")
		}
	}

	return err
}

func (t *tool) deinitProvider(_ context.Context) error {
	t.toolx = t.toolx[:0]
	t.langchain = t.langchain[:0]
	t.crewai = t.crewai[:0]

	return nil
}
