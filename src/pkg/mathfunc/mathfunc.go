package mathfunc

import (
	"errors"
)

/**
 * Add: adds two 64-bit floats
 * @param a first float
 * @param b second float
 */
func Add(a, b float64) float64 {
	return a + b
}

/**
 * Subtract: subtracts two 64-bit floats
 * @param a first float
 * @param b second float
 */
func Subtract(a, b float64) float64 {
	return a - b
}

/**
 * Multiply: multiplies two 64-bit floats
 * @param a first float
 * @param b second float
 */
func Multiply(a, b float64) float64 {
	return a * b
}

/**
 * Divide: divides two 64-bit floats. Returns error if b is zero.
 * @param a first float
 * @param b second float
 */
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

/**
 * AbsoluteValue: returns absolute value of a 64-bit float
 * @param a float value
 */
func AbsoluteValue(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

/**
 * Modulo: returns the remainder of division by two 64-bit floats. Returns error if b is zero.
 * @param a first float value
 * @param b second float value
 */
func Modulo(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	quotient := int64(a / b)
	remainder := a - float64(quotient)*b
	if (a < 0 && b > 0 || a > 0 && b < 0) && remainder != 0 {
		remainder += b
	}
	return remainder, nil
}

/**
 * Factorial: divides two 64-bit floats.
 * Works only on natural numbers. Decimals are rounded, negative numbers return an error.
 * @param a float value (internally converted to integer)
 */
func Factorial(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate factorial of negative numbers")
	}
	var output uint64 = 1
	var i uint64 = 1
	for i = 1; i <= uint64(a); i++ {
		previous := output
		output *= i
		if previous > output {
			return 0, errors.New("factorial too big")
		}
	}
	return float64(output), nil
}
