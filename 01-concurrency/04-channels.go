package main

import "fmt"

func main() {
	/*
		ch := make(chan int)
		data := <-ch //blocking
		ch <- 100
		fmt.Println(data)
	*/

	ch := make(chan int)
	ch <- 100
	data := <-ch //blocking
	fmt.Println(data)
}
