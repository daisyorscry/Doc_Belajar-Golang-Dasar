package main

import "fmt"

func main() {
	if true {
		fmt.Println("this is true")
	} else if false {
		fmt.Println("this is false")
	} else {
		fmt.Println("unknown")
	}

}
