package main

import (
	"fmt"
	"time"
)

var superbStart int = 1
var superbStep int = 1
var superbLength int = 0
var threats = make(chan bool, 16)

func isSquare(number int) bool {
	for n := 0; n*n <= number; n++ {
		if number == n*n {
			return true
		}
	}
	return false
}

func findSuperb() (start int, step int, length int) {
	startTime := time.Now()
	for i := 1; i < SQUARES_LIMIT; i++ {
		if i % 1000 == 0 {
			fmt.Printf("Trying %dth square.\n", i)
		}
		threats <- true
		square := i * i
		//fmt.Printf("Firing GoRoutine with start: %d.\n", prime)
		go tryOutStart(square)
	}
	for i := 0; i < len(threats); i++ {
		threats <- true
	}
	fmt.Println("#########")
	fmt.Printf("With the given LIMITS, the best sequence has start: %d, step: %d and length: %d. \n",
		superbStart, superbStep, superbLength)
	fmt.Printf("Total time: %v. \n", time.Since(startTime))
	fmt.Println("#########")
	return superbStart, superbStep, superbLength
}

func tryOutStart(start int) {
	for step := 1; step < STEPS_LIMIT; step++ {
		tryOutSequence(start, step)
	}
	_ = <-threats
}

func tryOutSequence(start int, step int) {
	last := start
	length := 0
	for ; isSquare(last); length++ {
		last = last + step
	}
	//fmt.Printf("Sequence with start: %d and step: %d was: %d long. \n", start, step, length)
	if length > superbLength {
		superbLength = length
		superbStart = start
		superbStep = step
	}
}

func checkSquares(start int, step int, length int) bool {
	last := start
	for i := 0; i < length; i++ {
		if isSquare(last) == false {
			return false
		}
		last = last + step
	}
	return true
}

/*
#########
With the given LIMITS, the best sequence has start: 1, step: 24 and length: 3.
Total time: 29.2225018s.
#########
 */

var SQUARES_LIMIT = 16384
var STEPS_LIMIT = 16384
func main() {
	//findSuperb()
	fmt.Println(checkSquares(1, 24, 3))
}