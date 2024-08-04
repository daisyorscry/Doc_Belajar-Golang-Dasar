package main

import "fmt"

func main() {

	// membuat tipe data map map([type_key] type_value)
	person := map[string]string{
		"name":    "Daisy",
		"address": "your_address",
		"country": "indonesia",
	}
	// mengambil semua data di map
	fmt.Println(person) //result => map[address:your_address country:indonesia name:Daisy]

	// mengambil salah satu data dari map menggunakan key
	fmt.Println(person["name"]) // resukt => Daisy

	// mengubah data di map
	person["name"] = "daisy_update"
	fmt.Println(person["name"]) // result => daisy_update

	// mengambil panjang data di map
	fmt.Println(len(person)) // result => 3

	// menghapus data di map menggunakan key
	delete(person, "name")
	fmt.Println(person) //result => map[address:your_address country:indonesia]

}
