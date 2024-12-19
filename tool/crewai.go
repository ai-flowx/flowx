package tool

import (
	"context"
)

type CrewAi struct {
	Type string
}

func (c *CrewAi) Init(_ context.Context) error {
	return nil
}

func (c *CrewAi) Deinit(_ context.Context) error {
	return nil
}

func (c *CrewAi) Run(ctx context.Context, invokes []*Invoke) error {
	return nil
}
