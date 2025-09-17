package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	x = 19.99
	x = true
	x = "Sint culpa nulla elit eu dolor eu esse qui culpa officia exercitation non."
	x = struct{}{}
	fmt.Println(x)

	// x = 200
	x = "Anim cillum duis officia Lorem enim labore."
	if val, ok := x.(int); ok {
		z := val * 2
		fmt.Println(z)
	} else {
		fmt.Println("x is not an int")
	}

	// x = 200
	// x = "Cupidatat excepteur ex mollit in."
	// x = true
	x = struct{}{}
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is an string")
	case bool:
		fmt.Println("x is an bool")
	default:
		fmt.Println("Unknown type")

	}
}
