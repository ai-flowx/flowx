package memory

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
)

const (
	searchLimit     = 3
	searchThreshold = 0.35
)

type Memory interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Reset(context.Context) error
	Save(context.Context, interface{}, map[string]interface{}, string) error
	Search(context.Context, string, int, float64) ([]interface{}, error)
}

type Storage interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Reset(context.Context) error
	Save(context.Context, interface{}, map[string]interface{}) error
	Search(context.Context, string, int, float64) ([]interface{}, error)
}

type Config struct {
	Addr    string
	Logger  hclog.Logger
	Storage Storage
}

type memory struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Memory {
	return &memory{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (m *memory) Init(_ context.Context) error {
	return nil
}

func (m *memory) Deinit(_ context.Context) error {
	return nil
}

func (m *memory) Reset(_ context.Context) error {
	return nil
}

func (m *memory) Save(ctx context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	if agent != "" {
		meta["agent"] = agent
	}

	if err := m.cfg.Storage.Save(ctx, value, meta); err != nil {
		return errors.Wrap(err, "failed to save\n")
	}

	return nil
}

func (m *memory) Search(ctx context.Context, query string, limit int, threshold float64) ([]interface{}, error) {
	if limit <= 0 {
		limit = searchLimit
	}

	if threshold <= 0 {
		threshold = searchThreshold
	}

	buf, err := m.cfg.Storage.Search(ctx, query, limit, threshold)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save\n")
	}

	return buf, nil
}
