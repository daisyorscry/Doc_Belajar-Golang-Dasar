package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {

	conn := GetConnection()
	defer conn.Close()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	_, err := conn.ExecContext(ctx, "DELETE FROM name WHERE name=$1", "jerry")
	if err != nil {
		panic(err)
	}

	fmt.Println("sukses delete ke database")
}
