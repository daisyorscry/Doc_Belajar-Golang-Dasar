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
	password := "123"
	//
	// cara untuk menghindari sql injection itu mudah
	// cukup untuk tiidak pernah hard code value yang akan di insert

	// taruh dulu kedala variabel dan untuk valuenya cukup gunakan $1 untuk value yang paling pertana dan seterusny
	query := "SELECT name, password FROM admin WHERE name = $1 AND password = $2"

	// nah misal setelah stript di parameter query contex ada lagi yang namannya arguments, ini untuk membaca semua variabel untuk value yang akan masuk ke dalam value daatabasenya
	err := conn.QueryRowContext(ctx, query, name, password).Scan(&person.name, &person.password)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(person.name, person.password)
}
