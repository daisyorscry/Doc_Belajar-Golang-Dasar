package main

import "fmt"

// Interface ada sebuah immplementasi dari method-method yang akan kita buat

// jadi semacam abstrak yang nanti akan di implementasikan oleh structs

// buat interfacenya  terlebih dahulu
type Service interface {
	Get_Name() string
}

// kemudian buat struct yang akan mengimplementasikan interface di atas
type Person2 struct {
	Name string
	Age  int
}

type Animal struct {
	Name string
	Age  int
}

// untuk mengimplementasikannya kaitkan structnya kedalam sebuah method yang itu adalah implementasi dari interface
// kemudian masukkan structnya seperti yang telah di bahas di struct
func (person Person2) Get_Name() string {
	return person.Name
}

func (animal Animal) Get_Name() string {
	return animal.Name
}

// misalnya disini mengambil return value dari function yang sudah dibuat dari interace melakukan print terhadap nama yang ada di struct menggunakan function abstrak dari interface
func Hello(value Service) {
	fmt.Println("hello", value.Get_Name())
}

func main() {

	// buat data strucnya
	person := Person2{Name: "Daisy", Age: 18}
	animal := Animal{Name: "Rex", Age: 5}

	// kemudian kirim ke method yang sudah mengimplementasikan func interface
	Hello(person)
	Hello(animal)
}
