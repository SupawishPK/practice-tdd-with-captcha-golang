package captcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenStringOperandContain1_resultShoudlReturnOne(t *testing.T) {
	so := StringOperand{
		value: 1,
	}

	r := so.result()

	assert.Equal(t, "One", r)
}

func Test_whenStringOperandContain2_resultShoudlReturnTwo(t *testing.T) {
	so := StringOperand{
		value: 2,
	}

	r := so.result()

	assert.Equal(t, "Two", r)
}

func Test_whenStringOperandContain9_resultShoudlReturnNine(t *testing.T) {
	so := StringOperand{
		value: 9,
	}

	r := so.result()

	assert.Equal(t, "Nine", r)
}
