package sort

import (
	v "go_study/sort_visualizer"
	"math/rand"
)

func VisualizeSort(fn v.Visualizable) {
	input := make([]int, 0)
	for i := 16; i > 0; i -= 1 {
		input = append(input, rand.Intn(100))
	}

	visualizer := v.NewVisualizer()
	visualizer.Play(fn, input)
}
