package main

import "fmt"

func main() {

	learn_golang := true

	switch learn_golang {
	case true:
		fmt.Println("this is learn golang")
	case false:
		fmt.Println("this is not learn golang")
	default:
		fmt.Println("die")

	}

}
