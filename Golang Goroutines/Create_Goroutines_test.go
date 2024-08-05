package main

import (
	"fmt"
	"sync"
	"testing"
)

func CreateGolangGoroutines(number int) {
	fmt.Println("ini adalah goroutines ke =>", number)
}

// ini adalah test untuk membuat goroutines
// go routines berjalan secara countcurrent dan pararel
// artinya akan di masukkan ke dalam antrian untuk kemudian di proses
// ketika prosesnya sudah selesai maka akan di ambil hasil prosesnya
// itu lah mengapa ketika kita menjalankan goroutinesnya melalui test di bawah print-nya dulu yang di eksekusi bukan function yang di tambahkan anotation go untuk membuat goroutines-nya

func TestCreateGolangGoroutines(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		CreateGolangGoroutines(1)
	}()

	wg.Wait()
	fmt.Println("selesai membuat go routines")
}
