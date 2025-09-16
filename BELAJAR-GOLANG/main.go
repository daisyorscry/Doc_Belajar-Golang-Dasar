package main

import (
	"belajar-go-h1/domain/service"
	"belajar-go-h1/pkg/database"
)

func main() {

	name, age := seviceBuildName("name1", 12)
	name1, age1 := seviceBuildName("name2", 13)

	service.Yourname(name, age)
	service.Yourname(name1, age1)

	_ = database.InitMysql()

}

func seviceBuildName(nameParam string, ageParam int) (string, int) {
	name := nameParam
	age := ageParam

	return name, age
}
