package mathfunc

import (
	"errors"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	AddTestCase(t, 0, 0, 0)
	AddTestCase(t, 1, 1, 2)
	AddTestCase(t, 1, -1, 0)
	AddTestCase(t, 1.23, 1.23, 2.46)
	AddTestCase(t, 1.23, -1.23, 0)
	AddTestCase(t, 1/3, 1/3, 2/3)
	AddTestCase(t, math.Pi, math.Pi, 2*math.Pi)
	AddTestCase(t, math.MaxFloat64, -math.MaxFloat64, 0)
}

func AddTestCase(t *testing.T, inputA float64, inputB float64, expectedOutput float64) {
	output := Add(inputA, inputB)
	if output != expectedOutput {
		t.Errorf("Add(%f, %f) = %f; should be %f", inputA, inputB, output, expectedOutput)
	}
}

func TestSubstract(t *testing.T) {
	SubstractTestCase(t, 0, 0, 0)
	SubstractTestCase(t, 1, 1, 0)
	SubstractTestCase(t, 1, -1, 2)
	SubstractTestCase(t, -1, 1, -2)
	SubstractTestCase(t, -1, -1, 0)

	SubstractTestCase(t, 1.23, 1.23, 0)
	SubstractTestCase(t, 1.23, -1.23, 2.46)
	SubstractTestCase(t, -1.23, 1.23, -2.46)
	SubstractTestCase(t, -1.23, -1.23, 0)

	SubstractTestCase(t, 1/3, -1/3, 2/3)
	SubstractTestCase(t, math.Pi, -math.Pi, 2*math.Pi)
	SubstractTestCase(t, math.MaxFloat64, math.MaxFloat64, 0)
}

func SubstractTestCase(t *testing.T, inputA float64, inputB float64, expectedOutput float64) {
	output := Substract(inputA, inputB)
	if output != expectedOutput {
		t.Errorf("Substract(%f, %f) = %f; should be %f", inputA, inputB, output, expectedOutput)
	}
}

func TestMultiply(t *testing.T) {
	MultiplyTestCaseCombine(t, 0, 0, 0)
	MultiplyTestCaseCombine(t, 1, 1, 1)
	MultiplyTestCaseCombine(t, 1, 5, 5)
	MultiplyTestCaseCombine(t, 5, 5, 25)
}

func MultiplyTestCaseCombine(t *testing.T, inputA float64, inputB float64, expectedOutput float64) {
	MultiplyTestCase(t, inputA, inputB, expectedOutput)
	MultiplyTestCase(t, inputA, -inputB, -expectedOutput)
	MultiplyTestCase(t, -inputA, inputB, -expectedOutput)
	MultiplyTestCase(t, -inputA, -inputB, expectedOutput)
	MultiplyTestCase(t, inputB, inputA, expectedOutput)
}

func MultiplyTestCase(t *testing.T, inputA float64, inputB float64, expectedOutput float64) {
	output := Multiply(inputA, inputB)
	if output-expectedOutput > math.Pow(10, -10) {
		t.Errorf("Multiply(%f,%f) = %f; should be %f", inputA, inputB, output, expectedOutput)
	}
}

func TestDivide(t *testing.T) {
	DivideTestCase(t, 1, 1, 1, nil)
	DivideTestCase(t, 1, 0, 0, errors.New("cannot divide by zero"))
	DivideTestCase(t, 0, 1, 0, nil)

	DivideTestCaseCombine(t, 0, 5, 0)
	DivideTestCaseCombine(t, 1, 5, 0.2)
	DivideTestCaseCombine(t, 2, 5, 0.4)
	DivideTestCaseCombine(t, 3, 5, 0.6)
	DivideTestCaseCombine(t, 4, 5, 0.8)
	DivideTestCaseCombine(t, 5, 5, 1)
	DivideTestCaseCombine(t, 6, 5, 1.2)

	DivideTestCaseCombine(t, 1, 3, 0.3333333333)
}

func DivideTestCaseCombine(t *testing.T, inputA float64, inputB float64, expectedOutput float64) {
	DivideTestCase(t, inputA, inputB, expectedOutput, nil)
	DivideTestCase(t, inputA, -inputB, -expectedOutput, nil)
	DivideTestCase(t, -inputA, inputB, -expectedOutput, nil)
	DivideTestCase(t, -inputA, -inputB, expectedOutput, nil)
}

func DivideTestCase(t *testing.T, inputA float64, inputB float64, expectedOutput float64, expectedError error) {
	output, err := Divide(inputA, inputB)
	// Check 10 decimals
	if output-expectedOutput > math.Pow(10, -10) {
		t.Errorf("Divide(%f,%f) = %f; should be %f", inputA, inputB, output, expectedOutput)
	}
	if (err == nil && expectedError != nil) || (err != nil && expectedError == nil) || (err != nil && expectedError != nil && err.Error() != expectedError.Error()) {
		t.Errorf("Divide(%f,%f) err = %s; should be %s", inputA, inputB, err, expectedError)
	}
}

func TestAbsoluteValue(t *testing.T) {
	AbsoluteValueTestCase(t, 0, 0)
	AbsoluteValueTestCase(t, -1, 1)
	AbsoluteValueTestCase(t, -(2 ^ 63), (2 ^ 63))
	AbsoluteValueTestCase(t, -math.MaxFloat64, math.MaxFloat64)
	AbsoluteValueTestCase(t, -math.Pi, math.Pi)
	AbsoluteValueTestCase(t, -math.Ln2, math.Ln2)
}

func AbsoluteValueTestCase(t *testing.T, input float64, expectedOutput float64) {
	output := AbsoluteValue(input)
	if output != expectedOutput {
		t.Errorf("AbsoluteValue(%f) = %f; should be %f", input, output, expectedOutput)
	}
}

func TestModulo(t *testing.T) {
	ModuloTestCase(t, 1, 1, 0, nil)
	ModuloTestCase(t, 1, 0, 0, errors.New("cannot divide by zero"))
	ModuloTestCase(t, 0, 1, 0, nil)

	ModuloTestCase(t, 0, 5, 0, nil)
	ModuloTestCase(t, 1, 5, 1, nil)
	ModuloTestCase(t, 2, 5, 2, nil)
	ModuloTestCase(t, 3, 5, 3, nil)
	ModuloTestCase(t, 4, 5, 4, nil)
	ModuloTestCase(t, 5, 5, 0, nil)
	ModuloTestCase(t, 6, 5, 1, nil)

	ModuloTestCase(t, -1, 5, 4, nil)
	ModuloTestCase(t, -2, 5, 3, nil)
	ModuloTestCase(t, -3, 5, 2, nil)
	ModuloTestCase(t, -4, 5, 1, nil)
	ModuloTestCase(t, -5, 5, 0, nil)
	ModuloTestCase(t, -6, 5, 4, nil)

	ModuloTestCase(t, 0, -5, 0, nil)
	ModuloTestCase(t, 1, -5, -4, nil)
	ModuloTestCase(t, 2, -5, -3, nil)
	ModuloTestCase(t, 3, -5, -2, nil)
	ModuloTestCase(t, 4, -5, -1, nil)
	ModuloTestCase(t, 5, -5, 0, nil)
	ModuloTestCase(t, 6, -5, -4, nil)

	ModuloTestCase(t, -1, -5, -1, nil)
	ModuloTestCase(t, -2, -5, -2, nil)
	ModuloTestCase(t, -3, -5, -3, nil)
	ModuloTestCase(t, -4, -5, -4, nil)
	ModuloTestCase(t, -5, -5, 0, nil)
	ModuloTestCase(t, -6, -5, -1, nil)
}

func ModuloTestCase(t *testing.T, inputA float64, inputB float64, expectedOutput float64, expectedError error) {
	output, err := Modulo(inputA, inputB)
	if output != expectedOutput {
		t.Errorf("Modulo(%f,%f) = %f; should be %f", inputA, inputB, output, expectedOutput)
	}
	if (err == nil && expectedError != nil) || (err != nil && expectedError == nil) || (err != nil && expectedError != nil && err.Error() != expectedError.Error()) {
		t.Errorf("Modulo(%f,%f) err = %s; should be %s", inputA, inputB, err, expectedError)
	}
}

func TestFactorial(t *testing.T) {
	FactorialTestCase(t, 0, 1, nil)
	FactorialTestCase(t, 1, 1, nil)
	FactorialTestCase(t, 2, 2, nil)
	FactorialTestCase(t, 3, 6, nil)
	FactorialTestCase(t, 4, 24, nil)
	FactorialTestCase(t, 5, 120, nil)
	FactorialTestCase(t, 10, 3628800, nil)

	FactorialTestCase(t, -1, 0, errors.New("cannot calculate factorial of negative numbers"))
	FactorialTestCase(t, 100000, 0, errors.New("factorial too big"))
}

func FactorialTestCase(t *testing.T, input float64, expectedOutput float64, expectedError error) {
	output, err := Factorial(input)
	if output != expectedOutput {
		t.Errorf("Factorial(%f) = %f; should be %f", input, output, expectedOutput)
	}
	if (err == nil && expectedError != nil) || (err != nil && expectedError == nil) || (err != nil && expectedError != nil && err.Error() != expectedError.Error()) {
		t.Errorf("Factorial(%f) err = %s; should be %s", input, err, expectedError)
	}
}

func TestPower(t *testing.T) {
	PowerTestCase(t, 0, 0, 0, errors.New("0^0 is undefined"))
	PowerTestCase(t, 123, -1, 0, errors.New("invalid exponent: '-1', has to be >= 0"))
	PowerTestCase(t, 34.2, 0, 1, nil)
	PowerTestCase(t, 34.2, 1, 34.2, nil)
	PowerTestCase(t, 5, 2, 25, nil)
	PowerTestCase(t, -5, 2, 25, nil)
	PowerTestCase(t, -5, 3, -125, nil)
	PowerTestCase(t, 10, 4, 10000, nil)
	PowerTestCase(t, 10, 4.4, 10000, nil)
	PowerTestCase(t, 10.2, 4.4, 10824.321600, nil)
	PowerTestCase(t, 10, 5, 100000, nil)
	PowerTestCase(t, 25, 8, 152587890625, nil)
	PowerTestCase(t, 525789, 8, 5841064044963377783181066373525779412512931840.000000, nil)
	PowerTestCase(t, 525789, 20157, 0, errors.New("result of 525789.000^20157 is too big"))

}

func PowerTestCase(t *testing.T, base float64, exp float64, expectedOutput float64, expectedError error) {
	output, err := Power(base, exp)
	// Check 10 decimals
	if math.Abs(output-expectedOutput) > math.Pow(10, -10) {
		t.Errorf("Power(%f, %f) = %f; should be %f", base, exp, output, expectedOutput)
	}
	if (err == nil && expectedError != nil) || (err != nil && expectedError == nil) || (err != nil && expectedError != nil && err.Error() != expectedError.Error()) {
		t.Errorf("Power(%f, %f) err = %s; should be %s", base, exp, err, expectedError)
	}
}
