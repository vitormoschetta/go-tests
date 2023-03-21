package sample2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var calculator = Calculator{}
var (
	a = 4
	b = 2
)

func TestSum(t *testing.T) {
	// Arrange
	expected := 6

	// Act
	calculator.Sum(a, b)

	// Assert
	assert.Equal(t, expected, calculator.expected)
}

func TestSubtract(t *testing.T) {
	// Arrange
	expected := 2

	// Act
	calculator.Subtract(a, b)

	// Assert
	assert.Equal(t, expected, calculator.expected)
}

func TestMultiply(t *testing.T) {
	// Arrange
	expected := 8

	// Act
	calculator.Multiply(a, b)

	// Assert
	assert.Equal(t, expected, calculator.expected)
}

func TestDivide(t *testing.T) {
	// Arrange
	expected := 2

	// Act
	calculator.Divide(a, b)

	// Assert
	assert.Equal(t, expected, calculator.expected)
}
