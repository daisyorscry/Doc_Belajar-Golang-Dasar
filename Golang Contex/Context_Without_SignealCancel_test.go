package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// ini adalah contoh kalau kita menggunakan goroutines tetapi setiap kita menjalankan gorouitnesnya itu tidak pernah diclose
// artinya akan terjadi memory leak dikarenakan goroutinesnya tidak pernah di close

// ini adalah simulasi dimana kita mengirimkan data melalui channel secara terus terusaan mengugnakan infinity loop
// dan kita akan menutup perulangannya jika counternya sudah naik menjadi 10
// akan tetapi goroutinesnya tidak akan ter close
// jika seperti ini misal ada 1juta request dan goroutinesnya tetap berjalan maka akan memperlambat aplikasi

// buat function yang return channel
func ContextSignalWithoutCancel() chan int {

	// buat  channelnya dan simpan di variabel destination
	destination := make(chan int)

	// buat gorotuinesnya menggunakan anonymous func
	go func() {
		// close channelnya menggunakan defer
		defer close(destination)

		// ini adalah data yang akan dikirimkan melalui channel
		counter := 1

		// lakukan perulangan untuk mengirimkan data ke channel
		for {
			destination <- counter
			counter++ //data akan terus terusan di ubah dengan increment
			time.Sleep(500 * time.Millisecond)

		}
	}()

	// balikkan data channelnya
	return destination
}

func TestContextSignalWithoutCancel(t *testing.T) {

	fmt.Println(runtime.NumGoroutine())

	// buat variabel untuk menampung data channelnya
	destination := ContextSignalWithoutCancel()

	// literasi menggunakan range untuk data dari channelnya
	for n := range destination {
		fmt.Println("counter =>", n)

		// nah jika data yang di kirim sudah mencapai 10 maka akan di break perulangqnnya
		if n == 10 {
			break
		}

	}

	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())

	// nah secara code mungkin kita berffikir bahwa goroutinesnya akan berhenti
	// nyatanya goroutinesnya masih menggantung dan tidak ter close
	// hal ini sangat berbahaya jika banyak goroutines yang berjalan dan tidak di close
	// untuk close goroutinesnya gunakan context signal dengan mengirim signal cancel kedalam contextnya
}
