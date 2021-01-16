package main

import (
	"fmt"
	"os"
)

// Echo prints CLI arguments
func Echo(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	Echo(os.Args)
}
