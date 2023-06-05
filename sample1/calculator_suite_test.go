package sample1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
	a        int
	b        int
	expected int
}

// This will run before each test in the suite
// Isso será executado antes de cada teste na suíte
func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

// This will run before each test in the suite
// Isso será executado antes de cada teste na suíte
func (suite *CalculatorTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	suite.a = 4
	suite.b = 2
}

func (suite *CalculatorTestSuite) TestSum() {
	// Arrange
	suite.expected = 6

	// Act
	actual := Sum(suite.a, suite.b)

	// Assert
	suite.Equal(suite.expected, actual)
}

func (suite *CalculatorTestSuite) TestSubtract() {
	// Act
	suite.expected = 2

	// Act
	actual := Subtract(suite.a, suite.b)

	// Assert
	suite.Equal(suite.expected, actual)
}

func (suite *CalculatorTestSuite) TestMultiply() {
	// Arrange
	suite.expected = 8

	// Act
	actual := Multiply(suite.a, suite.b)

	// Assert
	suite.Equal(suite.expected, actual)
}

func (suite *CalculatorTestSuite) TestDivide() {
	// Arrange
	suite.expected = 2

	// Act
	actual := Divide(suite.a, suite.b)

	// Assert
	suite.Equal(suite.expected, actual)
}
