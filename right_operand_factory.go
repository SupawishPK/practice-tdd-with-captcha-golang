package captcha

func newRightOperand(pattern int, rightOperand int) Operand {
	return StringOperand{rightOperand}
}
