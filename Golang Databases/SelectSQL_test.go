package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

// ini adalah contoh select dan menyimpan datanya di struct dan slice kemudian di iterasi menggunakan for in range
// iterasi menggunakan struct ini sedikit lebih lambat teapi sangat baik untuk di baca dan codenya lebih bersih
func TestSelect(t *testing.T) {

	type Person struct {
		Id, Nim int64
		Name    string
	}

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()

	// query juga mengembalikan 2 tipe data yaitu result dan erro
	// nah kalau query kita wajib menyimpan resultnya
	rows, err := conn.QueryContext(ctx, "SELECT * FROM admin")
	if err != nil {
		panic(err)
	}

	// jadi untuk membaca data hasil query setelah menyimpannya di struct
	// menggunakan struct adalah cara yang lebih baik untuk menyimpan sebuah data dapripada menggunakan variabelnya langsung

	// cara kita meliterasi data adalah menggunakan infinity loop

	// rows ini adalah return value dari query contex yang berisi data dari database dengan mengugnakan function next
	// cara kerja rows.next ini adalah meliterasi 1 per 1 data di database dan ketika datanya habis maka perulangannya akan berhenti
	//
	for rows.Next() {
		var person Person

		//rows.scan ini  adalah untuk menarik datanya dan menyimpannya di variabel
		// pastikan sesuai dengan urutan di tabel
		err := rows.Scan(&person.Id, &person.Name, &person.Nim)

		if err != nil {
			panic(err)
		}

		fmt.Println(person.Id)
		fmt.Println(person.Nim)
		fmt.Println(person.Name)

	}
	defer rows.Close()

	fmt.Println("sukses insert ke database")

}

// ini iterasi menggunakan variabel secara eksplisit
func TestSelectVariabel(t *testing.T) {

	conn := GetConnection()
	conn.Close()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()
	// time.Sleep(time.Second * 9)
	rows, err := conn.QueryContext(ctx, "SELECT * FROM name where id % 2 != 0")
	if err != nil {
		panic(err)
	}

	for rows.Next() {

		var id int

		var name string
		var nim int
		err := rows.Scan(&id, &name, &nim)

		if err != nil {
			panic(err)
		}

		fmt.Println("id=>", id, "name=>", id, "nim=>", id)

	}
	defer rows.Close()

	fmt.Println("sukses insert ke database")

}

// ini adalah contoh penggunaan semua tipe data untuk iterasi data
func TestSelectAllTipeData(t *testing.T) {

	// ini adalah contoh ketika kita meliterasi data di dalam database dengan tipe data yang berbeda beda
	type Person struct {
		Id, Nim  int64
		Name     string
		Rating   float32
		Birthday time.Time
		Married  bool
	}

	conn := GetConnection()
	conn.Close()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()
	rows, err := conn.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var persons []Person
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Id, &person.Name, &person.Nim, &person.Rating, &person.Birthday, &person.Married)

		if err != nil {
			panic(err)
		}
		persons = append(persons, person)
	}
	defer rows.Close()

	fmt.Println("sukses insert ke database")

	for _, person := range persons {
		fmt.Println("id=>", person.Id, "Name=>", person.Name, "NIM=>", person.Nim, "Birthday=>", person.Birthday, "Rating=>", person.Rating, "Married=>", person.Married)
	}

}

func TestSelectAllTipeDataNullable(t *testing.T) {

	// di golang sendiri itu tidak mengetahui tipe data null itu apa
	// nah untuk meliterasi data yang bersifat nullabel kita bisa mengugakan sql.null_tipedata
	type Person struct {
		Id, Nim  sql.NullInt64
		Name     sql.NullString
		Rating   sql.NullFloat64
		Birthday sql.NullTime
		Married  bool
	}

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()
	rows, err := conn.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Id, &person.Name, &person.Nim, &person.Rating, &person.Birthday, &person.Married)

		if err != nil {
			panic(err)
		}
		if person.Id.Valid {
			fmt.Println("id=>", person.Id.Int64)
		}
		if person.Name.Valid {
			fmt.Println("Name=>", person.Name.String)
		}
		if person.Nim.Valid {
			fmt.Println("NIM=>", person.Nim.Int64)
		}
		if person.Rating.Valid {
			fmt.Println("Rating=>", person.Rating.Float64)
		}
		fmt.Println("Birthday=>", person.Birthday.Time)

		fmt.Println("Married=>", person.Married)
	}
	defer rows.Close()

	fmt.Println("sukses insert ke database")

}
