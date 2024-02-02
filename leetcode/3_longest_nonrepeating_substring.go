package leetcode

import "go_study/structures"

// Given a string s, find the length of the longest substring
// without repeating characters
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3

// My solution
func LengthOfLongestSubstring(s string) int {
	lst := structures.NewLinkedList[byte]()
	max := 0

	for i := range s {
		el := s[i]

		if lst.Contains(el) {
			lst.RemoveFirst(el)
		}

		lst.Add(el)

		if lst.Count > max {
			max = lst.Count
		}
	}

	return max
}

// Best solution
func LengthOfLongestSubstringBest(s string) int {
	start := 0
	longest := 0
	used := map[byte]int{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if _, ok := used[c]; ok && used[c] >= start {
			start = used[c] + 1
		}

		longest = max(longest, i-start+1)
		used[c] = i
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
