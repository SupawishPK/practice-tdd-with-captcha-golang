package captcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var placeholder = 0

func TestOperator_whenInputIs1_itShouldBePlus(t *testing.T) {
	c := New(placeholder, placeholder, 1, placeholder)

	o := c.getOperator()

	assert.Equal(t, "+", o)
}

func TestOperator_whenInputIs2_itShouldBeMinus(t *testing.T) {
	c := New(placeholder, placeholder, 2, placeholder)

	o := c.getOperator()

	assert.Equal(t, "-", o)
}

func TestOperator_whenInputIs3_itShouldBeDivide(t *testing.T) {
	c := New(placeholder, placeholder, 3, placeholder)

	o := c.getOperator()

	assert.Equal(t, "/", o)
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs1_itShouldReturnString1(t *testing.T) {
	c := New(1, 1, placeholder, placeholder)

	l := c.getLeftOperand()

	assert.Equal(t, "1", l)
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs2_itShouldReturnString2(t *testing.T) {
	captcha := New(1, 2, placeholder, placeholder)

	l := captcha.getLeftOperand()

	assert.Equal(t, "2", l)
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs9_itShouldReturnString9(t *testing.T) {
	captcha := New(1, 9, placeholder, placeholder)

	l := captcha.getLeftOperand()

	assert.Equal(t, "9", l)
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs8_itShouldReturnString8(t *testing.T) {
	c := New(1, 8, placeholder, placeholder)

	l := c.getLeftOperand()

	assert.Equal(t, "8", l)
}

func TestLeftOperand_whenPatternIs2AndLeftOperandIs1_itShouldReturnStringOne(t *testing.T) {
	c := New(2, 1, placeholder, placeholder)

	l := c.getLeftOperand()

	assert.Equal(t, "One", l)
}

func TestLeftOperand_whenPatternIs2AndLeftOperandIs2_itShouldReturnStringTwo(t *testing.T) {
	c := New(2, 2, placeholder, placeholder)

	l := c.getLeftOperand()

	assert.Equal(t, "Two", l)
}

func TestLeftOperand_whenPatternIs2AndLeftOperandIs9_itShouldReturnStringNine(t *testing.T) {
	c := New(2, 9, placeholder, placeholder)

	l := c.getLeftOperand()

	assert.Equal(t, "Nine", l)
}

func TestLeftOperand_whenPatternIs2AndLefOperandIs8_itShouldReturnStringEight(t *testing.T) {
	c := New(2, 8, placeholder, placeholder)

	l := c.getLeftOperand()

	assert.Equal(t, "Eight", l)
}

func TestRightOperand_whenPatternIs1_andLeftOperandIs1_itShouldReturn1(t *testing.T) {
	c := New(1, placeholder, placeholder, 1)

	r := c.RightOperand()

	assert.Equal(t, "One", r)
}
