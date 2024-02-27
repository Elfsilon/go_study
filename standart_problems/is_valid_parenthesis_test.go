package standartproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidParenthesis(t *testing.T) {
	testCases := []struct {
		input  string
		result bool
	}{
		{
			input:  "()()()",
			result: true,
		},
		{
			input:  "()()(()",
			result: false,
		},
		{
			input:  "(()(())())",
			result: true,
		},
		{
			input:  "())",
			result: false,
		},
		{
			input:  "((()())",
			result: false,
		},
		{
			input:  "(((a + b) / (b - c) + 2)",
			result: false,
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			res := IsValidParenthesis(tC.input)
			assert.Equal(t, res, tC.result)
		})
	}
}
