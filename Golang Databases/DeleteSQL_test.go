package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	// time.Sleep(time.Second * 9)
	_, err := conn.ExecContext(ctx, "DELETE FROM name WHERE name=$1", "jerry")
	if err != nil {
		panic(err)
	}

	fmt.Println("sukses insert ke database")
}
