package leetcode

import "strings"

// My submission beats 100% of solutions by speed and 78% by memory!!!!
func LongestCommonPrefix(strs []string) string {
	var b strings.Builder

	approved := true
	for i := 0; approved; i++ {
		var suspicion byte

		for j := 0; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				approved = false
				break
			}

			if suspicion == 0 {
				suspicion = strs[j][i]
			} else if strs[j][i] != suspicion {
				approved = false
				break
			}
		}

		if approved {
			b.WriteByte(suspicion)
		}
	}

	return b.String()
}
