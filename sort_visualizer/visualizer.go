package sortvisualizer

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Visualizable func(arr []int) *Snapshotter

type Visualizer struct {
	speed int
}

func NewVisualizer() *Visualizer {
	return &Visualizer{
		speed: 10,
	}
}

func (v *Visualizer) SetSpeed(value int) {
	v.speed = value
}

func (v *Visualizer) Play(fn Visualizable, input []int) {
	snapshotter := fn(input)
	name := GetFunctionName(fn)

	for _, snapshot := range snapshotter.snapshots {
		ClearTerminal()

		var b strings.Builder

		b.WriteString(fmt.Sprintf("Name: %v\n", ApplyColor(name, VividYellow)))
		b.WriteString(fmt.Sprintf("Iteration: %v\n", ApplyColor(fmt.Sprint(snapshot.iteration), VividMagenta)))
		b.WriteString("[")

		for elemIndex, element := range snapshot.array {
			curr := fmt.Sprint(element)

			if elemIndex != len(snapshot.array)-1 {
				curr = fmt.Sprintf("%v, ", curr)
			}

			for rawColorIndex, snapIndex := range snapshot.indexes {
				if elemIndex == snapIndex {
					colorIndex := rawColorIndex % len(snapshotter.colors)
					curr = ApplyColor(curr, snapshotter.colors[colorIndex])
				}
			}

			b.WriteString(curr)
		}

		b.WriteString("]\n")

		if snapshot.description != "" {
			b.WriteString(fmt.Sprintf("%v\n", snapshot.description))
		}

		fmt.Println(b.String())
		time.Sleep(time.Duration(v.speed) * 50 * time.Millisecond)
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
