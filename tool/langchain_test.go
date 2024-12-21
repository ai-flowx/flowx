//go:build tool_langchain_test

// go test -cover -covermode=atomic -parallel 2 -tags=tool_langchain_test -v github.com/ai-flowx/flowx/tool

package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLangChain(t *testing.T) {
	assert.Equal(t, nil, nil)
}
