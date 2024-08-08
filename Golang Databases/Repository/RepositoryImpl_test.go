package repository

import (
	"context"
	"fmt"
	model "golang_database/Model"
	"testing"
	"time"
)

func TestUserRepository_Insert(t *testing.T) {

	conn := GetConnections()
	defer conn.Close()

	UserRepository := UserRepository{
		DB: conn,
	}

	daisy := model.Users{
		Name: "daisy",
		Nim:  1002230006,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	result, err := UserRepository.Insert(ctx, daisy)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
