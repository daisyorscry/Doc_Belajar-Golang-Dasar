package main

import (
	"fmt"
)

func maimn() {

	// learn_slice()
	// learn_if_else_ifelse_expression()
	// learn_forloop()
	// learn_continue()
	// learn_break()
	// learn_bilangan_ganjil_continue()

	// nama, _ := learn_func_parameter("jerry", 18)
	// fmt.Println("nama kamu adalah", nama)
	// fmt.Println("dan umur kamu adalah", umur)

	// firstname, middlename, lastname := learn_named_return_value()
	// fmt.Println(firstname)
	// fmt.Println(middlename)
	// fmt.Println(lastname)

	// sum := variabel_arguments(10, 20, 30, 40, 50)
	// fmt.Println(sum)

	// sum := variabel_arguments_foreach(10, 20, 30, 40, 50)
	// fmt.Println(sum)

	// num := []int{10, 20, 30, 40, 50}
	// sum := variabel_arguments_foreach_with_slice(num...)
	// fmt.Println(sum)

	// variabel_func := function_as_variabel
	// fmt.Println(variabel_func("jerry"))

	// variabel_parameter := function_as_parameter_func
	// fmt.Println(variabel_parameter("jerry", filter_name))

	// filter := func(name string) string {
	// 	return name + "baik"
	// }

	// fmt.Println(filter_name_for_anonymous_func("jerry", filter))

	// closures

	// counter := 0
	// increments := func() {
	// 	fmt.Println("hello")
	// 	counter++
	// }

	// increments()
	// increments()
	// increments()

	// println(counter)

	// deffer, panic
	// run(true)

	// var person Person
	// person.Name = "jerry"
	// person.Address = "jl. amil mena"
	// person.Age = 18

	// println(person.Name)
	// println(person.Address)
	// println(person.Age)

	// person2 := Person{
	// 	Name:    "jerry",
	// 	Address: "JL. AMIL MENA",
	// 	Age:     18,
	// }
	// fmt.Println(person2)

	// person2.sayHello("agus")

	// person3 := Person{
	// 	Name:    "jerry",
	// 	Address: "jll. amil menas",
	// 	Age:     18,
	// }

	// SayHello(person3)

	// data := CreateMap("jerry")
	// fmt.Println(data)

	// result := Random()
	// new_result_string := result.(string)
	// new_result_int := result.(int)

	// fmt.Println(new_result_string)
	// fmt.Println(new_result_int) ==> panic

	// switch value := result.(type) {
	// case string:
	// 	fmt.Println("jadi string", value)
	// case int:
	// 	fmt.Println("jadi int", value)
	// default:
	// 	fmt.Println("unknown")

	// }

	// address1 := Address{"kendari", "sultra", "indonesia"}
	// address2 := address1 //pointer
	// fmt.Println(address1)
	// fmt.Println(address2)
	// address2.City = "tangsel"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// address1 := Address{"kendari", "sultra", "indonesia"}
	// address2 := &address1 //pointer
	// fmt.Println(address1)
	// fmt.Println(address2)
	// address2.City = "tangsel"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// address2 = &Address{"tansel", "banten", "indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2)

	// *address2 = Address{"tansel", "banten", "malaysia"}
	// fmt.Println(address1)
	// fmt.Println(address2)

	// address := &Address{}
	// ChanceAddress(address)

	// fmt.Println(address)

	// man := Man{"jerry"}
	// man.Married()

	// println(man.Name)
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

// func ChanceAddress(address *Address) {
// 	address.Country = "indonesia"

// }

// type Address struct {
// 	City, Province, Country string
// }

func Random() any {
	return "ok"
}

func CreateMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(people People) {

	fmt.Println("hello", people.GetName())
}

type People interface {
	GetName() string
}

type Person struct {
	Name, Address, Country string
	Age                    int
}

func (person Person) sayHello(name string) {
	fmt.Println("hello", name, "my name is", person.Name)
}

func run(error bool) {
	defer loggin()
	if error {
		panic("close program")
	}
}
func loggin() {
	massage := recover()
	fmt.Println("errror aplikasi => ", massage)
	fmt.Println("logging")
}

func filter_name_for_anonymous_func(name string, filter func(name string) string) string {
	return "nama kamu setelah di filter anonymous " + name

}

func filter_name(name string) string {
	return name + " baik"

}

func function_as_parameter_func(name string, filter func(string) string) string {
	return "nama kamu akan di tambahkan " + filter(name)

}

func function_as_variabel(name string) string {
	return "hello nama kamu adalah " + name

}

func variabel_arguments_foreach_with_slice(numbers ...int) int {

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func variabel_arguments_foreach(numbers ...int) int {

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func variabel_arguments(angka ...int) int {
	total := 0
	for i := 0; i < len(angka); i++ {
		total += angka[i]
	}
	return total
}

func learn_named_return_value() (firstname, middlename, lastname string) {
	firstname = "jerry1"
	middlename = "jerry2"
	lastname = "jerry3"

	return firstname, middlename, lastname

}

func learn_func_parameter(name string, umur int) (string, int) {
	return name, umur
}

func learn_slice() {
	values := [...]int{10, 20, 30, 40, 50}

	slice := values[4:]
	fmt.Println(slice)

	slice[0] = 80
	fmt.Println(slice)

	slice2 := append(slice, 90)
	fmt.Println(slice2)
	fmt.Println(values)

	new_slice := make([]string, 2, 5)
	new_slice[0] = "jerry"
	new_slice[1] = "asep"

	fmt.Println(new_slice)

	ini_slice := []int{10, 20, 30, 40}

	fmt.Println(ini_slice)

	// tipe data map
	person := map[string]string{
		"name": "jerry",
		"umur": "18",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["umur"])

	fmt.Println(len(person))
	delete(person, "name")
	fmt.Println(person)
}
func learn_if_else_ifelse_expression() {
	name := "jerry"

	if name != "jerry" {
		fmt.Println(name)
	} else if name == "jerry" {
		fmt.Println(" nama benar")
	} else {
		fmt.Println("nama tidak valid")
	}
}
func learn_forloop() {

	for ganjil := 0; ganjil < 100; ganjil++ {
		if ganjil%2 != 0 {
			fmt.Println("bilangan ganil => ", ganjil)
		}
	}
}

func learn_bilangan_ganjil_continue() {
	for ganjil := 0; ganjil < 100; ganjil++ {
		if ganjil%2 == 0 {
			continue
		}
		fmt.Println(ganjil)
	}
}

func learn_continue() {
	person := []string{"jerry", "asep", "udin"}

	for ganil := 0; ganil < len(person); ganil++ {
		if person[ganil] == "asep" {
			continue
		}
		fmt.Println(person[ganil])
	}
}

func learn_break() {
	person := []string{"jerry", "asep", "udin"}

	for ganil := 0; ganil < len(person); ganil++ {
		if person[ganil] == "asep" {
			break
		}
		fmt.Println(person[ganil])
	}
}
