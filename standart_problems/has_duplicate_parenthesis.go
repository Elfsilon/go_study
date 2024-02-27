package standartproblems

import (
	"go_study/structures/stack"
)

// Suppose that [expression] is valid
func HasDuplicateParenthesis(expression string) bool {
	s := stack.NewStack[string]()

	for _, r := range expression {
		char := string(r)

		if char == closing_parenthesis {
			top := ""
			taken := 0

			for top != opening_parenthesis {
				top, _ = s.Peek()
				s.Pop()
				taken++
			}

			// If we took only "(" parenthesis
			if taken == 1 {
				return true
			}
		} else {
			s.Push(char)
		}
	}

	return false
}
