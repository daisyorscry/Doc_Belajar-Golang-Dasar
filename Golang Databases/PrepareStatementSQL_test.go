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

	// ini adalah contoh kalau kita menggunakan prepare statement, result dari prepare contex itu 2 2tatement dan error
	// statement ini sangat baik di lakukan untuk menghindari terjadinya sql injection juga
	// karena mau ga mau kita harus
	// statement ini yang akan di panggil untuk melakukan exec contex
	admin := "admin"
	password := "123"
	for i := 0; i < 100; i++ {
		_, err = stmt.ExecContext(ctx, admin, password)
		if err != nil {
			panic(err)
		}
	}

	defer stmt.Close()

}
