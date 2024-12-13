package store

import (
	"context"

	"github.com/pkg/errors"
)

const (
	ProviderChroma = "chroma"
	ProviderRagx   = "ragx"
)

var (
	storeList = map[string]Store{
		ProviderChroma: &Chroma{},
		ProviderRagx:   &Ragx{},
	}
)

type Store interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Reset(context.Context) error
	Save(context.Context, interface{}, map[string]interface{}, string) error
	Search(context.Context, string, int, float64) ([]interface{}, error)
}

type Config struct {
	Provider string
	Api      string
	Token    string
}

type Collection struct {
	Id   []string
	Meta map[string]interface{}
	Text string
}

type store struct {
	cfg *Config
	st  Store
}

func New(_ context.Context, cfg *Config) Store {
	return &store{
		cfg: cfg,
		st:  storeList[cfg.Provider],
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (s *store) Init(ctx context.Context) error {
	if s.st == nil {
		return errors.New("invalid store\n")
	}

	if err := s.st.Init(ctx); err != nil {
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

func (s *store) Save(ctx context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	if err := s.st.Save(ctx, value, meta, agent); err != nil {
		return errors.Wrap(err, "failed to save\n")
	}

	return nil
}

func (s *store) Search(ctx context.Context, query string, limit int, threshold float64) ([]interface{}, error) {
	buf, err := s.st.Search(ctx, query, limit, threshold)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search\n")
	}

	return buf, nil
}
