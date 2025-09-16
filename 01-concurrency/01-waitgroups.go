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
		wg.Add(1) // increment the counter by 1
		go fn(id, wg)
	}
	wg.Wait() // block until the counter becomes 0
	fmt.Println("done!!")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter by 1
	fmt.Printf("fn[%d] started...\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed...\n", id)
}
