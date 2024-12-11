package store

import (
	"context"
)

type Ragx struct {
	Collection []Collection
}

type Collection struct {
	Id   []string
	Meta map[string]interface{}
	Text string
}

func (r *Ragx) Init(_ context.Context) error {
	return nil
}

func (r *Ragx) Deinit(_ context.Context) error {
	return nil
}

func (r *Ragx) Reset(_ context.Context) error {
	return nil
}

func (r *Ragx) Save(_ context.Context, value interface{}, meta map[string]interface{}, agent string) error {
	return nil
}

func (r *Ragx) Search(_ context.Context, query string, limit int, threshold float64) ([]interface{}, error) {
	return nil, nil
}

func (r *Ragx) embedding(_ context.Context, text string, meta map[string]interface{}) ([]Collection, error) {
	return nil, nil
}
