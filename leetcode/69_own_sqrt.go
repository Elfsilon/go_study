package leetcode

type Issue69 struct{}

func (i *Issue69) MySqrt(x int) int {
	if x <= 1 {
		return x
	}

	res := 0
	start, end := 1, x/2

	for start <= end {
		mid := (start + end) / 2
		msqr := mid * mid

		// Perfect square
		if msqr == x {
			return mid
		}

		if msqr < x {
			// Try to find a closer value in the next iter
			res, start = mid, mid+1
		} else {
			end = mid - 1
		}
	}

	return res
}
