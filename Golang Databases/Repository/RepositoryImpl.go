package repository

import (
	"context"
	"database/sql"
	model "golang_database/Model"
	"log"
)

// buat sebuah struct yang isinya sql.db yang nanti akan mengimplementasi interace tadi
type UserRepository struct {
	DB *sql.DB
}

// ini adalah contoh untuk membuat salah satu iplementasi dari interfacenya

// buat sebuah function yang adalah implementasi dari interfacenya tadi

// nah di awal sebuah functionya pastikan untuk memasukkan strucnya agar dia implentasi dari interfacenya
// func <struct sql.db sebagai impmlentasi> insert(parameter)return
func (r *UserRepository) Insert(ctx context.Context, user model.Users) (model.Users, error) {

	// ini adalah sql scriptnya
	query := "INSERT INTO users (name, nim) VALUES ($1, $2) RETURNING id"

	// buat sebuah variabel yang nanti akan mengembalikan nilai dari
	var id int

	// ini adlaah contoh membuat query row contex.
	// jadi panggil dulu struct yang menyimpan sql.dbnya kemudian masukkan fild apa yang mau di pakai misalaya queryrowcontext
	err := r.DB.QueryRowContext(ctx, query, user.Name, user.Nim).Scan(&id)
	if err != nil {
		// nah kalau error return struct kosong aja terus kembalikan erronya
		return model.Users{}, err
	}

	user.Id = int32(id)
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user model.Users) (model.Users, error) {

	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	tx, err := r.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return model.Users{}, err
	}
		
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Printf("Rollback error: %v", rerr)
			}
		} else {
			if err := tx.Commit(); err != nil {
				log.Printf("Commit error: %v", err)
			}
		}
	}()

	// ini adalah sql scriptnya
	query := "UPDATE users SET name = $1, nim = $2  RETURNING id"

	// buat sebuah variabel yang nanti akan mengembalikan nilai dari
	var id int

	// ini adlaah contoh membuat query row contex.
	// jadi panggil dulu struct yang menyimpan sql.dbnya kemudian masukkan fild apa yang mau di pakai misalaya queryrowcontext
	err = tx.QueryRowContext(ctx, query, user.Name, user.Nim).Scan(&id)
	if err != nil {
		// nah kalau error return struct kosong aja terus kembalikan erronya
		return model.Users{}, err
	}

	user.Id = int32(id)
	return user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]model.Users, error) {

	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  true,
	}

	tx, err := r.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Printf("Rollback error: %v", rerr)
			}
		} else {
			if err := tx.Commit(); err != nil {
				log.Printf("Commit error: %v", err)
			}
		}
	}()

	// ini adalah sql scriptnya
	query := "SELECT id,name,nim FROM users"

	// ini adlaah contoh membuat query row contex
	// jadi panggil dulu struct yang menyimpan sql.dbnya kemudian masukkan fild apa yang mau di pakai misalaya queryrowcontext
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		// nah kalau error return struct kosong aja terus kembalikan erronya
		return nil, err
	}

	defer rows.Close()

	// buat sebuah variabel yang nanti akan mengembalikan nilai dari
	var users []model.Users
	for rows.Next() {
		var user model.Users
		if err := rows.Scan(&user.Id, &user.Name, &user.Nim); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (model.Users, error) {

	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	tx, err := r.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return model.Users{}, err
	}

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Printf("Rollback error: %v", rerr)
			}
		} else {
			if err := tx.Commit(); err != nil {
				log.Printf("Commit error: %v", err)
			}
		}
	}()

	query := "SELECT  id, name, nim FROM users where id = $1"

	var user model.Users
	err = tx.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Name, &user.Nim)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Users{}, err
		}
		return model.Users{}, err
	}
	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) (bool, error) {
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	tx, err := r.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return false, err
	}

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.Printf("Rollback error: %v", rerr)
			}
		} else {
			if err := tx.Commit(); err != nil {
				log.Printf("Commit error: %v", err)
			}
		}
	}()

	query := "DELETE FROM users where id = $1 "

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		// Jika tidak ada baris yang dihapus, return false
		return false, nil
	}

	return true, nil

}


// teesting