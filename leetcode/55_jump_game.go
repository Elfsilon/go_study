package leetcode

// My submission is not super effective, but it's pretty unique among
// all the others
// This solution can be boosted by reverse-traversing the list,
// this will boost time complexity to O(N), while current is something
// between O(N) and O(N^2), that sucks

type Issue55 struct{}

func (issue *Issue55) CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	i := 0
	for {
		// Forward
		for {
			if i == len(nums)-1 {
				return true
			}
			if nums[i] == 0 {
				break
			}
			i++
		}

		// Backward
		i--
		minToJumpOverZero := 2
		for nums[i] < minToJumpOverZero {
			if i == 0 {
				return false
			}
			i--
			minToJumpOverZero++
		}

		maxJump := issue.findMaxJump(nums, i, minToJumpOverZero)
		i += maxJump
	}
}

func (issue *Issue55) findMaxJump(nums []int, i, min int) int {
	for max := nums[i]; max >= min; max-- {
		if i+max < len(nums) {
			return max
		}
	}

	return min
}
