package captcha

import (
	"testing"
)

func TestOperator_whenInputIs1_ItShouldBePlus(t *testing.T) {
	placeholder := 0
	c := New(placeholder, placeholder, 1, placeholder)

	o := c.getOperator()

	if o != "+" {
		t.Errorf("expected + but got %v", o)
	}
}

func TestOperator_whenInputIs2_ItShouldBeMinus(t *testing.T) {
	placeholder := 0
	c := New(placeholder, placeholder, 2, placeholder)

	o := c.getOperator()

	if o != "-" {
		t.Errorf("expected - but get %v", o)
	}
}

func TestOperator_whenInputIs3_ItShouldBeDivide(t *testing.T) {
	placeholder := 0
	c := New(placeholder, placeholder, 3, placeholder)

	o := c.getOperator()

	if o != "/" {
		t.Errorf("expected / but got %v", o)
	}
}

func TestLeftOperand_WhenPatternIs1_ItShouldReturnNumber(t *testing.T) {
	placeholder := 0
	leftOperand := 2
	captcha := New(1, leftOperand, placeholder, placeholder)

	l := captcha.getLeftOperand()

	if l != "2" {
		t.Errorf("expected 2 but got %v", l)
	}
}
