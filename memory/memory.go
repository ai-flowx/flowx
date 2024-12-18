package memory

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ai-flowx/flowx/store"
)

const (
	searchLimit     = 3
	searchThreshold = 0.35

	typeLongTerm  = "longterm"
	typeShortTerm = "shortterm"
)

type Memory interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Reset(context.Context) error
	Save(context.Context, string, map[string]interface{}, string) error
	Search(context.Context, string, int32, float32) ([]interface{}, error)
}

type Config struct {
	Store store.Store
	Type  string
}

type memory struct {
	cfg   *Config
	long  *LongTerm
	short *ShortTerm
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
	var err error

	if m.cfg.Type == typeLongTerm {
		m.long = &LongTerm{
			Store: m.cfg.Store,
			Type:  m.cfg.Type,
		}
		err = m.long.Init(ctx)
	} else if m.cfg.Type == typeShortTerm {
		m.short = &ShortTerm{
			Store: m.cfg.Store,
			Type:  m.cfg.Type,
		}
		err = m.short.Init(ctx)
	} else {
		err = errors.New("invalid memory type\n")
	}

	return err
}

func (m *memory) Deinit(ctx context.Context) error {
	var err error

	if m.cfg.Type == typeLongTerm {
		if m.long != nil {
			err = m.long.Deinit(ctx)
		}
	} else if m.cfg.Type == typeShortTerm {
		if m.short != nil {
			err = m.short.Deinit(ctx)
		}
	} else {
		err = errors.New("invalid memory type\n")
	}

	return err
}

func (m *memory) Reset(ctx context.Context) error {
	var err error

	if m.cfg.Type == typeLongTerm {
		if m.long != nil {
			err = m.long.Reset(ctx)
		}
	} else if m.cfg.Type == typeShortTerm {
		if m.short != nil {
			err = m.short.Reset(ctx)
		}
	} else {
		err = errors.New("invalid memory type\n")
	}

	return err
}

func (m *memory) Save(ctx context.Context, text string, meta map[string]interface{}, agent string) error {
	var err error

	if m.cfg.Type == typeLongTerm {
		if m.long != nil {
			err = m.long.Save(ctx, text, meta, agent)
		}
	} else if m.cfg.Type == typeShortTerm {
		if m.short != nil {
			err = m.short.Save(ctx, text, meta, agent)
		}
	} else {
		err = errors.New("invalid memory type\n")
	}

	return err
}

func (m *memory) Search(ctx context.Context, query string, limit int32, threshold float32) ([]interface{}, error) {
	var buf []interface{}
	var err error

	if limit <= 0 {
		limit = searchLimit
	}

	if threshold <= 0 {
		threshold = searchThreshold
	}

	if m.cfg.Type == typeLongTerm {
		if m.long != nil {
			buf, err = m.long.Search(ctx, query, limit, threshold)
		}
	} else if m.cfg.Type == typeShortTerm {
		if m.short != nil {
			buf, err = m.short.Search(ctx, query, limit, threshold)
		}
	} else {
		err = errors.New("invalid memory type\n")
	}

	return buf, err
}
