package jd

const (
	DropTargetJ = 1 << iota
	DropTargetU
	DropTargetD
	DropTargetG
	DropTargetE
)

var (
	DropTargetIndexes = map[string]int{
		SwitchDropTargetJ: 0,
		SwitchDropTargetU: 1,
		SwitchDropTargetD: 2,
		SwitchDropTargetG: 3,
		SwitchDropTargetE: 4,
		LampDropTargetJ:   0,
		LampDropTargetU:   1,
		LampDropTargetD:   2,
		LampDropTargetG:   3,
		LampDropTargetE:   4,
	}
)
