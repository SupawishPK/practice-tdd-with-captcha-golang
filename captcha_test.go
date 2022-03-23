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
