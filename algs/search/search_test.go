package search

import (
	"fmt"
	"testing"
)

func genBigSlice(length int) []string {
	arr := make([]string, 0, length)

	for i := 0; i <= length; i++ {
		arr = append(arr, fmt.Sprintf("Element %d", i))
	}
	return arr
}

func checkResults(t *testing.T, err error, found int, expected int) {
	if err != nil {
		t.Fatalf("Got an error: %s", err)
	}

	if found != expected {
		t.Fatalf("Found index is %d, but %d expected", found, expected)
	}
}

var testArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// --- Linear search ---

func TestLinearEmpty(t *testing.T) {
	found, err := Linear(10, []int{})
	checkResults(t, err, found, -1)
}

func TestLinearEasiest(t *testing.T) {
	found, err := Linear(1, testArray)
	checkResults(t, err, found, 0)
}

func TestLinearHardest(t *testing.T) {
	found, err := Linear(6, testArray)
	checkResults(t, err, found, 5)
}

// --- Binary search ---

func TestBinaryEmpty(t *testing.T) {
	found, err := Binary(10, []int{})
	checkResults(t, err, found, -1)
}

func TestBinaryEasiest(t *testing.T) {
	found, err := Binary(1, testArray)
	checkResults(t, err, found, 0)
}

func TestBinaryHardest(t *testing.T) {
	found, err := Binary(6, testArray)
	checkResults(t, err, found, 5)
}

func TestBinaryExceptional(t *testing.T) {
	var found int
	var err error

	found, err = Binary(0, []int{1, 2})
	checkResults(t, err, found, -1)

	found, err = Binary(3, []int{1, 2})
	checkResults(t, err, found, -1)

	found, err = Binary(1, []int{1})
	checkResults(t, err, found, 0)
}

// --- Benchmarks ---

var sizes = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
}

func runBenchmarksWith(b *testing.B, fn SearchFn[string]) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("input_size_%d", size.input), func(b *testing.B) {
			arr := genBigSlice(size.input)
			target := fmt.Sprintf("Element %d", len(arr))

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				fn(target, arr)
			}
		})
	}
}

func BenchmarkLinear(b *testing.B) {
	runBenchmarksWith(b, Linear)
}

func BenchmarkBinary(b *testing.B) {
	runBenchmarksWith(b, Binary)
}
