package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for id := range 100 {
		wg.Go(func() {
			fn(id)
		})
	}
	wg.Wait() // block until the counter becomes 0
	fmt.Println("done!!")
}

func fn(id int) {
	fmt.Printf("fn[%d] started...\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed...\n", id)
}
