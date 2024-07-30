package main

import "fmt"

type Struc struct {
	Name string
}

func (struc *Struc) Chance_name() {
	struc.Name = "nama ditambahkan " + struc.Name
}

func mainnn() {
	nama := Struc{"jerry"}
	fmt.Println(nama)
	nama.Chance_name()
	fmt.Println(nama.Name)

}
