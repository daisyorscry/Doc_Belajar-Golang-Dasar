package main

import (
	"fmt"
	"sync"
	"testing"
)

// wait grub adalah struct untuk memastikan bahwa semua goroutines telah berjalan
// penggunaan time.sleep untuk goroutines tidak berjalan karena kita tidak tau kapan gorouitnes itu telah selesai
func TestCreateWaitGrub(t *testing.T) {
	// untuk membuat nya harus menggunakan type var
	var wg sync.WaitGroup

	// di dalam sync.WaitGrub ada 3 func
	// ADD, untuk mententukan berapa goroutines yang berjalan
	// DONE, adalah untuk memberitau bahwa goroutines telah selesai berjalan
	// WAIT, adalah time.sleep dimana goroutines akan di tunggu hingga selesai melakukan pekerjaanny

	// deklarasikan bahwa ada 1 goroutines yang akan berjalan
	wg.Add(1)
	go func() {
		// untuk menutup goroutines ketika seleesai berjalan panggil func DONE di deffer
		fmt.Println("helllo aku adalah goroutines")
		defer wg.Done()
	}()

	// gorutines akan ditunggu hingga selesai melakukan pekerjaan
	wg.Wait()

}
