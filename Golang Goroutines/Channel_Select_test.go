package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// ini sama aja cuman pakai in out channel dan menjelaskan penggunaan select
func TestSelect(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)
	// coba buat 2 channel dengan buffering 3
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)

	// jan lupa tutup channelnya
	defer close(ch1)
	defer close(ch2)

	// nah ini untuk function yang mengirimkan data
	go func() {
		defer wg.Done()
		InSelectChannel(ch1, ch2)

	}()
	// ini function untuk mmelakukan for in range
	// go OutForInRangeChannel(ch1, ch2)

	// ini function yang melakukan penggunaan select
	go func() {
		defer wg.Done()
		OutSelectChannel(ch1, ch2)

	}()

	wg.Wait()
	// tunggu 20detik sebelum aplikasinya benar benar sudah di tutup
}

func InSelectChannel(ch1 chan<- string, ch2 chan<- string) {
	// buat perulangan untuk mengirim 2 data kedalam channel
	for i := 0; i < 2; i++ {
		ch1 <- "data ke " + strconv.Itoa(i)
		fmt.Println("mengirim data ke channel sukses ch1", i)
		time.Sleep(1 * time.Second)
	}

	// buat perulangan untuk mengirim 2 data kedalam channel
	for x := 0; x < 2; x++ {
		ch2 <- "data ke " + strconv.Itoa(x)
		fmt.Println("mengirim data ke channel sukses chs2", x)
		time.Sleep(1 * time.Second)
	}
}

// coba tangkap datanya mengugnakan for in range
func OutForInRangeChannel(ch1 <-chan string, ch2 <-chan string) {
	for val := range ch1 {
		fmt.Println("diterima", val)
	}

	for val := range ch2 {
		fmt.Println("diterima", val)
	}
}

// nah disini kita mengkap datanya menggunakan select

// select digunakan untuk memproses beberapa channel sekaligus
// karena golang bisa berjalan secara pararel dan countcurrency
// maka ketika kita mengugnakan select siapa yang duluan datanya sampai dia yang duluan di proses
// berbeda dengan for in range yang meliterasi datanya satu per satu
// select dapat memprosesnya secara bersamaan
// penggunaannya mirip aja seperti switch
func OutSelectChannel(ch1 <-chan string, ch2 <-chan string) {

	// saya membuatnya sebagai infinity loop agar semua data yang dikirim dapat di proses

	// definisikan counter == 0 untuk membuat batas perulangan ketika counternya sudah di naikkan
	counter := 0

	// buat infinity loop
	for {
		// case 1 ketika datanya di terima akan di masukkan kedalam goroutines val dan akan di print
		select {
		case val := <-ch1:
			// print datanya
			fmt.Println("data diterima ", val)
			// naikkan counternya ketika datanya berhasil di proses
			counter++

		// case 2 ketika datanya di terima akan di masukkan kedalam goroutines val dan akan di print
		case val := <-ch2:
			// print datanya
			fmt.Println("data diterima ", val)
			// naikkan counternya ketika datanya berhasil di proses
			counter++
		//di select kita juga bisa membuat data default ketika data dari channelnya belum sampai
		// ini memungkinnkan kita seperti membuat animasi loading ketika datanya belum sempat di proses
		default:
			fmt.Println("menunggu data")
			time.Sleep(1 * time.Second)
		}
		// lakukan validasi agar infinity loopnya berhenti ketika counternya sudah di naikkan
		if counter == 4 {
			break
		}
	}
}
