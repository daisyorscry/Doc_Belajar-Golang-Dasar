package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// saat membuat objek sebagai parameter, maka sebenarnya objeknya di copy,dan tidak akan merubah objek aslinya sebelum di pointer
// misal kita melakukan update untuk data dari strucnya menggunakan parameter struct

// ini tidak akan merubah data dari strucnya karena ketika kita menggunakan struct sebagai parameter maka structnya akan di copy sehingga datanya tidak akan sama
// func ChanceAddress(address Address) {
// 	address.City = "tangsel"
// }

// func main() {

// 	fix_address := Address{"kendari", "sulawesi_tenggara", "indonesia"}
// 	ChanceAddress(fix_address)
// 	fmt.Println(fix_address)
// }

// gunakan pointer agar datanya bisa mengacu pada data dari strucnya
func ChanceAddress(address *Address) {
	address.City = "tangsel"
	address.Province = "banten"
	address.Country = "indonesia_update"
}

func main() {
	fix_address := Address{"kendari", "sulawesi_tenggara", "indonesia"}
	ChanceAddress(&fix_address)
	fmt.Println(fix_address)
}
