package main

// from main.go

import (
	"log"
	"os"
	"strconv"
	"testing"
)

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

// BenchmarkEcho is a function responsible for benchmarking the echo
func BenchmarkEcho(b *testing.B) {
	defer quiet()()
	args := []string{}

	for i := 0; i < b.N; i++ {
		args = append(args, "args"+strconv.Itoa(i))
		Echo(args)
	}
}
