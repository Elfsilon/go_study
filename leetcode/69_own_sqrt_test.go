package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	cases := [][2]int{
		{0, 0},
		{1, 1},
		{2, 1},
		{4, 2},
		{8, 2},
		{9, 3},
		{1128, 33},
		{90400, 300},
	}

	issue := Issue69{}
	for _, testCase := range cases {
		t.Run("My sqrt(x)", func(t *testing.T) {
			r := issue.MySqrt(testCase[0])
			assert.Equal(t, testCase[1], r)
		})
	}
}
