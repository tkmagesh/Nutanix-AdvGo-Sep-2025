package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	ID    int     `label:"Product ID"`
	Name  string  `label:"Product Name"`
	Price float64 `label:"Price (USD)"`
	Stock int     // No tag here
}

func PrintLabeledFields(input any) {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() != reflect.Struct {
		fmt.Println("Input must be a struct")
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		label := field.Tag.Get("label")
		if label == "" {
			continue // Skip fields without 'label' tag
		}

		value := val.Field(i).Interface()
		fmt.Printf("%s: %v\n", label, value)
	}
}

func main() {
	p := Product{
		ID:    101,
		Name:  "Laptop",
		Price: 1299.99,
		Stock: 25,
	}

	PrintLabeledFields(p)
}
