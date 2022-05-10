package main

import (
	"binary-diagnostic/concurrent-solution"
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

func GetMaxInColumn(lines []string, colIndex int) int {
	num1 := 0
	num0 := 0
	for _, line := range lines {
		val := line[colIndex]
		valInt, _ := strconv.Atoi(string(val))
		if valInt == 1 {
			num1 += 1
		} else {
			num0 += 1
		}
	}

	if num1 >= num0 {
		return 1
	} else {
		return 0
	}
}

func GetMinInColumn(lines []string, colIndex int) int {
	num1 := 0
	num0 := 0
	for _, line := range lines {
		val := line[colIndex]
		valInt, _ := strconv.Atoi(string(val))
		if valInt == 1 {
			num1 += 1
		} else {
			num0 += 1
		}
	}

	if num1 >= num0 {
		return 0
	} else {
		return 1
	}
}

func StringToIntSlice(line string) []int {
	trimmed := strings.TrimSpace(line)
	lineInts := []int{}

	for _, character := range trimmed {
		valInt, _ := strconv.Atoi(string(character))
		lineInts = append(lineInts, valInt)
	}

	return lineInts
}

func TrimRows(lines []string, colIndex int, keepNum int) []string {
	trimmed := []string{}

	for _, line := range lines {
		valInt, _ := strconv.Atoi(string(line[colIndex]))
		if valInt == keepNum {
			trimmed = append(trimmed, line)
		}
	}

	return trimmed
}

func GetOxygenRating(lines []string, colIndex int) []int {
	rowLength := len(lines)
	trimmed := strings.TrimSpace(lines[0])

	if rowLength == 1 {
		return StringToIntSlice(lines[0])
	}

	var newVec []string
	maxNum := GetMaxInColumn(lines, colIndex)

	if maxNum == 1 {
		newVec = TrimRows(lines, colIndex, 1)
	} else {
		newVec = TrimRows(lines, colIndex, 0)
	}

	nextIndex := colIndex + 1

	if nextIndex < len(trimmed) {
		return GetOxygenRating(newVec, nextIndex)
	} else {
		return StringToIntSlice(newVec[0])
	}
}

func GetScrubberRating(lines []string, colIndex int) []int {
	rowLength := len(lines)
	trimmed := strings.TrimSpace(lines[0])

	if rowLength == 1 {
		return StringToIntSlice(lines[0])
	}

	var newVec []string
	minNum := GetMinInColumn(lines, colIndex)

	if minNum == 1 {
		newVec = TrimRows(lines, colIndex, 1)
	} else {
		newVec = TrimRows(lines, colIndex, 0)
	}

	nextIndex := colIndex + 1

	if nextIndex < len(trimmed) {
		return GetScrubberRating(newVec, nextIndex)
	} else {
		return StringToIntSlice(newVec[0])
	}
}

func GetDecimal(values []int) int {
	highestPower := len(values) - 1
	currentValue := 0
	base := 2

	for i := 0; i <= highestPower; i++ {
		power := highestPower - i
		poweredBase := math.Pow(float64(base), float64(power))
		currentInt := int(values[i])
		result := currentInt * int(poweredBase)
		currentValue = currentValue + result
	}

	return currentValue
}

func Calculate() {
	lines := ReadLines()
	value := GetDecimal(GetOxygenRating(lines, 0)) * GetDecimal(GetScrubberRating(lines, 0))
	fmt.Println("Multiplied Value: ", value)
}

func main() {
	Calculate()
	conc.RunConcurrent()
}
