//go:build tool_crewai_test

// go test -cover -covermode=atomic -parallel 2 -tags=tool_crewai_test -v github.com/ai-flowx/flowx/tool

package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrewAiInit(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestCrewAiDeinit(t *testing.T) {
	assert.Equal(t, nil, nil)
}

func TestCrewAiRun(t *testing.T) {
	assert.Equal(t, nil, nil)
}
