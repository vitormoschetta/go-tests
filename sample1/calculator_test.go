package sample1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a        = 4
	b        = 2
	expected int
)

func TestSum(t *testing.T) {
	// Arrange
	expected = 6

	// Act
	actual := Sum(a, b)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestSubtract(t *testing.T) {
	// Arrange
	expected = 2

	// Act
	actual := Subtract(a, b)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestMultiply(t *testing.T) {
	// Arrange
	expected = 8

	// Act
	actual := Multiply(a, b)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestDivide(t *testing.T) {
	// Arrange
	expected = 2

	// Act
	actual := Divide(a, b)

	// Assert
	assert.Equal(t, expected, actual)
}
