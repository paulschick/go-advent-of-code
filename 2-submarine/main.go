package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func GetDirectionsMap(lines []string) map[string]int {
  // holds the direction and the total sum of all
  // magnitudes for that direction
  directions := make(map[string]int)

  for _, line := range lines {
    lineSplit := strings.Split(line, " ") 

    currentMagnitude, exists := directions[lineSplit[0]]
    
    directionTrimmed := strings.TrimSpace(lineSplit[1])
    lineValue, err := strconv.ParseInt(directionTrimmed, 0, 64)

    if err != nil {
      log.Fatal(err)
    }

    if exists {
      directions[lineSplit[0]] = currentMagnitude + int(lineValue)
    } else {
      directions[lineSplit[0]] = int(lineValue)
    }
  }

  return directions
}

func main() {
  lines := ReadLines()
  directions := GetDirectionsMap(lines) 
  fmt.Println(directions)
}
