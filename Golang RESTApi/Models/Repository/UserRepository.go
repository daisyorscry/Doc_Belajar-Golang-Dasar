package repository

import (
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	Login(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Register(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error)
	Update(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
}
