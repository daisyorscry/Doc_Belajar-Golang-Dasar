package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ContextSignalDeadline(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			// ketika signal contexnya di kirim maka ctx.done akan di eksekusi dan goroutinesnya akan berhenti
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++

				// ini adalah simulasi ketika proses dari goroutinesnya lambat
				time.Sleep(1 * time.Second)
			}

		}
	}()
	return destination
}

func TestContextSignalDeadline(t *testing.T) {

	fmt.Println(runtime.NumGoroutine())

	// disini kita menambahkan sebuah deadline
	// contex with deadline itu adalah sebuah contex yang akan menjalankan goroutines tetapi waktu prosesnya kita tentukan secara otomatis
	// misal prosesnya kita mau berhenti di jam 12 siang
	// nah maka ketika goroutinesnya berjalan di jam 12 siang goroutinesnya akan otomatis mati
	// nah disini kita membuat context yang akan membawa signal cancel ketika goroutinesya mau kita close
	// dengan menggunakan cancel ini jauh lebih baik

	// buat sebuah contex
	ctx := context.Background()

	// contex with timeout ini me return 2 objek ctx sebagai interface dan cancel sebagai func dan memiliki paramter contex dan timeout
	// kirim contextnya ke dalam function dan kirim berapa lama proses yang harus di jaalankan goroutines
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5))

	// panggil cancelnya di bawah agar ketika aplkasinya telah dijalanakn dan sebelum menutup aplikasinya kirim signal cancel
	// ini yang otomatis akan di eksekusi ketika programnya berhenti dan emngirim singnal cancel kedalam contextnya
	defer cancel()

	destination := ContextSignalDeadline(ctx)
	for n := range destination {
		fmt.Println("counter =>", n)
		if n == 10 {
			break
		}

	}

	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())

}
