package sample2

import (
	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
	calculator *Calculator
	a          int
	b          int
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calculator = &Calculator{}
	suite.a = 4
	suite.b = 2
}

func (suite *CalculatorTestSuite) TestSum() {
	// Arrange
	expected := 6

	// Act
	suite.calculator.Sum(suite.a, suite.b)

	// Assert
	suite.Equal(expected, suite.calculator.expected)
}

func (suite *CalculatorTestSuite) TestSubtract() {
	// Arrange
	expected := 2

	// Act
	suite.calculator.Subtract(suite.a, suite.b)

	// Assert
	suite.Equal(expected, suite.calculator.expected)
}

func (suite *CalculatorTestSuite) TestMultiply() {
	// Arrange
	expected := 8

	// Act
	suite.calculator.Multiply(suite.a, suite.b)

	// Assert
	suite.Equal(expected, suite.calculator.expected)
}

func (suite *CalculatorTestSuite) TestDivide() {
	// Arrange
	expected := 2

	// Act
	suite.calculator.Divide(suite.a, suite.b)

	// Assert
	suite.Equal(expected, suite.calculator.expected)
}
