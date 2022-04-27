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
		so := StringOperand{value: c.leftOperand}
		return so.result()
	}
	return strconv.Itoa(c.leftOperand)
}

func (c Captcha) RightOperand() string {
	if c.pattern == 2 {
		return strconv.Itoa(c.rightOperand)
	}
	so := StringOperand{value: c.rightOperand}
	return so.result()
}
