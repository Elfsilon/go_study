package leetcode

var roman map[rune]int = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	runes := []rune(s)

	i := 0
	sum := 0

	for i < len(runes) {
		curr := runes[i]

		cond := i == len(runes)-1 || (curr != 'I' || runes[i+1] != 'V' && runes[i+1] != 'X') &&
			(curr != 'X' || runes[i+1] != 'L' && runes[i+1] != 'C') &&
			(curr != 'C' || runes[i+1] != 'D' && runes[i+1] != 'M')

		if cond {
			sum += roman[curr]
			i++
		} else {
			sum += roman[runes[i+1]] - roman[runes[i]]
			i += 2
		}
	}

	return sum
}
