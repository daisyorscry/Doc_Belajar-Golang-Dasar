package main

import "fmt"

func main() {

	// ini adalah break digunakan jika kondisi perulanganya sudah terpenuhi dan akan menghentikan perulangannya
	// perulangan akan terhenti jika menemui angka 8
	for i := 0; i < 100; i++ {
		if i == 8 {
			break
		}
		fmt.Println(i)
	}

	// continue adalah mengskip jika kondisinya terpenuhi dan tetap melakukan perulangan
	// mencoba mencetak angka ganjil dari 0 hingga 100
	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

}
