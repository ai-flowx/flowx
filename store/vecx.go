package store

import (
	"context"
)

type Vecx struct{}

func (v *Vecx) Init(_ context.Context, name string) error {
	return nil
}

func (v *Vecx) Deinit(_ context.Context) error {
	return nil
}

func (v *Vecx) Reset(_ context.Context) error {
	return nil
}

func (v *Vecx) Save(_ context.Context, text string, meta map[string]interface{}, agent string) error {
	return nil
}

func (v *Vecx) Search(_ context.Context, query string, limit int32, threshold float32) ([]interface{}, error) {
	return nil, nil
}
