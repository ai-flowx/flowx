package agent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initParserTest(ctx context.Context) parser {
	return parser{}
}

func TestParse(t *testing.T) {
	ctx := context.Background()
	p := initParserTest(ctx)

	content := ""
	_, err := p.parse(ctx, content)
	assert.Equal(t, nil, err)
}
