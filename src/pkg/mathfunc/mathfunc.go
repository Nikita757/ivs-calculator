package mathfunc

import (
	"errors"
)

func Add(a float64, b float64) float64

func Substract(a float64, b float64) float64

func Multiply(a float64, b float64) float64

func Divide(a float64, b float64) (float64, error) // b != 0

// Advanced math operations

func Modulo(a float64, b float64) (float64, error) // b != 0

func AbsoluteValue(a float64) float64

func Factorial(a float64) (float64, error) // a has to be an integer - could return an error, or round

func Power(base float64, exponent float64) (float64, error) // exponent has to be a natural number (including 0?)

func Root(x float64, n float64) (float64, error)
