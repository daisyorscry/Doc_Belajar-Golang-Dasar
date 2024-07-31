package main

import "fmt"

func main() {

	runn(true)
}

// defer adalah sebuah cara dimana kita bisa menjalankan program sebelum aplikais kita benar benar berhenti
func deffer() {
	fmt.Println("app is die.................")

	// recover adalah cara untuk menangkap error yang di hasilkan dari panic
	massage := recover()

	fmt.Println("aplikasi error", massage)
}

func runn(error bool) {
	defer deffer()

	// panic adalah cara kita menghentikan secara paksa aplikasi kita ketika ada error
	if error {
		panic("error")
	}

}
