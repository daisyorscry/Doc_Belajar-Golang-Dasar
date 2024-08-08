package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// ini adalah contoh insert di database
func TestInsert(t *testing.T) {

	// import dulu connectionnya kemudian masukkan kedalam variabel
	conn := GetConnection()
	// jangan lupa close connetionya agar tidak terjadi
	conn.Close()

	// buat contex, sebenarnya bisa tanpa contex tetapi ini hanya memanfaatkan fitur dari golang untuk sebuah cara yang lebih bagus
	ctx := context.Background()

	// buat contex timeout agar databaes tau ketika transaksinya berjalan berapa lama timeout transaksinya harus bekerja
	// jika lebih dari batas timeout yang telah di tentukan maka transaksinya akan di cance
	ctx, cancel := context.WithTimeout(ctx, time.Minute*30)

	// ini adalah car kita menutu connetion ketika semua proses telah selesai
	defer cancel()

	for i := 0; i < 9153; i++ {

		// ini adalah cara untuk melakukan eksekusi ke database panggil variabel connection di variabel connetion itu ada beberapa function yang bisa digunakan
		// begin ini untuk memulai transactional
		// conn.Begin()
		// conn.BeginTx()

		// exec digunankan untuk perintal sql yang tidak memiliki kembalian seperti insert update delete
		// conn.Exec()
		// conn.ExecContext()

		// ini adalah untuk membuat prepare statement yaitu ketika kita membuat perintah sql yang sama dan mau menggunakan 1 connection yang sama
		// conn.Prepare()
		// conn.PrepareContext()

		// query ini untuk iterasi data di database
		// conn.Query()
		// conn.QueryContext()

		// kalau query row untuk iterasi data di databse tetapi hanya mengambil 1 data
		// conn.QueryRow()
		// conn.QueryRowContext()

		// exec contex mengembalikan 2 tipe data  yaitu result dan error
		_, err := conn.ExecContext(ctx, "INSERT INTO name(name, nim) values('jerry', 1002230006)")

		// ini adalah cara untuk hadle errornya
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("sukses insert ke database")
}
