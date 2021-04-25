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
	return 0, nil
}
