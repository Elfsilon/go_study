package hacks

// There is the list of numbers, each number has a pair
// but one hasn't. Find this number
//
// Example:
// nums = [1, 2, 3, 2, 4, 1, 3]
// result = 4
//
// Hack: XOR can be used for each number, because
// X1 XOR X1 = 0, so this will filter all the pairs
func FindOneWithoutPair(nums []int) int {
	r := 0
	for _, n := range nums {
		r ^= n
	}

	return r
}
