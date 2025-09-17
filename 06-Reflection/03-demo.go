package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Product struct {
	ID    int     `label:"Product ID"`
	Name  string  `label:"Product Name"`
	Price float64 `label:"Price (USD)"`
	Stock int     `label:"Available Units"`
}

func GenerateMarkdownDoc(input any) string {
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Struct {
		return "Input is not a struct"
	}

	var sb strings.Builder
	sb.WriteString("| Field   | Label            | Type    |\n")
	sb.WriteString("|---------|------------------|---------|\n")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		label := field.Tag.Get("label")
		if label == "" {
			label = "(no label)"
		}
		sb.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
			field.Name, label, field.Type.Name()))
	}

	return sb.String()
}

func main() {
	md := GenerateMarkdownDoc(Product{})
	fmt.Println(md)
}
