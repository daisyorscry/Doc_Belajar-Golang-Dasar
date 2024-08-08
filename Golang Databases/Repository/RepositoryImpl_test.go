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

	// ini test untuk insert database
	result, err := UserRepository.Insert(ctx, daisy)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// ini test untuk update database
	result, err = UserRepository.Update(ctx, daisy)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestDeleteDatabase(t *testing.T) {

	conn := GetConnections()
	defer conn.Close()

	UserRepository := UserRepository{
		DB: conn,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// ini test untuk delete database
	deleted, err := UserRepository.Delete(ctx, 1023)
	if err != nil {
		t.Fatal("error saat menghapus data", err)
	}

	if !deleted {
		t.Fatal("data id tidak ditemukan")
	}

}

func TestFindAll(t *testing.T) {

	conn := GetConnections()
	defer conn.Close()

	UserRepository := UserRepository{
		DB: conn,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// testing context
	// time.Sleep(time.Second * 5)

	results, err := UserRepository.FindAll(ctx)
	if err != nil {
		t.Fatal("error saat membaca data", err)
	}

	for _, result := range results {
		fmt.Println(result)
	}

}

func TestFindById(t *testing.T) {

	conn := GetConnections()
	defer conn.Close()

	ctx := context.Background()

	UserRepository := UserRepository{
		DB: conn,
	}

	result, err := UserRepository.FindById(ctx, 1007)
	if err != nil {
		t.Fatal("error find users", err)
	}

	fmt.Println(result)
}
