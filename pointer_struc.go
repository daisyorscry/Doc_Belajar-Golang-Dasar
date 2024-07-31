package main

import "fmt"

type Struct struct {
	Name string
}

// ketika kita membuat method maka sebenarnya data di struct nya akan di copy
// jadi datanya tidak akan sama dengan data struc yang aslinya

// nah misalnya kita melakukan update nama di scruct nya
// namanya tidak akan ter update
// func (struc Struct) Chance_name() {
// 	struc.Name = "update nama" + struc.Name
// }

// func main() {
// 	nama := Struct{"jerry"}
// 	fmt.Println(nama)
// 	nama.Chance_name()
// 	fmt.Println(nama.Name)

// }

// gunakan pointer ke sctrucnya agar datanya mengacu pada struct yang sama
func (struc *Struct) Chance_name() {
	struc.Name = "update nama " + struc.Name
}

func main() {
	nama := Struct{"jerry"}
	fmt.Println(nama)
	nama.Chance_name()
	fmt.Println(nama.Name)

}
