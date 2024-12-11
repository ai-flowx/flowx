package memory

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"

	"github.com/ai-flowx/flowx/store"
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

type Config struct {
	Addr   string
	Logger hclog.Logger
	Store  store.Store
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

func (m *memory) Init(ctx context.Context) error {
	if err := m.cfg.Store.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	return nil
}

func (m *memory) Deinit(ctx context.Context) error {
	if err := m.cfg.Store.Deinit(ctx); err != nil {
		return errors.Wrap(err, "failed to deinit\n")
	}

	return nil
}

func (m *memory) Reset(ctx context.Context) error {
	if err := m.cfg.Store.Reset(ctx); err != nil {
		return errors.Wrap(err, "failed to reset\n")
	}

	return nil
}

func (m *memory) Save(ctx context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	if err := m.cfg.Store.Save(ctx, value, meta, agent); err != nil {
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

	buf, err := m.cfg.Store.Search(ctx, query, limit, threshold)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search\n")
	}

	return buf, nil
}
