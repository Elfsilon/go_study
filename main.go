package main

import (
	"go_study/concurrency"
)

func main() {
	// sort.VisualizeSort(sort.QuickSort)
	// issue := leetcode.Issue55{}
	// result := issue.CanJump([]int{2, 3, 1, 1, 4}) // true
	// result := issue.CanJump([]int{3, 2, 1, 0, 4}) // false
	// result := issue.CanJump([]int{2, 0}) // true
	// result := issue.CanJump([]int{3, 0, 0, 0}) // true
	// result := issue.CanJump([]int{3, 0, 0, 0, 0}) // false
	// result := issue.CanJump([]int{2, 5, 0, 0}) // true
	// result := issue.CanJump([]int{3, 0, 8, 2, 0, 0, 1}) // true

	concurrency.RunTweets()
}
