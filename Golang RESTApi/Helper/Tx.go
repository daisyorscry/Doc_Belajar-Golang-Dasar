package helper

import (
	"database/sql"
	"log"
)

func TxHandler(tx *sql.Tx, err error) {
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			log.Printf("Transaction rollback error: %v", rerr)
		} else {
			log.Printf("Transaction rolled back due to error: %v", err)
		}
	} else {
		if cerr := tx.Commit(); cerr != nil {
			log.Printf("Transaction commit error: %v", cerr)
		}
	}
}
