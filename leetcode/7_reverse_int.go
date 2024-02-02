package leetcode

import "math"

func ReverseInt(x int) int {
	res := 0

	for x != 0 {
		res = res*10 + x%10
		x /= 10

		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}

	return res

}
