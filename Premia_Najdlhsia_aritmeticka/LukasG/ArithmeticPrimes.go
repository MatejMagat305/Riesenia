package main

import (
	"fmt"
	"time"
)

var lastPrime int = 1
var bestStart int = 2
var bestStep int = 1
var bestLength int = 0
var routines = make(chan bool, 16)

func isPrime(number int) bool {
	if number <= 3  {
		return true
	}
	if number % 2 == 0 || number % 3 == 0 {
		return false
	}
	for i :=5; i*i <= number; i += 6 {
		if number % i == 0 || number % (i + 2) == 0 {
			return false
		}
	}
	return true
}

func nextPrime() int {
	if lastPrime == 1 {
		lastPrime = 2
		return lastPrime
	}
	number := lastPrime + 1
	for ; isPrime(number) == false; number++ { }
	lastPrime = number
	return lastPrime
}

func findBest() (start int, step int, length int) {
	startTime := time.Now()
	for i := 0; i < LIMIT_PRIMES; i++ {
		if i % 1000 == 0 {
			fmt.Printf("Trying %dth prime.\n", i)
		}
		routines <- true
		prime := nextPrime()
		//fmt.Printf("Firing GoRoutine with start: %d.\n", prime)
		go tryStart(prime)
	}
	for i := 0; i < len(routines); i++ {
		routines <- true
	}
	fmt.Println("#########")
	fmt.Printf("With the given LIMITS, the best sequence has start: %d, step: %d and length: %d. \n",
		bestStart, bestStep, bestLength)
	fmt.Printf("Total time: %v. \n", time.Since(startTime))
	fmt.Println("#########")
	return bestStart, bestStep, bestLength
}

func tryStart (start int) {
	for step := 1; step < LIMIT_STEPS ; step++ {
		trySequence(start, step)
	}
	_ = <-routines
}

func trySequence(start int, step int) {
	last := start
	length := 0
	for ; isPrime(last); length++ {
		last = last + step
	}
	//fmt.Printf("Sequence with start: %d and step: %d was: %d long. \n", start, step, length)
	if length > bestLength {
		bestLength = length
		bestStart = start
		bestStep = step
	}
}

func printPrimes() {
	for i := 0; i < LIMIT_PRIMES; i++ {
		fmt.Println(nextPrime())
	}
}

func checkSequence(start int, step int, length int) bool {
	last := start
	for i := 0; i < length; i++ {
		if isPrime(last) == false {
			return false
		}
		last = last + step
	}
	return true
}

/*
#########
With the given LIMITS, the best sequence has start: 199, step: 210 and length: 10.
Total time: 3.0003ms.
#########

#########
With the given LIMITS, the best sequence has start: 4943, step: 60060 and length: 13.
Total time: 5m8.7147278s.
#########
*/

var LIMIT_PRIMES = 32768
var LIMIT_STEPS = 32768
func main() {
	//findBest()
	fmt.Println(checkSequence(4943, 60060, 13))
	fmt.Println(checkSequence(199, 210, 10))
}