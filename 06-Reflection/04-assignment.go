package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	lAdd := logWrapper(add)
	lSubtract := logWrapper(subtract)

	lAdd(100, 200)
	lSubtract(100, 200)
}

func add(x, y int) {
	fmt.Println("Add Result : ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result : ", x-y)
}

func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}

func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started")
	op(x, y)
	log.Println("Operation completed")
}

func logWrapper(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		op(x, y)
		log.Println("Operation completed")
	}
}

/* Create a generalized logWrapper that can work with any function with any number of arguments of any type */
