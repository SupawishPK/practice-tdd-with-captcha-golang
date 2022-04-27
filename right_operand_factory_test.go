package captcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenPatternIs1_shouldReturnStringOperand(t *testing.T) {
	placeholder := 0

	o := newRightOperand(1, placeholder)

	assert.IsType(t, StringOperand{}, o)
}
