package sortvisualizer

import (
	"fmt"
	"strings"
	"time"
)

type Visualizable interface {
	Name() string
	Iterate(arr []int) *Snapshotter
}

type Snapshotter struct {
	snapshots []Snapshot
}

func NewSnapshotter() *Snapshotter {
	return &Snapshotter{
		snapshots: make([]Snapshot, 0),
	}
}

func (s *Snapshotter) CaptureSnapshot(array []int, swaps []Swap) {
	arrCopy := make([]int, len(array), cap(array))
	copy(arrCopy, array)
	s.snapshots = append(s.snapshots, NewSnapshot(arrCopy, swaps))
}

type Snapshot struct {
	array []int
	swaps []Swap
}

func NewSnapshot(array []int, swaps []Swap) Snapshot {
	return Snapshot{
		array: array,
		swaps: swaps,
	}
}

type Swap struct {
	i int
	j int
}

func NewSwap(i, j int) Swap {
	return Swap{
		i: i,
		j: j,
	}
}

func NewSingleSwap(i, j int) []Swap {
	return []Swap{
		{
			i: i,
			j: j,
		},
	}
}

var colors = map[bool]string{
	true:  VividCyan,
	false: VividBlue,
}

type Visualizer struct {
	fn    Visualizable
	speed int
}

func NewVisualizer(fn Visualizable) *Visualizer {
	return &Visualizer{
		fn:    fn,
		speed: 10,
	}
}

func (v *Visualizer) SetFn(fn Visualizable) {
	v.fn = fn
}

func (v *Visualizer) Play(input []int) {
	snapshotter := v.fn.Iterate(input)

	for i, snapshot := range snapshotter.snapshots {
		ClearTerminal()

		var b strings.Builder

		b.WriteString(fmt.Sprintf("Name: %v\n", ApplyColor(v.fn.Name(), VividYellow)))
		b.WriteString(fmt.Sprintf("Iteration: %v\n", ApplyColor(fmt.Sprint(i), VividMagenta)))

		swaps := make(map[int]bool)
		for _, p := range snapshot.swaps {
			swaps[p.i] = true
			swaps[p.j] = false
		}

		b.WriteString("[")

		for elemIndex, element := range snapshot.array {
			curr := fmt.Sprint(element)

			if elemIndex != len(snapshot.array)-1 {
				curr = fmt.Sprintf("%v, ", curr)
			}

			if colorValue, ok := swaps[elemIndex]; ok {
				curr = ApplyColor(curr, colors[colorValue])
			}

			b.WriteString(curr)
		}
		b.WriteString("]\n")

		fmt.Println(b.String())

		time.Sleep(time.Duration(v.speed) * 50 * time.Millisecond)
	}
}
