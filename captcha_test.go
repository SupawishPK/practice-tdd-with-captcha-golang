package captcha

import (
	"testing"
)

var placeholder = 0

func TestOperator_whenInputIs1_itShouldBePlus(t *testing.T) {
	c := New(placeholder, placeholder, 1, placeholder)

	o := c.getOperator()

	if o != "+" {
		t.Errorf("expected + but got %v", o)
	}
}

func TestOperator_whenInputIs2_itShouldBeMinus(t *testing.T) {
	c := New(placeholder, placeholder, 2, placeholder)

	o := c.getOperator()

	if o != "-" {
		t.Errorf("expected - but get %v", o)
	}
}

func TestOperator_whenInputIs3_itShouldBeDivide(t *testing.T) {
	c := New(placeholder, placeholder, 3, placeholder)

	o := c.getOperator()

	if o != "/" {
		t.Errorf("expected / but got %v", o)
	}
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs1_itShouldReturnString1(t *testing.T) {
	leftOperand := 1
	captcha := New(1, leftOperand, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "1" {
		t.Errorf("expected 1 but got %v", l)
	}
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs2_itShouldReturnString2(t *testing.T) {
	leftOperand := 2
	captcha := New(1, leftOperand, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "2" {
		t.Errorf("expected 2 but got %v", l)
	}
}
