package conc

import (
	"fmt"
	"sync"
)

// Subscription example, and a good repo for reference
// https://github.com/lotusirous/go-concurrency-patterns/blob/main/14-adv-subscription/main.go

// Use a mutex to manage the state. Probably that would hold the lines?
// or other pieces of data that are calculated and then reused by other parts of the program?
// https://gobyexample.com/mutexes

type Rating struct {
	Binary  []int
	Decimal int
}

type LinesContainer struct {
	Mu             sync.Mutex
	Lines          []string
	OxygenRating   Rating
	ScrubberRating Rating
}

func (lc *LinesContainer) SetLines(lines []string) {
	lc.Mu.Lock()
	defer lc.Mu.Unlock()
	lc.Lines = lines
}

func RunConcurrent() {
	fmt.Println("Running Concurrent Solution")
}
