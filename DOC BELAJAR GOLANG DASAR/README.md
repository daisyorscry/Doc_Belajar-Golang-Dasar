# Doc_Belajar-Golang-Dasar

### Variabel

define variabel di golang ada 2 cara menggunakan var atau :=

```
package main

import "fmt"

func main() {
	// membuat variabel menggunakan kata kunci var
	var name1 string = "Daisy"

	// membuat variabel jauh lebih mudah
	name2 := "Daisy"

	fmt.Println(name1)
	fmt.Println(name2)

}

```

### tipe data

```
string => "this is string"
int => int8, int16, int32, int64
float => float32, float64, complex64, complex128
bool => true or false
array => array := [jerry
jerry
jerry8]int{10,20,30,40}
slice => create_slice := make([]string, panjangnya berapa "2", limit nya berapa "5")
map => person := map[string]string{this data key and value misal "name" : "your name"}
struct => type Person struct{ this data }
constanta => const name string := "your name"

```
### Operasi Tipe Data String
```
package main

import "fmt"

func main() {

	// membuat variabel
	name := "Daisy "

	// fmt package import dari "fmt"
	fmt.Println(name) //result => Daisy
	//len() operasi untuk mengambil index of string. spasi tetap di hitung index
	fmt.Println(len(name)) //result => 6
}
```


###  Tipe Data Constanta
```
package main

import (
	"fmt"
)

func main() {

	// constant adalah tipe data yang tidak bisa dii ubah
	const APPNAME = "GOLANG"
	// const APPNAME = "GOLANG update" //contanta tidak dapat di ubah
	fmt.Println(APPNAME)
}

```
### MAX AND LIMIT Tipe Data int
![MIX & MAX INT](image.png)


###  Type declaration || membuat tipe data sendiri
```
package main

func main() {
	type Person string

	var name Person = "Daisy"
	var address Person = "your address"

	println(name)
	println(address)
}

```

###  Operasi pada tipe data array
```
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


```

###  Operasi pada tipe data slice
```
package main

import "fmt"

func main() {
	// membuat slice dari array

	day := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturyday", "sunday"}
	// mengambil data array dan menjadikan sebagai slice nama_array[mulai dari mana : batas akhirnya  dimana]
	slice1 := day[3:5]
	fmt.Println(slice1) //result => [thursday friday]

	// mengambil dari awal dan nilai batasnya di tentukan
	slice2 := day[:3]
	fmt.Println(slice2) //result => [monday tuesday wednesday]

	// mengambil dari batas yang ditentukan dan batas akhir tidak tentukan
	slice3 := day[3:]
	fmt.Println(slice3) //result => [thursday friday saturyday sunday]

	// mengambil semua
	slice4 := day[:]
	fmt.Println(slice4)

	// mmengubah data array menggunakan slice
	slice5 := day[3:]
	slice5[0] = "thursday_update"
	slice5[1] = "friday_update"

	fmt.Println(slice5) //result => [thursday_update friday_update saturyday sunday]

	// menambahkan data di slice
	// di slice kalau kapasitasnya sudah penuh maka akan dibutkan slice yang baru
	slice6 := append(slice5, "day_ditambahkan")
	fmt.Println(slice6) //result => [thursday_update friday_update saturyday sunday day_ditambahkan]

	// slice bisa mengubah data di array tetapi kalau di tambahkan data dari array sebelumnya tidak akan muncul karena telah dibuatkan slice yang baru
	fmt.Println(day) // => [monday tuesday wednesday thursday_update friday_update saturyday sunday]

	// kata kunci make(tipe_data, panjang_slice, kapasitas dari slicenya) untuk membuat slice baru
	new_slice := make([]int, 2, 10)
	new_slice[0] = 10
	new_slice[1] = 20
	fmt.Println(new_slice) //result => [10 20]

	// melihat panjang dari slicenya
	fmt.Println(len(new_slice))

	// melihat kapasitas dari slicenya
	fmt.Println(cap(new_slice))

	// menambahkan slice
	// ERROR
	// kalau seperti ini akan error karena panjang dari slicenya sudah penuh dan harus menggunakan append
	// new_slice[2] = 30

	new_slice = append(new_slice, 30)
	fmt.Println(new_slice) // result [10 20 30]

}

```

### tipe data map di golang
```
package main

import "fmt"

func main() {

	// membuat tipe data map map([type_key] type_value)
	person := map[string]string{
		"name":    "Daisy",
		"address": "your_address",
		"country": "indonesia",
	}
	// mengambil semua data di map
	fmt.Println(person) //result => map[address:your_address country:indonesia name:Daisy]

	// mengambil salah satu data dari map menggunakan key
	fmt.Println(person["name"]) // resukt => Daisy

	// mengubah data di map
	person["name"] = "daisy_update"
	fmt.Println(person["name"]) // result => daisy_update

	// mengambil panjang data di map
	fmt.Println(len(person)) // result => 3

	// menghapus data di map menggunakan key
	delete(person, "name")
	fmt.Println(person) //result => map[address:your_address country:indonesia]

}


```

### if else expression 

```
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

```

### switch

```
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
```

### for looping
```
package main

import "fmt"

func main() {

	// membuat perulangan dari 0 hingga 100
	// angka++ adalah increment , artinya menaikkan 1 persatu nilanya
	// angka-- adalah dicrement , artinya menurunkan 1 persatu nilainya
	// for define variabel; kondisi; increment || dicrement
	for angka := 0; angka < 100; angka++ {
		fmt.Println(angka)
	}

}

```

### for each

```
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

```


### continue and break

```
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

```

### function 

```
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

```


### struct

```
package main

import "fmt"

func main() {

	Create_Struc() //call function create strunc, result => {Daisy Tangsel Banten Indonesia}

	// nah ketika kita membuat objek baru sekarang kita memiliki function dari structnya yang biasa disebut method
	person := Person{Name: "Daisy_Struct"} //create object Struct and print this name

	fmt.Println(person.Data_Sctruct())

}

// struc adalah prototype dari data kita

// karena golang buka object oriented golang menyediakan fitur struct ini bagai representasi dari data kita
type Person struct {
	Name, City, Province, Country string
}

// kita bisa membuat data untuk structnya seperti ini
func Create_Struc() {
	daisy := Person{
		Name:     "Daisy",
		City:     "Tangsel",
		Province: "Banten",
		Country:  "Indonesia",
	}
	// nah sekarng kita mempunyai field baru yaitu function Data_Struct
	fmt.Println(daisy.Data_Sctruct())
	fmt.Println(daisy)
}

// selanjutnya ada method
// di golang ketika membuat function dan ingin memasukkan data dari struc itu disebut struct
// untuk implementasinya deklarasi struc sebelum nama functionnya

func (person Person) Data_Sctruct() string {
	return "hello " + person.Name
}

```

### interface 


```
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

```

### pointer

```
package main

import "fmt"

type Pointer struct {
	City, Province, Country string
}

func main() {

	// PASS BY VALUE
	// kalau kita reference maka variabel akan di copy tidak akan reference, artinya pointer 1 memiliki data yang berbeda dengan pointer 2

	// ------------------- CASE 1 => POINTER -----------------------------

	// pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	// pointer2 := pointer1
	// fmt.Println(pointer1)

	// // kita telah melakukan reference dan kemudian kita merubah datanya di pointer 2 yang reference ke pointer 1
	// // liat yang terjadi data di pointer 1 tidak akan berubah
	// // hal ini dikrenakan golang mengcopy strucnya menjadi struct yang baru
	// // artinya pointer 1 dan pointer 2 memiliki data yang berbeda

	// pointer2.City = "tangsel"
	// pointer2.Province = "banten"
	// pointer2.Country = "indonesia2"
	// fmt.Println(pointer2)
	// fmt.Println(pointer1)

	// ------------------- CASE 2 => POINTER -----------------------------

	// Menangani kondisi variabel pointer 1 tidak memiliki data yang sama dengan variabel pointer 1
	// POINTER => using simbol & to pointer pointer 1
	// pointer di maksudkan agar variabel ke 2 yang di reference memiliki data yang sama dengan pointer 1

	// pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	// pointer2 := &pointer1 //=> gunakan pointer & agar variabel pointer 1 dan variabel pointer 2 dapat memiliki data yang sama
	// fmt.Println(pointer1)

	// // perhatikan ketika kita mengubah datanya dan melakukan print ke variabel pointer 1 dan variabel pointer 2 maka datanya akan sama
	// pointer2.City = "tangsel"
	// pointer2.Province = "banten"
	// pointer2.Country = "indonesia2"
	// fmt.Println(pointer2)
	// fmt.Println(pointer1)

	// ---------------------OPERATOR ASTERIS *--------------------------------

	// ketika variabel pointer 2 adalah pointer dari variabel pointer 1 maka kita tidak bisa membuat objek dari di variabel yang sudah di pointer
	// misalnya kita membuat lagi seperti ini maka akan error 'pointer2 = Pointer{"kendari", "sulawesi_tenggara", "indonesia"}'
	// gunakan pointer lagi untuk mengatasi hal ini agar pointer 2 memiliki objek yang berbeda dengan pointer 1 'pointer2 = &Pointer{"kendari", "sulawesi_tenggara", "indonesia"}'

	// ------------------- CASE 3 => POINTER -----------------------------

	// pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	// pointer2 := &pointer1

	// pointer2.City = "tangsel"
	// pointer2.Province = "banten"
	// pointer2.Country = "indonesia2"

	// fmt.Println(pointer2)

	// fmt.Println(pointer1)
	// pointer2 = &Pointer{"kendari", "sulawesi_tenggara", "indonesia"} //GUNAKAN POINTER LAGI AGAR variabel pointer2 memiliki objek baru tetapi hal ini tidak membuat pointer 1 merubah nilainy
	// fmt.Println(pointer2)
	// fmt.Println(pointer1)

	// ------------------- CASE 4 => MENGGUNAKAN OPERATOR ASTERIS -----------------------------

	// misalnya kita mau membuat semua yang tadinya mengacu pada pointer 1 akan berubah semua dataya maka kita bisa menggunakan operator asteris
	pointer1 := Pointer{"kendari", "sulawesi_tenggara", "indonesia"}

	pointer2 := &pointer1

	pointer2.City = "tangsel"
	pointer2.Province = "banten"
	pointer2.Country = "indonesia2"

	fmt.Println(pointer1)
	fmt.Println(pointer2)

	*pointer2 = Pointer{"kendari", "sulawesi_tenggara", "indonesia"} //gunakan operator asteris agar pointer 1 mengacu pada pointer 2 agar pointer 1 dan 2 memiliki objek yang sama
	// siapapun yang akan mengacu pada pointer 1 maka akan dirubah datanya menggunakan asteris ke objek yang baru ke pointer 2
	fmt.Println(pointer1)
	fmt.Println(pointer2)

}

```


### pointer function

```
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// saat membuat objek sebagai parameter, maka sebenarnya objeknya di copy,dan tidak akan merubah objek aslinya sebelum di pointer
// misal kita melakukan update untuk data dari strucnya menggunakan parameter struct

// ini tidak akan merubah data dari strucnya karena ketika kita menggunakan struct sebagai parameter maka structnya akan di copy sehingga datanya tidak akan sama
// func ChanceAddress(address Address) {
// 	address.City = "tangsel"
// }

// func main() {

// 	fix_address := Address{"kendari", "sulawesi_tenggara", "indonesia"}
// 	ChanceAddress(fix_address)
// 	fmt.Println(fix_address)
// }

// gunakan pointer agar datanya bisa mengacu pada data dari strucnya
func ChanceAddress(address *Address) {
	address.City = "tangsel"
	address.Province = "banten"
	address.Country = "indonesia_update"
}

func main() {

	fix_address := Address{"kendari", "sulawesi_tenggara", "indonesia"}
	ChanceAddress(&fix_address)
	fmt.Println(fix_address)
}


```

### pointer struct 

```
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// saat membuat objek sebagai parameter, maka sebenarnya objeknya di copy,dan tidak akan merubah objek aslinya sebelum di pointer
// misal kita melakukan update untuk data dari strucnya menggunakan parameter struct

// ini tidak akan merubah data dari strucnya karena ketika kita menggunakan struct sebagai parameter maka structnya akan di copy sehingga datanya tidak akan sama
// func ChanceAddress(address Address) {
// 	address.City = "tangsel"
// }

// func main() {

// 	fix_address := Address{"kendari", "sulawesi_tenggara", "indonesia"}
// 	ChanceAddress(fix_address)
// 	fmt.Println(fix_address)
// }

// gunakan pointer agar datanya bisa mengacu pada data dari strucnya
func ChanceAddress(address *Address) {
	address.City = "tangsel"
	address.Province = "banten"
	address.Country = "indonesia_update"
}

func main() {

	fix_address := Address{"kendari", "sulawesi_tenggara", "indonesia"}
	ChanceAddress(&fix_address)
	fmt.Println(fix_address)
}

```

### defer panic and recover

```
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

```