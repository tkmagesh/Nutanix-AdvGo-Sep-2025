package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Println(sumInts([]int{3, 1, 4, 2, 5}))
		fmt.Println(sumFloats([]float64{3.0, 1.9, 4.3, 2.7, 5.2}))
	*/
	fmt.Println(sum([]int{3, 1, 4, 2, 5}))
	fmt.Println(sum([]float64{3.0, 1.9, 4.3, 2.7, 5.2}))

}

func sumInts(list []int) int {
	var result int
	for _, no := range list {
		result += no
	}
	return result
}

func sumFloats(list []float64) float64 {
	var result float64
	for _, no := range list {
		result += no
	}
	return result
}

type MyNumerics interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sum[T MyNumerics](list []T) T {
	var result T
	for _, no := range list {
		result += no
	}
	return result
}
