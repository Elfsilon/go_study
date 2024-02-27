package standartproblems

import "go_study/structures/stack"

func IsValidParenthesis(expression string) bool {
	s := stack.NewStack[string]()

	for _, r := range expression {
		char := string(r)

		if char == opening_parenthesis {
			s.Push(char)
		} else if char == closing_parenthesis {
			if s.IsEmpty() {
				return false
			}
			s.Pop()
		}
	}

	return s.IsEmpty()
}
