package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Directions struct {
  Aim int64
  Horizontal int64
  Depth int64
}

func NewDirections() *Directions {
  return &Directions {
    Aim: 0,
    Horizontal: 0,
    Depth: 0,
  }
}

func (d *Directions) SetForward(forward int64) {
  d.Horizontal = d.Horizontal + forward
  d.Depth = d.Depth + (d.Aim * forward)
}

func (d *Directions) SetUp(up int64) {
  d.Aim = d.Aim - up
}

func (d *Directions) SetDown(down int64) {
  d.Aim = d.Aim + down
}

func (d *Directions) CalculatePosition() int64 {
  return d.Horizontal * d.Depth
}

func (d *Directions) ParseMovements(lines []string) {
  for _, line := range lines {
    lineSplit := strings.Split(line, " ")
    // valueString := strings.TrimSpace(lineSplit[1])
    valueInt, err := strconv.ParseInt(strings.TrimSpace(lineSplit[1]), 0, 64)
    if err != nil {
      log.Fatal(err)
    }
    magnitude := int64(valueInt)

    direc := lineSplit[0]

    if direc == "forward" {
      d.SetForward(magnitude)
    } else if direc == "up" {
      d.SetUp(magnitude)
    } else if direc == "down" {
      d.SetDown(magnitude)
    }

  }
}

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
  directions := NewDirections()
  directions.ParseMovements(ReadLines())
  position := directions.CalculatePosition()
  fmt.Println(position)
}
