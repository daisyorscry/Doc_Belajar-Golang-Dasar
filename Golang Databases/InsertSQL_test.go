package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*30)
	defer cancel()
	for i := 0; i < 9153; i++ {
		// time.Sleep(time.Second * 9)
		_, err := conn.ExecContext(ctx, "INSERT INTO name(name, nim) values('jerry', 1002230006)")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("sukses insert ke database")
}

func TestInsertUsers(t *testing.T) {

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*30)
	defer cancel()
	for i := 0; i < 1000; i++ {
		// time.Sleep(time.Second * 9)
		_, err := conn.ExecContext(ctx, "INSERT INTO users(name, nim, rating, birthday, married) values('jerry', 1002230006, 5.0, null, false)")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("sukses insert ke database")
}
