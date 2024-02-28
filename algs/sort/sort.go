package sort

import (
	"fmt"
	v "go_study/sort_visualizer"
)

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func BubbleSort(arr []int) *v.Snapshotter {
	snapshotter := v.NewSnapshotter()

	sorted := false
	for i := 0; !sorted && i < len(arr)-1; i++ {
		sorted = true

		for j := i + 1; j < len(arr); j++ {
			snapshotter.CaptureSnapshot(v.NewSnapshot(i, arr, i, j))

			if arr[i] > arr[j] {
				swap(arr, i, j)
				sorted = false

				snapshotter.CaptureSnapshot(v.NewSnapshot(i, arr, i, j))
			}
		}
	}

	return snapshotter
}

func SelectionSort(arr []int) *v.Snapshotter {
	snapshotter := v.NewSnapshotter()

	l := 0
	for l < len(arr)-1 {
		minIndex := l
		min := arr[minIndex]

		var desc string
		for i := l + 1; i < len(arr); i++ {
			desc = fmt.Sprintf("Min: %v", min)

			snapshot := v.NewSnapshot(l, arr, l, i)
			snapshot.SetDescription(desc)
			snapshotter.CaptureSnapshot(snapshot)

			if arr[i] < min {
				min, minIndex = arr[i], i
			}
		}

		snapshot := v.NewSnapshot(l, arr, l, minIndex)
		snapshot.SetDescription(desc)
		snapshotter.CaptureSnapshot(snapshot)

		swap(arr, l, minIndex)

		snapshot = v.NewSnapshot(l, arr, l, minIndex)
		snapshot.SetDescription(desc)
		snapshotter.CaptureSnapshot(snapshot)

		l++
	}

	return snapshotter
}

func InsertionSort(arr []int) *v.Snapshotter {
	snapshotter := v.NewSnapshotter()

	for i := 2; i < len(arr); i++ {
		iteration := i - 2
		snapshotter.CaptureSnapshot(v.NewSnapshot(iteration, arr, i, i-1))

		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			snapshotter.CaptureSnapshot(v.NewSnapshot(iteration, arr, j, j+1, i))
			swap(arr, j, j+1)
			snapshotter.CaptureSnapshot(v.NewSnapshot(iteration, arr, j, j+1, i))
		}
	}

	return snapshotter
}
