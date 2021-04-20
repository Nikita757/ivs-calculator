package mathfunc

import (
	"fmt"
	"math"
)

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
			return 0, fmt.Errorf("result of %f^%d is too big", base, exp)
		}
	}
	return res, nil
}

func Root(x float64, n float64) (float64, error) {
	return 0, nil
}
