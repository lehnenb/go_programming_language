package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo prints CLI arguments
func Echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func main() {
	Echo(os.Args)
}
