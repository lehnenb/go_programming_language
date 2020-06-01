// Dup 2 prints the count and text of lines that appear
// more than once in the input.
// It reads from stdin or from a list of named files

package main

import (
	"bufio"
	"fmt"
	"os"
)

type keyVal struct {
	Filename string
	Line     string
}

func main() {
	files := os.Args[1:]
	counts := make(map[keyVal]int)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\n", line.Filename)
			fmt.Printf("%d\t%s\n", n, line.Line)
		}
	}
}

func countLines(f *os.File, counts map[keyVal]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[keyVal{Filename: f.Name(), Line: input.Text()}]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
