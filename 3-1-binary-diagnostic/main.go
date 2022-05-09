package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	// "unicode/utf8"
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

// Deprecated: following MostCommonNumber, this will be replaced with
// a conversion from a byte, not a rune
// func IntFromRune(character rune) int64 {
//   buf := make([]byte, 1) 
//   _ = utf8.EncodeRune(buf, character)
//   value, _ := strconv.Atoi(string(buf))
//   return int64(value)
// }

func MostCommonNumber(lines []string) []int64 {
  gamma := []int64 {}

  trimmed := strings.TrimSpace(lines[0])
  colCount := len(trimmed)

  for colIndex := 0; colIndex < colCount; colIndex++ {
    countOnes := 0
    countZeros := 0

    for _, row := range lines {
      valInt, _ := strconv.Atoi(string(row[colIndex]))

      if valInt == 0 {
        countZeros += 1
      } else if valInt == 1 {
        countOnes += 1
      }
    }

    if countOnes > countZeros {
      gamma = append(gamma, int64(1))
    } else if countOnes == countZeros {
      gamma = append(gamma, int64(1))
    } else {
      gamma = append(gamma, int64(0))
    }
  }
  return gamma
}

func main() {
  gamma := MostCommonNumber(ReadLines())
  fmt.Println("Gamma: ", gamma)
}
