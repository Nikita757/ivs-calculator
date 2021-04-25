package mathfunc

import (
	"fmt"
	"math"
)

/**
 * Power: returns base raised to the power of exp as a float64 value
 *
 * Only uses natural numbers (including 0) as the exponent. If epx is not an integer, it is floored, negative values of exp return an error.
 * Base can be any float64 value.
 *
 * @param base float value used as the base of the exponentiation
 * @param exp float value used as the exponent (internally converted to integer)
 */
func Power(base float64, exponent float64) (float64, error) {
	exp := int(exponent) // exponent has to be a natural number - including 0
	if exp < 0 {
		return 0, fmt.Errorf("invalid exponent: '%d', has to be >= 0", exp)
	}
	if exp == 0 && base == 0 {
		return 0, fmt.Errorf("0^0 is undefined")
	}

	var res float64 = 1
	for i := 0; i < exp; i++ {
		res *= base
		if math.IsInf(res, 0) {
			return 0, fmt.Errorf("result of %.3f^%d is too big", base, exp)
		}
	}
	return res, nil
}

func Root(x float64, n float64) (float64, error) {
	degree := float64(int(n))

	if degree == 0 {
		return 0, fmt.Errorf("can't calculate 0th root")
	} else if degree < 0 {
		return 0, fmt.Errorf("can't calculate root of a negative degree: %d", int(degree))
	}

	// if degree is even and x < 0
	if x < 0 && int(degree)%2 == 0 {
		return 0, fmt.Errorf("can't calculate root %d of a negative number: %.3f", int(degree), x)
	}

	// handle special cases
	if x == 0 || x == 1 || degree == 1 {
		return x, nil
	}

	res, old, tmpPow := 1.0, 0.0, 0.0
	oneOverDeg, degMinOne := 1/degree, degree-1
	eps := math.Pow10(-10)

	for math.Abs(old-res) > eps {
		old = res
		tmpPow = math.Pow(res, degMinOne)
		res = oneOverDeg * ((degMinOne * res) + (x / tmpPow))
	}

	return res, nil
}
