package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
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
  fmt.Println(lines[0])
  for _, character := range lines[0] {
    buf := make([]byte, 1)
    _ = utf8.EncodeRune(buf, character)
    value, _ := strconv.Atoi(string(buf))
    fmt.Println(value)
    fmt.Printf("%T\n", value)
  }
}
