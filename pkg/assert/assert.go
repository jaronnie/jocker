package assert

import "github.com/stretchr/testify/assert"

type JockerAssert struct{}

func (j *JockerAssert) Errorf(format string, args ...interface{}) {}

func NewJockerAssert() *assert.Assertions {
	return assert.New(&JockerAssert{})
}
