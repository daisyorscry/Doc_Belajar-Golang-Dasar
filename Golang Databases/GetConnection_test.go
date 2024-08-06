package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ini adalah contoh kita membuat sebuah connection untuk semua repository yang ada di service
// file url untuk connection ke database di buat dalam .env
// untuk mendownload library env lakukan perintah go get github.com/joho/godotenv
func TestGetConnection(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ini adlah contoh kalau connectionnya menggunakan pgx (recomended)
	// ctx := context.Background()
	// conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	// nah inii adalah open connection menggunakan standart library di golang ini
	// DATABASE_URL itu adalah env yang sudah di buat di atas
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	// kalau error panggil aja panic
	if err != nil {
		panic(err)
	}

	// kalau gada error ini akan di eksekusi artinya koneksinya sukses
	fmt.Println("sukses")

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(1000)
	conn.SetConnMaxIdleTime(time.Minute * 5)
	conn.SetConnMaxLifetime(time.Minute * 60)

}
