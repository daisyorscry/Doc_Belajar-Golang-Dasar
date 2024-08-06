package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

// ini adalah contoh select dan menyimpan datanya di struct dan slice kemudian di iterasi menggunakan for in range
func TestSelect(t *testing.T) {

	type Person struct {
		Id, Nim int64
		Name    string
	}

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()
	// time.Sleep(time.Second * 9)
	rows, err := conn.QueryContext(ctx, "SELECT * FROM nam")
	if err != nil {
		panic(err)
	}

	var persons []Person
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Id, &person.Name, &person.Nim)

		if err != nil {
			panic(err)
		}
		persons = append(persons, person)
	}
	defer rows.Close()

	fmt.Println("sukses insert ke database")

	for _, person := range persons {
		fmt.Println("id=>", person.Id, "name=>", person.Name, "nim=>", person.Nim)
	}

}

// ini iterasi menggunakan variabel secara eksplisit
func TestSelectVariabel(t *testing.T) {

	conn := GetConnection()

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

	type Person struct {
		Id, Nim  int64
		Name     string
		Rating   float32
		Birthday time.Time
		Married  bool
	}

	conn := GetConnection()

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()
	// time.Sleep(time.Second * 9)
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
	// time.Sleep(time.Second * 9)
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
		// if person.Birthday.Valid {
		// }
		fmt.Println("Married=>", person.Married)

	}

}
