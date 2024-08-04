package main

import (
	"fmt"
	"sync"
	"testing"
)

// ini adalah In dan Out Channel
// terkadang kita ingin membuat sebuah function yang hanya dapat mengirim dan hanya dapat menerima data
// untuk melakukan itu kita perlu yang namanya in dan out channel
// untuk membuatnya kita perlu menambahkan anotasi <- sesudah kata chan di parameter function untuk menyatakan bahwa function ini hanya dapat mengirim data
// dan anotasi <- sebelum kata chan untuk function yang hanya dapat menerima data
func TestInOutChannel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	// seperti biasa buat dulu channelnya
	channel2 := make(chan string)

	// nah kita panggil function in dan outnya
	go func() {
		defer wg.Done()
		InChannel(channel2)
	}()

	go func() {
		defer wg.Done()
		OutChannel(channel2)
	}()

	// tunggu selama 5 detik sebelum semuanya di tutup untuk memastikan bahawa semua goroutinesnya sudah berjalan
	wg.Wait()
	close(channel2)
}

// buat funntion yang hanya dapat mengirim data ke channel
// nah liat sesudah kata chan tambahkan anotasi <- kemudian tipe_datanya
func InChannel(channel2 chan<- string) {

	// kirim datanya ke string sesuai dengan tipe datanya
	channel2 <- "DATAd"
}

// nah ini adalah function yang hanya dapat menerima data channel
// liat cara pembuatannya gunakan anotas <- sebelum kata chan kemudian tipe_datant
func OutChannel(channel2 <-chan string) {

	// buat sebuah goroutines untuk menerima datanya
	data := <-channel2

	// tambpilkan data yang sudah dikirikan
	fmt.Println(data)
}
