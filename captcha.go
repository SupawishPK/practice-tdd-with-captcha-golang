package captcha

import "strconv"

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

func New(pattern int, leftOperand int, operator int, rightOperand int) Captcha {
	return Captcha{
		pattern:      pattern,
		leftOperand:  leftOperand,
		operator:     operator,
		rightOperand: rightOperand,
	}
}

func (c Captcha) getOperator() string {
	if c.operator == 1 {
		return "+"
	}
	if c.operator == 3 {
		return "/"
	}
	return "-"
}

func (c Captcha) getLeftOperand() string {
	if c.pattern == 2 {
		numbers := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		return numbers[c.leftOperand-1]
	}
	return strconv.Itoa(c.leftOperand)
}

func (c Captcha) RightOperand() string {
	numbers := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return numbers[c.rightOperand-1]
}
