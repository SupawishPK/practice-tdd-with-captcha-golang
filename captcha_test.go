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
	captcha := New(1, 1, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "1" {
		t.Errorf("expected 1 but got %v", l)
	}
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs2_itShouldReturnString2(t *testing.T) {
	captcha := New(1, 2, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "2" {
		t.Errorf("expected 2 but got %v", l)
	}
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs9_itShouldReturnString9(t *testing.T) {
	captcha := New(1, 9, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "9" {
		t.Errorf("expected 9 but got %v", l)
	}
}

func TestLeftOperand_whenPatternIs1AndLeftOperandIs8_itShouldReturnString8(t *testing.T) {
	captcha := New(1, 8, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "8" {
		t.Errorf("expected 8 but got %v", l)
	}
}

func TestLeftOperand_whenPatternIs2AndLeftOperandIs1_itShouldReturnStringOne(t *testing.T) {
	captcha := New(2, 1, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "One" {
		t.Errorf("expected One but got %v", l)
	}
}
