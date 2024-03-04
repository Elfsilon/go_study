package standartproblems

import (
	"fmt"
	"strings"
)

func RLE(source string) string {
	if source == "" {
		return ""
	}

	n := len(source)
	var b strings.Builder

	i := 0
	for i < n {
		count := 1
		for i < n-1 && source[i] == source[i+1] {
			count++
			i++
		}

		b.WriteByte(source[i])
		if count > 1 {
			b.WriteString(fmt.Sprint(count))
		}
		i++
	}

	return b.String()
}
