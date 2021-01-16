package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
  "sync"
)

type Line struct { 
  Number int
  Contents string
}

type Result struct { 
  Match bool 
  Line Line
}

func main() {
  word := flag.String("w", "", "Word to search for")
	filePath := flag.String("f", "", "File in which the word will be searched")

  flag.Parse()

	if *word == "" {
		fmt.Println("a word needs to be provided")
    os.Exit(1)
	}

	if *filePath == "" {
		fmt.Println("a file needs to be provided")
    os.Exit(1)
	}

  findWord(word, filePath)
}

func findWord(word *string, filePath *string) {
  var wg sync.WaitGroup

  file, err := os.Open(*filePath)
	scanner := bufio.NewScanner(file)
  
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  lineCh := make(chan Line)
  results := make(chan Result)

  for i:=0; i < 8; i++ {
    go lineWorker(i, *word, lineCh, results)
  }

  for i:=0; scanner.Scan(); i++ {
    wg.Add(1)
    lineCh <- Line { Number: i + 1, Contents: scanner.Text() } 

    go func() {
      result := <-results

      if (result.Match) {
        fmt.Printf("Word %s found in line %d \n", *word, result.Line.Number)
      }

      wg.Done()
    }()
  }

  wg.Wait()

  close(lineCh)
  close(results)
}

func lineWorker(i int, word string, lineCh <-chan Line, results chan <-Result) {
  for line := range lineCh {
    words := strings.Split(line.Contents, " ")
    match := false
    fmt.Printf("Worker #%d processing line #%d\n", i + 1, line.Number)

    for _, w := range words {
      if (w == word) {
        match = true
        break;
      }
    }

    results <- Result { Match: match, Line: line }
  }
}

