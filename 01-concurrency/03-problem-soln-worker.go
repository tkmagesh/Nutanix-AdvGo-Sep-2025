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
	primes := generatePrimes(start, end, 10)
	for primeNo := range primes {
		fmt.Printf("[main] Prime No : %d\n", primeNo)
	}
	fmt.Println("[main] Done!")
}

// producer
func generatePrimes(start, end int, workerCount int) <-chan int {
	primesCh := make(chan int)
	dataCh := make(chan int)

	// feeder
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()

	wg := &sync.WaitGroup{}
	for id := range workerCount {
		wg.Add(1)
		go primeWorker(id, wg, dataCh, primesCh)
	}

	// wait for the workers to complete and close the 'output' channel
	go func() {
		wg.Wait()
		close(primesCh)
	}()

	return primesCh
}

func primeWorker(workerId int, wg *sync.WaitGroup, dataCh <-chan int, primeCh chan<- int) {
	defer wg.Done()
	for no := range dataCh {
		fmt.Printf("[primeWorker] Worker [%d] processes %d\n", workerId, no)
		if isPrime(no) {
			primeCh <- no
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
