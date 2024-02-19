package leetcode

// Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.
// A subarray is a contiguous part of the array.
//
// Example 1:
// Input: nums = [1,0,1,0,1], goal = 2
// Output: 4
// Explanation: The 4 subarrays are bolded and underlined below:
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
//
// Example 2:
// Input: nums = [0,0,0,0,0], goal = 0
// Output: 15
func NumSubarraysWithSum(nums []int, goal int) int {
	if len(nums) == 0 {
		return 0
	}

	counter := 0

	startSize := goal
	if goal == 0 {
		startSize += 1
	}

	for wSize := startSize; wSize <= len(nums); wSize++ {
		for start := 0; ; start++ {
			end := start + wSize - 1
			if end >= len(nums) {
				break
			}

			sum := 0
			for i := start; i <= end; i++ {
				sum += nums[i]
			}

			if sum == goal {
				counter += 1
			}
		}
	}

	return counter
}
