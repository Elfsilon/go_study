package main

import (
	"fmt"
	"go_study/leetcode"
)

func main() {
	// sort.VisualizeSort(sort.QuickSort)
	prefix := leetcode.LongestCommonPrefix([]string{"football", "foogu", "fooooo"})
	fmt.Println(prefix)
}
