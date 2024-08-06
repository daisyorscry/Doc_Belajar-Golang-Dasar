package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	for i := 0; i < 10; i++ {
		// time.Sleep(time.Second * 9)
		_, err := conn.ExecContext(ctx, "UPDATE name SET name=$1 WHERE name=$2", "daisy", "jerry")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("sukses insert ke database")
}
