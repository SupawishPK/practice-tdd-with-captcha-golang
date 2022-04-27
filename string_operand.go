package captcha

type StringOperand struct {
	value int
}

func (o StringOperand) result() string {
	numbers := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return numbers[o.value-1]
}
