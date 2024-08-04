package main

import (
	"fmt"
	"sync"
	"testing"
)

// ini adalah channel
// channel adalah cara kita mengirim data ke goroutines dan di tangkap data juga oleh goroutinesnya
func TestCreateChannel(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)
	// membuat channel bisa dilakukan dengan menggunakan kata kunci make(chan nama_tip_data)
	channel := make(chan string)

	// lakukan defer agar memastikan bahwa ketika channel sudah selesai di eksekusi maka akan di close channelnya
	defer close(channel)

	// kita kirim channelnya ke function sebagai parameter yang nanti di function itu kita akan masukkan datanya
	go func() {
		defer wg.Done()
		SendDataFromChannel(channel)
	}()
	// ini adalah variabel yang menampung data dan mengeksekusi ketika datanya sudah selesai di tangkap
	data := <-channel
	fmt.Println(data)

	wg.Wait()
}

// ini adalah contoh menggunakan channel sebagai parameter di sebuah function
// cara membuatnya cukup func nama_func(alias_dari_channelnya chan tipe_datanya)
func SendDataFromChannel(channel chan string) {

	// ini adalah cara kita mengirim sebuah data ke dalam channel yang sudah dijadikan paramter
	channel <- "AKU ADALAH DATA CHANNELNYA"

	fmt.Println("sukses mengirim data ke channel")
	// eksekusi ini ketika datanya sudah sukses terkirim ke channel

}
