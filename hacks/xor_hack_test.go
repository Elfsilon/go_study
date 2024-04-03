package hacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase[Input any, Result any] struct {
	input  Input
	expect Result
}

type TestCases[Input any, Result any] []TestCase[Input, Result]

func TestFindOneWithoutPair(t *testing.T) {
	cases := TestCases[[]int, int]{
		{
			input:  []int{1, 1, 2},
			expect: 2,
		},
		{
			input:  []int{1, 2, 3, 3, 1, 2, 1, 4, 1},
			expect: 4,
		},
		{
			input:  []int{1, 5, 2, 3, 4, 1, 4, 3, 5},
			expect: 2,
		},
	}

	for _, testCase := range cases {
		t.Run("TestFindOneWithoutPair", func(t *testing.T) {
			r := FindOneWithoutPair(testCase.input)
			assert.Equal(t, testCase.expect, r)
		})
	}
}
