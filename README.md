# Doc_Belajar-Golang-Dasar

### Variabel

define variabel di golang ada 2 cara menggunakan var atau :=

```
var name = "your name" 

name := "your name"
```

### tipe data

```
string => "this is string"
int => int8, int16, int32, int64
float => float32, float64, complex64, complex128
bool => true or false
array => array := [8]int{10,20,30,40}
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
	name := "jerry "

	// fmt package import dari "fmt"
	fmt.Println(name) //result => jerry
	//len() operasi untuk mengambil index of string. spasi tetap di hitung index
	fmt.Println(len(name)) //result => 6
}
```

![MIX & MAX INT](image.png)




