/*
programmatic cancellation using context.WithCancel()
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Println("Hit ENTER to stop...")
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
		fmt.Println("[genNos] cancellation signal received")
		close(ch)
	}()
	return ch
}
