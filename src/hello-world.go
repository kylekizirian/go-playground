package main

import (
	"fmt"
	"math"
	"time"
)

func findLargestPrime(max uint64) uint64 {

	//TODO assert input is valid

	start := time.Now()

	primeList := make([]uint64, 1)
	primeList[0] = 2

	for i := uint64(3); i <= max; i = i + 2 {
		// check if prime
		isPrime := true
		for _, num := range primeList {
			if i%num == 0 {
				isPrime = false
				break
			}
			if float64(num) > math.Sqrt(float64(i)) {
				break
			}
		}
		if isPrime {
			primeList = append(primeList, i)
		}
	}
	endTime := time.Since(start)
	fmt.Printf("It took %s to find the largest prime under %d\n", endTime, max)
	return primeList[len(primeList)-1]
}

func main() {
	var primeUnder uint64 = 100000000
	largestPrimeUnder := findLargestPrime(primeUnder)
	fmt.Printf("Largest prime under %d = %d\n", primeUnder, largestPrimeUnder)
}
