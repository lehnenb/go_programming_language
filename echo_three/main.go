package main

import (
	"os"
	"strings"
)

// Echo prints CLI arguments
func Echo(args []string) {
	println(strings.Join(args[1:], " "))
}

func main() {
	Echo(os.Args)
}
