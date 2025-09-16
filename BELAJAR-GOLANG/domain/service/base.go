package service

import "fmt"

func Yourname(name string, age int) {
	sayHello, _ := fmt.Println("hello", name, "umur anda:", age)

	fmt.Println(sayHello)
}
