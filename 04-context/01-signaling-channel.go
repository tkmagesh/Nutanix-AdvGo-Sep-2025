package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genData(stopCh)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}

	fmt.Println("Done")
}

func genData(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)

	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-stopCh:
				fmt.Println("Stop signal received")
				break LOOP
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}
