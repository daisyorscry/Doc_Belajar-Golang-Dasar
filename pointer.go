package main

import "fmt"

type Pointer struct {
	City, Province, Country string
}

func mainny() {

	// PASS BY VALUE
	// kalau kita reference maka variabel akan di copy tidak akan reference, artinya pointer 1 memiliki data yang berbeda dengan pointer 2

	// pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	// pointer2 := pointer1
	// fmt.Println(pointer1)

	// pointer2.City = "tangsel"
	// pointer2.Province = "banten"
	// pointer2.Country = "indonesia2"
	// fmt.Println(pointer2)
	// fmt.Println(pointer1)

	// POINTER => using simbol & to pointer pointer 1
	// pointer di maksudkan agar variabel ke 2 yang di reference memiliki data yang sama dengan pointer 1

	// pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	// pointer2 := &pointer1
	// fmt.Println(pointer1)

	// pointer2.City = "tangsel"
	// pointer2.Province = "banten"
	// pointer2.Country = "indonesia2"
	// fmt.Println(pointer2)
	// fmt.Println(pointer1)

	// OPERATOR ASTERIS *
	// ketika variabel pointer 2 adalah pointer dari variabel 1 maka kita tidak bisa membuat objek dari di variabel yang sudah di pointer
	// misalnya kita membuat lagi seperti ini maka akan error 'pointer2 = Pointer{"kendari", "sulawesi_tenggara", "indonesia"}'
	// gunakan pointer lagi untuk mengatasi hal ini agar pointer 2 memiliki objek yang berbeda dengan pointer 1 'pointer2 = &Pointer{"kendari", "sulawesi_tenggara", "indonesia"}'

	// pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	// pointer2 := &pointer1

	// pointer2.City = "tangsel"
	// pointer2.Province = "banten"
	// pointer2.Country = "indonesia2"

	// fmt.Println(pointer2)

	// fmt.Println(pointer1)
	// pointer2 = &Pointer{"kendari", "sulawesi_tenggara", "indonesia"} //GUNAKAN POINTER LAGI AGAR variabel pointer2 memiliki objek baru tetapi hal ini tidak membuat pointer 1 merubah nilainy
	// fmt.Println(pointer2)
	// fmt.Println(pointer1)

	// misalnya kita mau membuat semua yang tadinya mengacu pada pointer 1 akan berubah semua dataya maka kita bisa menggunakan operator asteris
	pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	pointer2 := &pointer1

	pointer2.City = "tangsel"
	pointer2.Province = "banten"
	pointer2.Country = "indonesia2"

	fmt.Println(pointer1)
	fmt.Println(pointer2)

	*pointer2 = Pointer{"kendari", "sulawesi_tenggara", "indonesia"} //gunakan operator asteris agar pointer 1 mengacu pada pointer 2 agar pointer 1 dan 2 memiliki objek yang sama
	// siapapun yang akan mengacu pada pointer 1 maka akan dirubah datanya menggunakan asteris ke objek yang baru ke pointer 2
	fmt.Println(pointer1)
	fmt.Println(pointer2)

}
