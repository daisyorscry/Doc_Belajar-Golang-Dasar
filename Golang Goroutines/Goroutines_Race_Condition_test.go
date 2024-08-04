package main

import (
	"fmt"
	"sync"
	"testing"
)

// ini adalah simulasi dimana golang yang berjalan sebagai pararel dan bisa juga berjalan sebagai councurency
// hal ini dapat menjadi masalah dimana dapat terjadi race condition
// dimana race condition adalah ketika banyak goroutines mengakses data secara bersamaan dalam waktu nanoSecon secara bersamaan
// hal ini dapat mengakibatkan ketidakkonsesistennya data tersebut
func TestRaceCondition(t *testing.T) {
	// anggap ini datanya ada 0
	x := 0

	var wg sync.WaitGroup

	nummGoroutines := 100
	wg.Add(nummGoroutines)
	// kita membuat goroutines sebanyak 1k
	for i := 0; i <= nummGoroutines; i++ {
		// buat goroutinesnya menggunakan anonymous func
		go func() {
			defer wg.Done()
			// nah sekarang kita coba membuat 100 goroutines mengakses data yang sama
			for j := 1; j <= 100; j++ {
				// ketikaa gorouitines ini mengubah data ini maka yang terjadi adalah
				// beberapa goroutines dalam waktu nanoSecon yang sama dapat mengakses data yang sama
				// dan ketika goroutines itu merubah maka terjadi ketidakkonsistensi data
				// seharusnya setiap goroutines menambahkan 100 ke x
				// ada 1k gorotutines yang berjalan
				// dan 100 perubahan data akan dijalankan oleh 1k goroutines
				// artinya di data x akan ada 100k data
				// akan tetapi ketika race condition ini berjalan
				// datanya tidakakan konsisten setiap kali test ini berjalan
				x = x + 1
				fmt.Println("counter kess =>", x)
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter", x)
}

// untuk mengatasi hal ini kita bisa melakukan mutex lock untuk setiap perubahan yang akan di lakukan oleh goroutines
func TestResolfRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	numGoroutines := 1000
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// ketika goroutines akan merubah data ini hanya 1 goroutines yang akan masuk ke dalam proses ini dan akan di lock
				// artinya semua goroutines akan menunggu untuk melakuan proses merubah data sampai goroutines yang sudah terlock selesai melakukakn perubahan data
				mutex.Lock()
				x = x + 1
				// fmt.Println("counter kess =>", j)
				// ini adalah anotasi unlock untuk membuka kunci dan menyatakan  bahwa goroutines sudah melakukan perubahan data
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	// hasil dari data x yang diubah oleh goroutines akan konsisten sebanyak 100k
	fmt.Println("counters", x)
}
