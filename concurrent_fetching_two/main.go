package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func removeSubstrings(needles []string, haystack string) string {
	replacedHaystack := haystack

	for _, needle := range needles {
		replacedHaystack = strings.ReplaceAll(replacedHaystack, needle, "")
	}

	return replacedHaystack
}

func writeToFile(url string, contents io.ReadCloser) (int64, error) {
	filename := removeSubstrings([]string{"http://", "https://"}, url)
	path := fmt.Sprintf("/tmp/%s.html", filename)
	file, err := os.Create(path)

	defer file.Close()

	if err != nil {
		return 0, err
	}

	nbytes, err := io.Copy(file, contents)

	return nbytes, nil
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := writeToFile(url, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
