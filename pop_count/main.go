package main

import (
	"fmt"
	"os"
)


func main() {
  Bitwiseoperations();
}

//Bitwiseoperations Try a bunch of things with Bitwise operations
func Bitwiseoperations() {
  ten := 0b1010
  twelve := 0b1100


  fmt.Printf("%b \n", twelve)
  fmt.Printf("%b \n", twelve << 1)

  fmt.Printf("%d \n", ten)
  fmt.Printf("%d \n", twelve)
}

//FileTryouts Try a bunch of things with files
func FileTryouts() {
  var wd string
  wd, err := os.Getwd()

  if err != nil {
    panic("damn")
  }
  fmt.Printf("Working dir: %s \n", wd)

  f, err := os.Open("/Users/brunolehnen/Desktop/changes.txt")

  if err != nil {
    panic("damn /dev/null")
  }

  stat, _ := f.Stat()
  data := make([]byte, stat.Size())

  fmt.Printf("%d", stat.Size())
  f.Read(data)

  fmt.Printf("%s", data)

}
