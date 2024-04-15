package leetcode

func ClimbStairs(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	return ClimbStairs(n-1) + ClimbStairs(n-2)
}
