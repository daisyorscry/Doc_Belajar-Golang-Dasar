package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSQLInjection(t *testing.T) {
	conn := GetConnection()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	type Users struct {
		name     string
		password string
	}

	var person Users
	name := "admin"
	// name := "admin'; #"
	// name := "admin' --"
	password := "123" // Password should be a string
	query := "SELECT name, password FROM admin WHERE name = $1 AND password = $2"
	err := conn.QueryRowContext(ctx, query, name, password).Scan(&person.name, &person.password)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(person.name, person.password)
}
