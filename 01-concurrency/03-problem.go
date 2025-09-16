package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start and end..")
	fmt.Scanln(&start, &end)
	fmt.Printf("Generating prime numbers from %d to %d\n", start, end)
	primes := generatePrimes(start, end)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func generatePrimes(start, end int) []int {
	var primes []int
LOOP:
	for no := start; no <= end; no++ {
		// prime loop
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				// not a prime number
				// move on to the next number
				continue LOOP
			}
		}
		// prime number, add it to the slice
		primes = append(primes, no)
	}
	return primes
}
