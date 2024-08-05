package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// ini adalah pool
// penggunaannya mirip dengan array atau slice, tetapi ketika kita mengambil data dari poolnya kita harus mengambalikann datanya
// di pool udah aman sama yang namanya race condition
func TestPool(t *testing.T) {

	// ini adalah cara memnbuat pool/ deklarasi valiabel pool yang bernilai sync pool
	// var pool sync.

	// ini adalah cara dimana kita membuat sebuah interface kosong sebagai data ddefault ketika datanya tidak ada
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}

	// menggunakan pool.put untuk memasukkan data kedalam pool
	pool.Put("Data_1")
	pool.Put("Data_2")
	pool.Put("Data_3")
	pool.Put("Data_4")

	// ini adalah contoh perulangan dimana goroutines mencoba mengambil data dari pool satu persatu
	for i := 0; i < 10; i++ {

		// buat goroutines
		go func() {
			// ini adalah cara dimana goroutines mengambil data dari pool
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)

			// ketika datanya sudah di ambil maka kita harus mengamballikan datanya kedalam poolnya lagi
			pool.Put(data)

		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("selesai")

}
