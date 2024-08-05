package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Product struct {
	sync.Mutex
	name, product_name string

	quantity int
}

func (product *Product) Lock() {
	product.Mutex.Lock()

}

func (product *Product) Unlock() {
	product.Mutex.Unlock()
}

func (product *Product) Chance_Quantity(chance_quantity int) {
	product.quantity = product.quantity + chance_quantity
}

// Deadlock terjadi ketika dua atau lebih goroutine saling menunggu untuk resource yang sedang di-lock oleh yang lain. Dalam skenario ini, kita memiliki dua produk dan dua pengguna yang melakukan transaksi secara bersamaan. Mari kita jelaskan langkah-langkah terjadinya deadlock menggunakan animasi:
//
// Inisialisasi Produk dan Pengguna:
//
// user1 memiliki produk laptop dengan kuantitas 1000.
// user2 juga memiliki produk laptop dengan kuantitas 1000.
// Mulai Transaksi:
//
// Transaksi pertama dilakukan oleh user1 yang mengunci user1 kemudian user2.
// Transaksi kedua dilakukan oleh user2 yang mengunci user2 kemudian user1.

// Langkah-langkah Deadlock

// Transaksi 1 (user1 mengunci user1 kemudian mencoba mengunci user2):
//
// user1 mengunci resource user1.
// user1 mencoba mengunci resource user2, namun user2 sudah terkunci oleh transaksi 2.
// Transaksi 2 (user2 mengunci user2 kemudian mencoba mengunci user1):
//
// user2 mengunci resource user2.
// user2 mencoba mengunci resource user1, namun user1 sudah terkunci oleh transaksi 1.

// Animasi Terjadinya Deadlock
// Step 1:
//
// Transaction 1 oleh user1 mengunci user1.
// Transaction 2 oleh user2 mengunci user2.
// Step 2:
//
// Transaction 1 oleh user1 mencoba mengunci user2 (namun user2 sudah terkunci oleh Transaction 2).
// Transaction 2 oleh user2 mencoba mengunci user1 (namun user1 sudah terkunci oleh Transaction 1).
// Step 3:
//
// Kedua transaksi menunggu satu sama lain untuk melepaskan resource, menyebabkan deadlock.
// Kesimpulan
// Deadlock terjadi karena masing-masing goroutine mengunci resource yang berbeda dan kemudian mencoba mengunci resource yang sedang dipegang oleh goroutine lainnya, menyebabkan saling menunggu tanpa akhir. Dengan time.Sleep, kita memperpanjang waktu yang dihabiskan untuk mengunci resource, membuat kondisi deadlock lebih mudah terjadi.

func Transaction(user1 *Product, user2 *Product, amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	user1.Lock()
	fmt.Println("di lock oleh", user1.name, "akan merubah product", user1.product_name)
	user1.Chance_Quantity(-amount)

	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("di lock oleh", user2.name, "akan merubah product", user2.product_name)
	user2.Chance_Quantity(amount)

	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {

	user1 := Product{
		name:         "daisy1",
		product_name: "labtob",
		quantity:     1000,
	}

	user2 := Product{
		name:         "daisy2",
		product_name: "labtob",
		quantity:     1000,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// ini adalah skenario dimana terjadi deadlock
	// disini ada 2 locking yang dilakukan oleh mutex
	// nah ketika locking di lakukan ternyata ada proses lagi yang melakukan locking terhadap data yang mau di locking peroses sebelumya
	// artinya 2 goroutines ini saling melakukan locking proses
	// akibatnya kedua goroutines ini saling menunggu proses satu sama lain yang lagi melakukan locking
	// hal ini mengakibatkan deadlock dan aplikasinya akan mmenunggu terus sampai akhirnya nanti aplikasinya melakukan paniic untuk menghentikan aplikasinya
	go Transaction(&user1, &user2, 20, &wg)
	go Transaction(&user2, &user1, 30, &wg)

	wg.Wait()

	fmt.Println("name =>", user1.name, "productName =>", user1.product_name, "quantity =>", user1.quantity)
	fmt.Println("name =>", user2.name, "productName =>", user2.product_name, "quantity =>", user2.quantity)
}
