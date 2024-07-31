package main

import "fmt"

// secara default golang akan menjalanakan 1 function utama yaitu func main()
// ini adalah func awal yang akan di jalankan golang
func main() {

	Hello_World()                             //call function HelloWorld
	Parameter("aku adalah paramter")          // call function with parameter
	fmt.Println(Func_ReturnValue())           // call function with return value
	name, age := Func_Multiple_Return_Value() // call function with named return value
	fmt.Println(name, age)

	firstname, middlename, lastname, age := Named_Multipel_Return_Value() // call function with return named value
	fmt.Println(firstname, middlename, lastname, age)

	fmt.Println(Var_Argh("daisy1", "daisy2", "daisy3")) //call function with variabel arguments

	function := Func_Variabel
	fmt.Println(function("daisy")) //call function menggunakan variabel

	fmt.Println(fun_as_parameter("daisy", paramter_filter)) // call function yang dijadikan sebagai parameter

	fmt.Println(func_Type_declare("daisy", paramter_filter_type_declare)) // call function yang dijadikan sebagai parameter

	filter := func(name string) string {
		return "mr " + name
	}

	fmt.Println(fun_anonymous("daisy", filter))
}

// kita bisa membuat function baru dengan menggunakan kata kunci func diikuti dengan nama functionya dan kurung buka dan tutup
// kemudian panggil nama functionya di function main agar golang bisa menjalankan function tersebut
// untuk memanggil function bebas berapapun dan tanpa batas
func Hello_World() {
	fmt.Println("print")
}

// FUNCTION Parameter
// paramter adalah jalan masuk data dari luar function
// untuk membuatnya cukup masukkan nama variabel beserta tipe datanya di dalam kurung buka dan tutup function tersebut
// secara langsung ketika kita memanggil function itu maka wajib untuk menambahkan nilai yang di minta oleh paramternya
func Parameter(paramter string) {
	fmt.Println("ini adalah functionparamter =>", paramter)
}

// function return value
// ketika kita membuatf function terkadang kita ingin mengembalikan nilai tersebut ke dalam functionnya sendiri
// ini dapat dilkukan dengan menetapkan tipe data yang kita inginkan setelah membuat nama ffunctionnya

func Func_ReturnValue() string {
	return "aku adalah function return value"
}

// kita bisa membuat banyak return, tinggal buat aja kembaliannya dengan tipe data yang berbeda-beda
func Func_Multiple_Return_Value() (string, int) {
	return "hello aku adalah daisy dan umur aku", 18
}

// ini adalah cara untuk membuat named di return value function
func Named_Multipel_Return_Value() (firstname, middlename, lastname string, age int) {
	firstname = "daisy1"
	middlename = "daisy2"
	lastname = "daisy3"

	age = 18

	return firstname, middlename, lastname, age
}

// ketika kita mmembuat function ada masalah dimana kita tidak mau tau berapa banyak parameter yang dikirimkan oleh function
// kita bisa menggunakan variabel argumen untuk menngani hal ini
func Var_Argh(name ...string) string {
	var result string

	for _, value := range name {
		result += fmt.Sprintln("nama saya adalah name =>", value)
	}
	return result
}

// kita bisa menyimpan function di sebuah variabel jadi ketika kita memanggil sebuah functionya kita panggil menggu
func Func_Variabel(name string) string {
	return "hello" + name
}

// kita juga bisa menggunakan function untuk di jadikan parameter di function lain
func fun_as_parameter(name string, filter func(string) string) string {
	return filter(name)

}

func paramter_filter(name string) string {
	return "mr " + name
}

// jika seperti di atas terlalu rumit untuk menggunakan parameter sebagai function kita bisa menggunakan type declare untuk membuat functionnya seperti membuat objek

type Filter func(string) string

func func_Type_declare(name string, filter Filter) string {
	return filter(name)

}

func paramter_filter_type_declare(name string) string {
	return "mr " + name
}

// sebanarny menggunakan function sebagai parameter digunakan hanya untuk function yang singkat saja atau sederhana
// kita bisa menggunakan anonymous function untuk melakukan hal ini

type Anonymous func(string) string

func fun_anonymous(name string, anonymous Anonymous) string {
	return anonymous(name)

}
