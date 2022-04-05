package captcha

import "strconv"

type Captcha struct {
	pattern     int
	operator    int
	leftOperand int
}

func New(pattern int, leftOperand int, operator int, rightOperand int) Captcha {
	return Captcha{
		pattern:     pattern,
		leftOperand: leftOperand,
		operator:    operator,
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
		return "One"
	}
	return strconv.Itoa(c.leftOperand)
}
