package search

import (
	"fmt"
	"testing"
)

func genBigSlice(length int) []int {
	arr := make([]int, 0, length)

	for i := 0; i <= length; i++ {
		arr = append(arr, i)
	}
	return arr
}

func checkResults(t *testing.T, found int, expected int) {
	if found != expected {
		t.Fatalf("Found index is %d, but %d expected", found, expected)
	}
}

var testArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// --- Linear search ---

func TestLinearEmpty(t *testing.T) {
	found := Linear(10, []int{})
	checkResults(t, found, -1)
}

func TestLinearEasiest(t *testing.T) {
	found := Linear(1, testArray)
	checkResults(t, found, 0)
}

func TestLinearHardest(t *testing.T) {
	found := Linear(6, testArray)
	checkResults(t, found, 5)
}

// --- Binary search ---

func TestBinaryEmpty(t *testing.T) {
	found := Binary(10, []int{})
	checkResults(t, found, -1)
}

func TestBinaryEasiest(t *testing.T) {
	found := Binary(1, testArray)
	checkResults(t, found, 0)
}

func TestBinaryHardest(t *testing.T) {
	found := Binary(6, testArray)
	checkResults(t, found, 5)
}

func TestBinaryIterEmpty(t *testing.T) {
	found := BinaryIter(10, []int{})
	checkResults(t, found, -1)
}

func TestBinaryIterEasiest(t *testing.T) {
	found := BinaryIter(1, testArray)
	checkResults(t, found, 0)
}

func TestBinaryIterHardest(t *testing.T) {
	found := BinaryIter(6, testArray)
	checkResults(t, found, 5)
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

func runBenchmarksWith(b *testing.B, fn SearchFn[int]) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("input_size_%d", size.input), func(b *testing.B) {
			arr := genBigSlice(size.input)
			target := len(arr)

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

func BenchmarkBinaryIter(b *testing.B) {
	runBenchmarksWith(b, BinaryIter)
}
