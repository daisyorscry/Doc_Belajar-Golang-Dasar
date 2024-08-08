package repository

import (
	"context"
	"database/sql"
	model "golang_database/Model"
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
	return user, err
}

func (r *UserRepository) FindAll(ctx context.Context, user model.Users) ([]model.Users, error) {

	// ini adalah sql scriptnya
	query := "SELECT id,name,nim FROM admin"

	// ini adlaah contoh membuat query row contex
	// jadi panggil dulu struct yang menyimpan sql.dbnya kemudian masukkan fild apa yang mau di pakai misalaya queryrowcontext
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		// nah kalau error return struct kosong aja terus kembalikan erronya
		return nil, err
	}

	// buat sebuah variabel yang nanti akan mengembalikan nilai dari
	var users []model.Users
	for rows.Next() {
		var user model.Users
		if err := rows.Scan(user.Id, user.Name, user.Nim); err != nil {
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
	query := "SELECT id, name, nim where id = $1"

	var user model.Users
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&user.Id, user.Name, user.Nim)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Users{}, err
		}
		return model.Users{}, err
	}
	return user, err
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM admin where id "

	// var id int64
	_, err := r.DB.QueryRowContext(ctx, query, id)
	return err
}
