package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ReadLines() []string {
  content, err := ioutil.ReadFile("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  linesString := string(content)
  lines := strings.Split(linesString, "\n")
  return lines
}

func main() {
  lines := ReadLines()
  fmt.Println("Number of lines: ", len(lines))
}
