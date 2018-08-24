package vm

import (
	"testing"
)

type NoopWriter struct{}

func (nw *NoopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

var noopWriter = &NoopWriter{}

func BenchmarkPRINT(b *testing.B) {
	program := []int{
		ICONST, 42,
		PRINT,
		HALT,
	}
	for n := 0; n < b.N; n++ {
		vm := NewVM(program, make(map[int]interface{}))
		vm.Run(noopWriter)
	}
}

func BenchmarkCALL(b *testing.B) {
	program := []int{
		ICONST, 42,
		ICONST, 43,
		ICONST, 50,
		CALL, 10, 2,
		HALT,
		LOAD, 0,
		LOAD, 1,
		IADD,
		PRINT,
		CALL, 20, 1,
		RET,
		LOAD, 0,
		PRINT,
		RET,
	}
	for n := 0; n < b.N; n++ {
		vm := NewVM(program, make(map[int]interface{}))
		vm.Run(noopWriter)
	}
}
