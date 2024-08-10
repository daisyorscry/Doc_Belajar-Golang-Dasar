package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// ctx := context.Background()
	// conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(1000)
	conn.SetConnMaxIdleTime(time.Second * 5)
	conn.SetConnMaxLifetime(time.Minute * 60)

	return conn
}
