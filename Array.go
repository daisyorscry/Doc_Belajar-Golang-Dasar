package main

import "fmt"

func main() {

	array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array)      //print semua array
	fmt.Println(array[0])   //mengambil array by index
	fmt.Println(len(array)) //mendapatkan panjang array

	// mengubah nilai array

	array[0] = 10
	fmt.Println(array[0]) //mengambil nilai aarray yang telah di ubah

}
