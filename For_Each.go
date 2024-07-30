package main

import "fmt"

func main() {

	// for each atau di golang disebutnya for range adalah teknik membaca data di dalam tipe data collection seperti map, slice dan array
	data := []int{10, 20, 30, 40}

	// misal kalau menggunakan perulangan manual
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

	// ini menggunakan for range
	for keys, values := range data {
		fmt.Println(keys, "=>", values)
	}
}
