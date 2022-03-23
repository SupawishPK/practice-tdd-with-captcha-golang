package captcha

type Captcha struct {
	operator int
}

func New(pattern int, leftOperand int, operator int, rightOperand int) Captcha {
	return Captcha{
		operator: operator,
	}
}

func (c Captcha) getOperator() string {
	if c.operator == 1 {
		return "+"
	}
	return "-"
}
