package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSD(t *testing.T) {
	numbers := []float64{321, 43, 400, 10, 143, 400, 500, 640, 320, 4052}

	total := StandardDeviation(numbers)
	if total != 1200.121239245806 {
		t.Errorf("Result is incorrect")
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
