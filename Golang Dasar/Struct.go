package main

import "fmt"

func main() {

	Create_Struc() //call function create strunc, result => {Daisy Tangsel Banten Indonesia}

	// nah ketika kita membuat objek baru sekarang kita memiliki function dari structnya yang biasa disebut method
	person := Person{Name: "Daisy_Struct"} //create object Struct and print this name

	fmt.Println(person.Data_Sctruct())

}

// struc adalah prototype dari data kita

// karena golang buka object oriented golang menyediakan fitur struct ini bagai representasi dari data kita
type Person struct {
	Name, City, Province, Country string
}

// kita bisa membuat data untuk structnya seperti ini
func Create_Struc() {
	daisy := Person{
		Name:     "Daisy",
		City:     "Tangsel",
		Province: "Banten",
		Country:  "Indonesia",
	}
	// nah sekarng kita mempunyai field baru yaitu function Data_Struct
	fmt.Println(daisy.Data_Sctruct())
	fmt.Println(daisy)
}

// selanjutnya ada method
// di golang ketika membuat function dan ingin memasukkan data dari struc itu disebut struct
// untuk implementasinya deklarasi struc sebelum nama functionnya

func (person Person) Data_Sctruct() string {
	return "hello " + person.Name
}
