package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ContextSignalCancel(ctx context.Context) chan int {
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
				time.Sleep(500 * time.Millisecond)
			}

		}
	}()
	return destination
}

func TestContextSignalCancel(t *testing.T) {

	// nah disini kita membuat context yang akan membawa signal cancel ketika goroutinesya mau kita close
	// dengan menggunakan cancel ini jauh lebih baik

	// buat sebuah contex
	ctx := context.Background()

	// contex with cancel ini me return 2 objek ctx sebagai interface dan cancel sebagai func
	// kirim contextnya ke dalam function
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println(runtime.NumGoroutine())
	destination := ContextSignalCancel(ctx)

	for n := range destination {
		fmt.Println("counter =>", n)
		if n == 10 {
			break
		}

	}

	// panggil cancelnya di bawah agar ketika aplkasinya telah dijalanakn dan sebelum menutup aplikasinya kirim signal cancel
	// ini yang otomatis akan di eksekusi ketika programnya berhenti dan emngirim singnal cancel kedalam contextnya
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())

}
