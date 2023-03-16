package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJockerAssert(t *testing.T) {
	jassert := NewJockerAssert()
	assert.True(t, jassert.Equal(1, 1))
	assert.True(t, jassert.Greater(2, 1))
}
