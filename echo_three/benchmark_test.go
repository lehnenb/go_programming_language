package main

// from main.go

import (
	"strconv"
	"testing"
)

// BenchMarkEcho is a function responsible for benchmarking the echo
func BenchmarkEcho(b *testing.B) {
	args := []string{"um", "cepo", "de", "madeira"}

	for i := 0; i < b.N; i++ {
		Echo(args)
		args = append(args, "args"+strconv.Itoa(i))
	}
}
