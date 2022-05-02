package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSD(t *testing.T) {
	numbers := []float64{321, 43, 400, 10, 143, 400, 500, 640, 320, 4052}
	expectedRes := 1200.121239245806

	total := StandardDeviation(numbers)
	if total != expectedRes {
		t.Errorf("Result is: %.10f, should be: %.10f", total, expectedRes)
	}

	numbers = []float64{10, 20}
	expectedRes = 7.0710678119

	total = StandardDeviation(numbers)
	if math.Abs(total-expectedRes) > math.Pow(10, -10) {
		t.Errorf("Result is: %.10f, should be: %.10f", total, expectedRes)
	}

	numbers = []float64{987, 382, 928, 278, 273, 421, 515, 832, 286, 495}
	expectedRes = 275.2453652854

	total = StandardDeviation(numbers)
	if math.Abs(total-expectedRes) > math.Pow(10, -10) {
		t.Errorf("Result is: %.10f, should be: %.10f", total, expectedRes)
	}

	numbers = []float64{987.2, 382.5, 928, 278, 273, 421, 515, 832, 286, 495}
	expectedRes = 275.2496969420

	total = StandardDeviation(numbers)
	if math.Abs(total-expectedRes) > math.Pow(10, -10) {
		t.Errorf("Result is: %.10f, should be: %.10f", total, expectedRes)
	}
}

func benchmarkSD(file *os.File, b *testing.B) {
	reader := bufio.NewReader(file)
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

	for i := 0; i < b.N; i++ {
		StandardDeviation(f_numbers)
	}
}

func BenchmarkSD10(b *testing.B) {
	f, err := os.Open("./testdata10")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	benchmarkSD(f, b)

}

func BenchmarkSD100(b *testing.B) {
	f, err := os.Open("./testdata100")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	benchmarkSD(f, b)
}

func BenchmarkSD1000(b *testing.B) {
	f, err := os.Open("./testdata1000")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	benchmarkSD(f, b)
}
