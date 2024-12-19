package tool

import (
	"context"
)

type FlowX struct {
	Type string
}

func (f *FlowX) Init(_ context.Context) error {
	return nil
}

func (f *FlowX) Deinit(_ context.Context) error {
	return nil
}

func (f *FlowX) Run(ctx context.Context) error {
	return nil
}
