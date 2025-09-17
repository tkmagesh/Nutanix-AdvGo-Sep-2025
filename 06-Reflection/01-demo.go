package main

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateInsertQuery(tableName string, data any) (string, []any, error) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	if v.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("expected struct type")
	}

	var columns []string
	var placeholders []string
	var values []any

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue // skip fields with no db tag
		}
		columns = append(columns, tag)
		placeholders = append(placeholders, "?")
		values = append(values, v.Field(i).Interface())
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, values, nil
}

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func main() {
	user := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	query, args, err := GenerateInsertQuery("users", user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)
}
