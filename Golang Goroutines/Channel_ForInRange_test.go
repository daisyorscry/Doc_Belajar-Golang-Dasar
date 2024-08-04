package main

import (
	"fmt"
	"testing"
	"time"
)

// ini adalah test untuk membuat for in range kepada data channel
// terkadang kita ingin mengirim data kedalam channelnya dalam jumlah yang banyak
// nah daripada kita menggunakan seperti ini data := channel2 intuk menerima datanya dan ketika datanya banyak kita membuat seperti itu sebanyak datannya
// for in range dapat digunakan untuk meliterasi sebanyak apapun data di dalam channelnya
func TestRangeChannel(t *testing.T) {

	// seperti biasa buat dulu channelnya
	// nah setelah deklarasi tipe data ada yang namanya buffered channel, singkattnya itu adalah tempat untuk menampung datanya sebanyak buffednya
	// ketika kita membuat data kedalam channel dan channel itu tidak memiliki buffered maka ketika tidak ada goroutines yang menangkap data itu maka akan terjadi panic dan goroutinesnya akan di close
	// untuk mengatasi hal itu kita bisamenggunakan buffred channel untuk menampung sementara datanya yang di kirim dari channel yang kemudia ketika ada goroutines yang mengambil data channelnya maka dia akan mengambil di dalam buffred channelnya
	ch := make(chan int, 20)

	// okei buat sebuah anonim function untuk mengirim data 1 per satu sebanyak 10 data
	go func() {
		for i := 0; i < 10; i++ {
			// data yang di masukkan berupa integer sebanyak perulangan yang di lakukan
			ch <- i
			// print ini ketika datanya telah dikirim
			fmt.Println("data terkirim =>", i)
			// tunggu 1 seconn setelah datanya dikriim
			time.Sleep(time.Second)
		}
		// jangan lupa close channelnya setelah datanya selesai dikriim
		close(ch)
	}()

	// nah untuk membaca 1 per satu data yang  dikirim gunakan for in range
	// ini memudahkan kita membaca data yang dikirimkan apa aja dari channelnya
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
