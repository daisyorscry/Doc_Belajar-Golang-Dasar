package main

import (
	"context"
	"testing"
	"time"
)

func TestPrepareStatement(t *testing.T) {
	conn := GetConnection()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := "INSERT INTO admin(name, password) VALUES($1,$2)"
	stmt, err := conn.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000000; i++ {
		_, err = stmt.ExecContext(ctx, "admin", "123")
		if err != nil {
			panic(err)
		}
	}
	defer stmt.Close()

}
