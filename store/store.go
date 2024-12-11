package store

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

type Store interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Reset(context.Context) error
	Save(context.Context, interface{}, map[string]interface{}, string) error
	Search(context.Context, string, int, float64) ([]interface{}, error)
}

type Config struct {
	Addr   string
	Logger hclog.Logger
}

type store struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Store {
	return &store{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (s *store) Init(_ context.Context) error {
	return nil
}

func (s *store) Deinit(_ context.Context) error {
	return nil
}

func (s *store) Reset(_ context.Context) error {
	return nil
}

func (s *store) Save(_ context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	return nil
}

func (s *store) Search(_ context.Context, query string, limit int, threshold float64) ([]interface{}, error) {
	return nil, nil
}
