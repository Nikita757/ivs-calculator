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
	var sum1, sum2, res float64
	num_len := len(numbers)

	for i := 0; i < num_len; i++ {
		sum2 = mathfunc.Add(sum2, numbers[i])
	}
	sum2, _ = mathfunc.Divide(sum2, float64(num_len))
	sum2, _ = mathfunc.Power(sum2, 2);
	sum2 = mathfunc.Multiply(float64(num_len), sum2)

	for i := 0; i < num_len; i++ {
		pow, _ := mathfunc.Power(numbers[i], 2)
		sum1 = mathfunc.Add(sum1, pow)
	}
	res = mathfunc.Substract(sum1, sum2)
	res, _ = mathfunc.Divide(res, float64(num_len-1))
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
