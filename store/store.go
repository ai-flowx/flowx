package store

import (
	"context"

	"github.com/hashicorp/go-hclog"
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
	Logger   hclog.Logger
	Provider string
}

type Collection struct {
	Id   []string
	Meta map[string]interface{}
	Text string
}

type store struct {
	cfg    *Config
	_store Store
}

func New(_ context.Context, cfg *Config) Store {
	return &store{
		cfg:    cfg,
		_store: storeList[cfg.Provider],
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (s *store) Init(ctx context.Context) error {
	if s._store == nil {
		return errors.New("invalid store\n")
	}

	if err := s._store.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	return nil
}

func (s *store) Deinit(ctx context.Context) error {
	if s._store != nil {
		if err := s._store.Deinit(ctx); err != nil {
			return errors.Wrap(err, "failed to deinit\n")
		}
	}

	return nil
}

func (s *store) Reset(ctx context.Context) error {
	if err := s._store.Reset(ctx); err != nil {
		return errors.Wrap(err, "failed to reset\n")
	}

	return nil
}

func (s *store) Save(ctx context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	if err := s._store.Save(ctx, value, meta, agent); err != nil {
		return errors.Wrap(err, "failed to save\n")
	}

	return nil
}

func (s *store) Search(ctx context.Context, query string, limit int, threshold float64) ([]interface{}, error) {
	buf, err := s._store.Search(ctx, query, limit, threshold)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search\n")
	}

	return buf, nil
}
