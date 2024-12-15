package store

import (
	"context"
)

type Ragx struct {
	Url string
}

func (r *Ragx) Init(_ context.Context, name string) error {
	return nil
}

func (r *Ragx) Deinit(_ context.Context) error {
	return nil
}

func (r *Ragx) Reset(_ context.Context) error {
	return nil
}

func (r *Ragx) Save(_ context.Context, text string, meta map[string]interface{}, agent string) error {
	return nil
}

func (r *Ragx) Search(_ context.Context, query string, limit int32, threshold float32) ([]interface{}, error) {
	return nil, nil
}
