package leetcode

// Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.
// Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].
// For example, swapping at indices 0 and 2 in "abcd" results in "cbad".
//
// Input: s = "ab", goal = "ba"
// Output: true
func BuddyStrings(s string, goal string) bool {
	if len(s) == 1 || len(goal) == 1 || len(s) != len(goal) {
		return false
	}

	seen := make(map[byte]bool)
	diff := make(map[byte]byte)

	var k1 byte
	hasDuplicates := false

	for i := 0; i < len(s); i += 1 {
		if _, ok := seen[s[i]]; ok {
			hasDuplicates = true
		} else {
			seen[s[i]] = true
		}

		if s[i] != goal[i] {
			diff[s[i]] = goal[i]
			k1 = s[i]
		}
	}

	if len(diff) == 0 {
		return hasDuplicates
	} else if len(diff) != 2 {
		return false
	}

	v1, ok := diff[k1]
	if !ok {
		return false
	}

	v2, ok := diff[v1]
	if !ok {
		return false
	}

	if k1 == v2 {
		return true
	}

	return false
}
