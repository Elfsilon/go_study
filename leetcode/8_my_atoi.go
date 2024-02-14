package leetcode

import (
	"math"
)

var signRunes = map[rune]int{
	43: 1,
	45: -1,
}

func MyAtoi(s string) int {
	res := 0
	sign := 1
	signUsed := false

	for _, r := range s {

		if isWhitespace := r == 32; isWhitespace {
			continue
		}

		if val, isSign := signRunes[r]; isSign {
			if signUsed {
				break
			}
			signUsed = true
			sign = val
			continue
		}

		if isDigit := r < 48 || r > 57; isDigit {
			break
		}

		d := int(r - '0')
		res = res*10 + d
	}

	res *= sign

	if res > math.MaxInt32 {
		res = math.MaxInt32
	} else if res < math.MinInt32 {
		res = math.MinInt32
	}

	return int(int32(res))
}
