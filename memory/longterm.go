package memory

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ai-flowx/flowx/store"
)

type LongTerm struct {
	Store store.Store
}

func (l *LongTerm) Init(ctx context.Context) error {
	if err := l.Store.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	return nil
}

func (l *LongTerm) Deinit(ctx context.Context) error {
	if err := l.Store.Deinit(ctx); err != nil {
		return errors.Wrap(err, "failed to deinit\n")
	}

	return nil
}

func (l *LongTerm) Reset(ctx context.Context) error {
	if err := l.Store.Reset(ctx); err != nil {
		return errors.Wrap(err, "failed to reset\n")
	}

	return nil
}

func (l *LongTerm) Save(ctx context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	if err := l.Store.Save(ctx, value, meta, agent); err != nil {
		return errors.Wrap(err, "failed to save\n")
	}

	return nil
}

func (l *LongTerm) Search(ctx context.Context, query string, limit int, threshold float64) ([]interface{}, error) {
	buf, err := l.Store.Search(ctx, query, limit, threshold)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search\n")
	}

	return buf, nil
}
