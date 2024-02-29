package sort

import (
	"fmt"
	v "go_study/visualizer"
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

func QuickSort(arr []int) *v.Snapshotter {
	snapshotter := v.NewSnapshotter()
	iter := 0
	quickSortR(arr, arr, snapshotter, &iter)
	return snapshotter
}

func quickSortR(original []int, arr []int, snapshotter *v.Snapshotter, iter *int) {
	(*iter) += 1
	l := len(arr)

	if l == 2 {
		if arr[0] >= arr[1] {
			swap(arr, 0, 1)
		}
	} else if l > 1 {
		i, j := 0, l-1
		pivotIndex := (l + 1) / 2
		pivot, p := arr[pivotIndex], -1

		for p < 0 {
			for ; arr[i] < pivot; i++ {
				snapshot := v.NewSnapshot((*iter), original, i, j, pivotIndex)
				snapshot.SetDescription(fmt.Sprintf("pivot = %v, p = %v", pivot, p))
				snapshotter.CaptureSnapshot(snapshot)
			}
			for ; arr[j] > pivot; j-- {
				snapshot := v.NewSnapshot((*iter), original, i, j, pivotIndex)
				snapshot.SetDescription(fmt.Sprintf("mid = %v, p = %v", pivot, p))
				snapshotter.CaptureSnapshot(snapshot)
			}

			if i >= j {
				p = j
			} else {
				snapshot := v.NewSnapshot((*iter), original, i, j, pivotIndex)
				snapshot.SetDescription(fmt.Sprintf("mid = %v, p = %v", pivot, p))
				snapshotter.CaptureSnapshot(snapshot)

				swap(arr, i, j)

				snapshot = v.NewSnapshot((*iter), original, i, j, pivotIndex)
				snapshot.SetDescription(fmt.Sprintf("mid = %v, p = %v", pivot, p))
				snapshotter.CaptureSnapshot(snapshot)

				i++
				j--
			}
		}

		quickSortR(original, arr[:p], snapshotter, iter)
		quickSortR(original, arr[p:], snapshotter, iter)
	}
}
