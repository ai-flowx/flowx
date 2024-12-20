package tool

import (
	"context"
)

type LangChain struct{}

func (l *LangChain) Init(_ context.Context) error {
	return nil
}

func (l *LangChain) Deinit(_ context.Context) error {
	return nil
}

func (l *LangChain) Run(ctx context.Context, invokes []*Invoke) error {
	return nil
}
