package store

import (
	"context"

	"github.com/pkg/errors"
)

const (
	ProviderChroma = "chroma"
	ProviderRagx   = "ragx"
)

type Store interface {
	Init(context.Context, string) error
	Deinit(context.Context) error
	Reset(context.Context) error
	Save(context.Context, string, map[string]interface{}, string) error
	Search(context.Context, string, int32, float32) ([]interface{}, error)
}

type Config struct {
	Provider string
	Url      string
}

type Collection struct {
	Id      string
	Meta    map[string]interface{}
	Context string
	Score   float32
}

type store struct {
	cfg *Config
	st  Store
}

func New(_ context.Context, cfg *Config) Store {
	var st Store

	if cfg.Provider == ProviderChroma {
		st = &Chroma{
			Url: cfg.Url,
		}
	} else if cfg.Provider == ProviderRagx {
		st = &Ragx{
			Url: cfg.Url,
		}
	}

	return &store{
		cfg: cfg,
		st:  st,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (s *store) Init(ctx context.Context, name string) error {
	if s.st == nil {
		return errors.New("invalid store\n")
	}

	if err := s.st.Init(ctx, name); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	return nil
}

func (s *store) Deinit(ctx context.Context) error {
	if s.st != nil {
		if err := s.st.Deinit(ctx); err != nil {
			return errors.Wrap(err, "failed to deinit\n")
		}
	}

	return nil
}

func (s *store) Reset(ctx context.Context) error {
	if err := s.st.Reset(ctx); err != nil {
		return errors.Wrap(err, "failed to reset\n")
	}

	return nil
}

func (s *store) Save(ctx context.Context, text string, meta map[string]interface{}, agent string) error {
	if err := s.st.Save(ctx, text, meta, agent); err != nil {
		return errors.Wrap(err, "failed to save\n")
	}

	return nil
}

func (s *store) Search(ctx context.Context, query string, limit int32, threshold float32) ([]interface{}, error) {
	buf, err := s.st.Search(ctx, query, limit, threshold)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search\n")
	}

	return buf, nil
}
