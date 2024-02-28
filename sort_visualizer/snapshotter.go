package sortvisualizer

type Snapshotter struct {
	snapshots []Snapshot
	colors    []string
}

func NewSnapshotter() *Snapshotter {
	return &Snapshotter{
		snapshots: make([]Snapshot, 0),
		colors: []string{
			VividBlue,
			VividCyan,
			VividGreen,
			VividYellow,
			VividRed,
			VividMagenta,
		},
	}
}

func (s *Snapshotter) CaptureSnapshot(snapshot *Snapshot) {
	s.snapshots = append(s.snapshots, *snapshot)
}

type Snapshot struct {
	array       []int
	indexes     []int
	iteration   int
	description string
}

func NewSnapshot(iteration int, array []int, indexes ...int) *Snapshot {
	arrCopy := make([]int, len(array), cap(array))
	copy(arrCopy, array)

	return &Snapshot{
		iteration: iteration,
		array:     arrCopy,
		indexes:   indexes,
	}
}

func (s *Snapshot) SetDescription(value string) {
	s.description = value
}
