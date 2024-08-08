package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestTransactional(*testing.T) {

	conn := GetConnection()
	defer conn.Close()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// nah untuk melakukan begintx buat dulu txoptionnya
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	// lakukan begin dulu kirimlan context dan txoptionnnya
	tx, err := conn.BeginTx(ctx, &txOption)
	if err != nil {
		panic("error begin transaction")
	}

	script := "insert into admin (name, password) values($1, $2)"

	username := "admin"
	password := "123"

	// nah unutk melakukan exec cukup seperti ini
	_, err = tx.ExecContext(ctx, script, username, password)
	if err != nil {
		// nah ini untuk menghandle error dan akan melakukan rollback ketika ada eror
		if Rberr := tx.Rollback(); Rberr != nil {
			panic("rollback")
		}
	}

	// commit ini adalah function terakhir yang akan di eksekusi dan akan melakukan insert ketika tidak ada rollback
	if err = tx.Commit(); err != nil {
		panic("error commit")
	}

	fmt.Println("transactional sukses")
}
