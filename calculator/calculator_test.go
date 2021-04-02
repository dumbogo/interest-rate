package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) { // TODO: WIP
	res := Calculate(Config{})
	assert.Equal(t, uint32(0), res, "True is true")
}
