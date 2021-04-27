package profiler

import (
	"testing"
	"math/rand"
)

func TestStandardDeviation(t *testing.T) {
	numbers := []float64{321, 43, 400, 10, 143, 400, 500, 640, 320, 4052}
	
	total := StandardDeviation(numbers)
    if total != 1795.3364896617882 {
       t.Errorf("Result is incorrect")
    }
}

func BenchmarkStandardDeviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []float64{
			float64(rand.Int()), float64(rand.Int()), float64(rand.Int()),
			float64(rand.Int()), float64(rand.Int()), float64(rand.Int()),
			float64(rand.Int()), float64(rand.Int()), float64(rand.Int()),
			float64(rand.Int())}
		StandardDeviation(numbers)
    }	
}
