package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// saat membuat objek sebagai parameter, maka sebenarnya objeknya di copy,dan tidak akan merubah objek aslinya sebelum di pointer
func ChanceAddress(address *Address) {
	address.City = "tangsel"
}

func mainnna() {

	fix_address := Address{"kendari", "sulawesi_tenggara", "indonesia"}
	ChanceAddress(&fix_address)
	fmt.Println(fix_address)
}
