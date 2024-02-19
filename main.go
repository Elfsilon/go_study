package main

import (
	srtv "go_study/sort_visualizer"
	"math/rand"
)

type BubbleSort struct{}

func (b *BubbleSort) Iterate(arr []int) *srtv.Snapshotter {
	snapshotter := srtv.NewSnapshotter()

	for i := 0; i < len(arr)-1; i++ {

		for j := i + 1; j < len(arr); j++ {
			snapshotter.CaptureSnapshot(arr, srtv.NewSingleSwap(i, j))

			if arr[i] > arr[j] {
				swap(arr, i, j)
				snapshotter.CaptureSnapshot(arr, srtv.NewSingleSwap(i, j))
			}
		}
	}

	return snapshotter
}

func (b *BubbleSort) Name() string {
	return "Bubble Sort"
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	// fmt.Println(leetcode.BuddyStrings("aghthb", "agbthh")) // True
	// fmt.Println(leetcode.BuddyStrings("aghthb", "agfthg")) // False
	// fmt.Println(leetcode.BuddyStrings("aghthb", "agfjhg")) // False
	// fmt.Println(leetcode.BuddyStrings("agat", "agat"))     // True
	// fmt.Println(leetcode.BuddyStrings("agbt", "agbt"))     // False
	// fmt.Println(leetcode.BuddyStrings("abab", "baba"))     // Must be false! - duplicates is not the antidote

	bubble := BubbleSort{}
	visualizer := srtv.NewVisualizer(&bubble)

	input := make([]int, 0)
	for i := 16; i > 0; i -= 1 {
		input = append(input, rand.Intn(100))
	}

	visualizer.Play(input)
}
