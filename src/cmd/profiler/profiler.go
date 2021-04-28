package main

import (
	"bufio"
	"fmt"
	"io"
	"ivs-calculator/pkg/mathfunc"
	"os"
	"strconv"
	"strings"
)

func StandardDeviation(numbers []float64) float64 {
	var sum, mean, res float64
	len := len(numbers)

	for i := 0; i < len; i++ {
		sum = mathfunc.Add(sum, numbers[i])
	}
	mean, _ = mathfunc.Divide(sum, float64(len))

	for i := 0; i < len; i++ {
		pow1, _ := mathfunc.Power(numbers[i], 2)
		pow2, _ := mathfunc.Power(mean, 2)
		res = mathfunc.Add(res, mathfunc.Substract(pow1, mathfunc.Multiply(float64(len), pow2)))
	}

	div, _ := mathfunc.Divide(1, mathfunc.Substract(1, float64(len)))
	res = mathfunc.Multiply(res, div)
	res, _ = mathfunc.Root(res, 2)

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text := ""
	for {
		str, err := reader.ReadString('\n')
		text += str
		if err == io.EOF {
			break
		}
	}
	s_numbers := strings.Fields(text)
	f_numbers := make([]float64, len(s_numbers))
	for i := 0; i < len(f_numbers); i++ {
		f_numbers[i], _ = strconv.ParseFloat(s_numbers[i], 64)
	}

	fmt.Println(StandardDeviation(f_numbers))
}
