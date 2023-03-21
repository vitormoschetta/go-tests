package sample2

type Calculator struct {
	expected int
}

func (c *Calculator) Sum(a int, b int) {
	c.expected = a + b
}

func (c *Calculator) Subtract(a int, b int) {
	c.expected = a - b
}

func (c *Calculator) Multiply(a int, b int) {
	c.expected = a * b
}

func (c *Calculator) Divide(a int, b int) {
	c.expected = a / b
}
