package vm

type Frame struct {
	variables     map[int]int
	returnAddress int
}
