package main

import (
  "fmt"
	"io/ioutil"
	"log"
	"math"
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

func GetDecimal(values []int64) int64 {
  highestPower := len(values) - 1
  currentValue := int64(0)
  base := 2

  for i := 0; i <= highestPower; i++ {
    power := highestPower - i
    poweredBase := math.Pow(float64(base), float64(power))
    currentInt := int(values[i])
    result := currentInt * int(poweredBase)
    currentValue = currentValue + int64(result)
  }

  return currentValue
}

func main() {
  // lines := ReadLines()
  fmt.Println("test")
}
