package repository

import (
	"context"
	model "golang_database/Model"
)

// ini adalah contoh membuat repository pattern di golang

// buat dulu sebuah interface di golang
// parameternya context dan struct yang digunakan sebagai model untuk membuat representasi databasenya, kemudian pastikan mereturn modelnya lagi dan errornya
type User interface {
	Insert(ctx context.Context, user model.Users) (model.Users, error)
	Update(ctx context.Context, user model.Users) (model.Users, error)
	Delete(ctx context.Context, user model.Users) (bool, error)
	FindById(ctx context.Context, id int64) (model.Users, error)
	FindAll(ctx context.Context) ([]model.Users, error)
}
