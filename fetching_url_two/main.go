package main

// Fetch prints the content found at a URL.

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
		os.Exit(1)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		fetch(url)
	}
}
