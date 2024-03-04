package standartproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRLE(t *testing.T) {
	testCases := []struct {
		input  string
		result string
	}{
		{
			input:  "A",
			result: "A",
		},
		{
			input:  "AAA",
			result: "A3",
		},
		{
			input:  "AAABGGGAAFFFXY",
			result: "A3BG3A2F3XY",
		},
		{
			input:  "FFFFFFFFFFFFFFFBBY",
			result: "F15B2Y",
		},
		{
			input:  "",
			result: "",
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			res := RLE(tC.input)
			assert.Equal(t, res, tC.result)
		})
	}
}
