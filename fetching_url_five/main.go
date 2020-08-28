package main

// Fetch prints the content found at a URL.

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
		os.Exit(1)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs  %7d  %s \n", secs, nbytes, url)
}

func main() {
	start := time.Now()

	for _, url := range os.Args[1:] {
		fetch(url)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("Total time: %.2fs", secs)
}
