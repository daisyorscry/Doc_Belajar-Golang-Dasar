package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ContextSignalTimeout(ctx context.Context) chan int {
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

func TestContextSignalTimeout(t *testing.T) {

	// disini kita menambahkan sebuah timeout
	// contex with timeout itu adalah sebuah contex yang akan menjalankan goroutines tetapi waktu prosesnya kita tentukan
	// misal prosesnya lebih lama dari timeout yang kita tentukan maka goroutinesnya akan di cancel
	// nah disini kita membuat context yang akan membawa signal cancel ketika goroutinesya mau kita close
	// dengan menggunakan cancel ini jauh lebih baik

	// buat sebuah contex
	ctx := context.Background()

	// contex with timeout ini me return 2 objek ctx sebagai interface dan cancel sebagai func dan memiliki paramter contex dan timeout
	// kirim contextnya ke dalam function dan kirim berapa lama proses yang harus di jaalankan goroutines
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	fmt.Println(runtime.NumGoroutine())
	destination := ContextSignalTimeout(ctx)

	// panggil cancelnya di bawah agar ketika aplkasinya telah dijalanakn dan sebelum menutup aplikasinya kirim signal cancel
	// ini yang otomatis akan di eksekusi ketika programnya berhenti dan emngirim singnal cancel kedalam contextnya
	defer cancel()

	for n := range destination {
		fmt.Println("counter =>", n)
		if n == 10 {
			break
		}

	}

	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())

}
