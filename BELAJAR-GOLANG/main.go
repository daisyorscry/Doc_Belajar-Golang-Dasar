package main

import (
	"fmt"
	"log"
	"net/http"

	"belajar-go-h1/domain/service"
	"belajar-go-h1/pkg/database"
)

func main() {

	name, age := seviceBuildName("name1", 12)
	name1, age1 := seviceBuildName("name2", 13)

	service.Yourname(name, age)
	service.Yourname(name1, age1)

	_ = database.InitMysql()

	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}

}

func seviceBuildName(nameParam string, ageParam int) (string, int) {
	name := nameParam
	age := ageParam

	return name, age
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s!", name)
}
