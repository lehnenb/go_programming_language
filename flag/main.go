package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
  lines, err := readByLine(filePath);

  if (err != nil) {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  numLines := len(lines)
  lineCh := make(chan Line, numLines)
  results := make(chan Result, numLines)

  for i:=0; i < 3; i++ {
    go lineWorker(*word, lineCh, results)
  }

  for i, line := range lines {
    lineCh <- Line { Number: i, Contents: line } 
  }

  close(lineCh)

  for i:=0; i < numLines; i++ {
    result := <-results

    if (result.Match) {
      fmt.Printf("Word %s found in line %d \n", *word, result.Line.Number)
    }
  }

  close(results)
}

func lineWorker(word string, lineCh <-chan Line, results chan <-Result) {
  for line := range lineCh {
    words := strings.Split(line.Contents, " ")
    match := false

    for _, w := range words {
      if (w == word) {
        match = true
        break;
      }
    }

    results <- Result { Match: match, Line: line }
  }
}

func readByLine(filePath *string) ([]string, error) {
  lines := make([]string, 1)
  file, err := os.Open(*filePath)
	scanner := bufio.NewScanner(file)

  if err != nil {
    return nil, err
  }

  for(scanner.Scan()) {
    text := scanner.Text()
    lines = append(lines, text)
  }

  return lines, nil
}
