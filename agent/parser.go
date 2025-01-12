package agent

import (
	"context"
)

type parser struct{}

func (p *parser) parse(ctx context.Context, content string) (Action, error) {
	return Action{}, nil
}
