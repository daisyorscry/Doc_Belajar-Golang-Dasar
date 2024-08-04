package main

import "fmt"

func main() {
	// membuat slice dari array

	day := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturyday", "sunday"}
	// mengambil data array dan menjadikan sebagai slice nama_array[mulai dari mana : batas akhirnya  dimana]
	slice1 := day[3:5]
	fmt.Println(slice1) //result => [thursday friday]

	// mengambil dari awal dan nilai batasnya di tentukan
	slice2 := day[:3]
	fmt.Println(slice2) //result => [monday tuesday wednesday]

	// mengambil dari batas yang ditentukan dan batas akhir tidak tentukan
	slice3 := day[3:]
	fmt.Println(slice3) //result => [thursday friday saturyday sunday]

	// mengambil semua
	slice4 := day[:]
	fmt.Println(slice4)

	// mmengubah data array menggunakan slice
	slice5 := day[3:]
	slice5[0] = "thursday_update"
	slice5[1] = "friday_update"

	fmt.Println(slice5) //result => [thursday_update friday_update saturyday sunday]

	// menambahkan data di slice
	// di slice kalau kapasitasnya sudah penuh maka akan dibutkan slice yang baru
	slice6 := append(slice5, "day_ditambahkan")
	fmt.Println(slice6) //result => [thursday_update friday_update saturyday sunday day_ditambahkan]

	// slice bisa mengubah data di array tetapi kalau di tambahkan data dari array sebelumnya tidak akan muncul karena telah dibuatkan slice yang baru
	fmt.Println(day) // => [monday tuesday wednesday thursday_update friday_update saturyday sunday]

	// kata kunci make(tipe_data, panjang_slice, kapasitas dari slicenya) untuk membuat slice baru
	new_slice := make([]int, 2, 10)
	new_slice[0] = 10
	new_slice[1] = 20
	fmt.Println(new_slice) //result => [10 20]

	// melihat panjang dari slicenya
	fmt.Println(len(new_slice))

	// melihat kapasitas dari slicenya
	fmt.Println(cap(new_slice))

	// menambahkan slice
	// ERROR
	// kalau seperti ini akan error karena panjang dari slicenya sudah penuh dan harus menggunakan append
	// new_slice[2] = 30

	new_slice = append(new_slice, 30)
	fmt.Println(new_slice) // result [10 20 30]

}
