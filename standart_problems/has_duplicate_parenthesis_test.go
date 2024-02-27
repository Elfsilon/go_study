package standartproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasDuplicateParenthesis(t *testing.T) {
	testCases := []struct {
		input  string
		result bool
	}{
		{
			input:  "(a+b)",
			result: false,
		},
		{
			input:  "((a+b))",
			result: true,
		},
		{
			input:  "((a+b)*(a+b))",
			result: false,
		},
		{
			input:  "(((a+b)*(a+b)))",
			result: true,
		},
		{
			input:  "((a+b)*((a+b)))",
			result: true,
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			res := HasDuplicateParenthesis(tC.input)
			assert.Equal(t, res, tC.result)
		})
	}
}
