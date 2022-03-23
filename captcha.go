package captcha

type Captcha struct {
	operator string
}

func New(pattern int, leftOperand int, operator int, rightOperand int) Captcha {

	return Captcha{
		operator: convertNumberToOperator(operator),
	}
}

func (c Captcha) getOperator() string {
	return c.operator
}

func convertNumberToOperator(number int) string {
	if number == 1 {
		return "+"
	}
	return "-"
}
