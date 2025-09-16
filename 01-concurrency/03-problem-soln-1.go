package main

import (
	"fmt"
	"sync"
)

// consumer
func main() {
	var start, end int
	fmt.Println("Enter the start and end..")
	fmt.Scanln(&start, &end)
	fmt.Printf("Generating prime numbers from %d to %d\n", start, end)
	primes := generatePrimes(start, end)
	for primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

// producer
func generatePrimes(start, end int) <-chan int {
	var primes = make(chan int)
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		// prime loop
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				primes <- no
			}
		}()
	}
	go func() {
		wg.Wait()
		close(primes)
	}()
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
