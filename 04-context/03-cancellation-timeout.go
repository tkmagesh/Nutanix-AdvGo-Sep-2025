/*
time based cancellation using context.WithTimeout()
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	go func() {
		fmt.Println("Will stop after 5 secs.. but Hit ENTER to manually stop...")
		fmt.Scanln()
		cancel()
	}()
	ch := genNos(ctx)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genNos(ctx context.Context) <-chan int {

	ch := make(chan int)
	go func() {
	LOOP:
		for no := 1; ; no++ {
			select {
			case <-ctx.Done():
				break LOOP
			default:
				ch <- no * 10
			}
		}
		switch err := ctx.Err(); err {
		case context.Canceled:
			fmt.Println("[genNos] programmatic cancellation signal received")
		case context.DeadlineExceeded:
			fmt.Println("[genNos] timeout cancellation signal received")
		default:
			fmt.Println("[genNos] cancellation due to unknown error")
		}
		close(ch)
	}()
	return ch
}
