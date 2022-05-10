package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
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

func gammaCompare(countOnes int, countZeros int) bool {
  if countOnes >= countZeros {
    return true
  } else {
    return false
  }
}

func epsilonCompare(countOnes int, countZeros int) bool {
  if countOnes < countZeros {
    return true
  } else {
    return false
  }
}

func GetGamma(lines []string) []int64 {
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

    if gammaCompare(countOnes, countZeros) {
      gamma = append(gamma, int64(1))
    } else {
      gamma = append(gamma, int64(0))
    }
  }
  return gamma
}

func GetEpsilon(lines []string) []int64 {
  epsilon := []int64 {}

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

    if epsilonCompare(countOnes, countZeros) {
      epsilon = append(epsilon, int64(1))
    } else {
      epsilon = append(epsilon, int64(0))
    }
  }
  return epsilon
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
  lines := ReadLines()
  multiple := GetDecimal(GetGamma(lines)) * GetDecimal(GetEpsilon(lines))
  fmt.Println("Multiplied Result: ", multiple)
}
