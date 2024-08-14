package repository

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error) {
	SQL := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1"

	var user entity.User
	err := tx.QueryRowContext(ctx, SQL, userId).Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, helper.RepositoryErr(err, "user not found")
		}
		return entity.User{}, helper.RepositoryErr(err, "error finding user by id")
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error) {
	SQL := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE username = $1"

	var user entity.User
	err := tx.QueryRowContext(ctx, SQL, username).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, helper.RepositoryErr(err, "user not found")
		}
		return entity.User{}, helper.RepositoryErr(err, "error finding user by username")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	SQL := "SELECT id, username, email, created_at, updated_at FROM users WHERE username = $1 AND password = $2"

	var loggedInUser entity.User
	err := tx.QueryRowContext(ctx, SQL, user.Username, user.Password).Scan(&loggedInUser.Id, &loggedInUser.Username, &loggedInUser.Email, &loggedInUser.CreatedAt, &loggedInUser.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, helper.RepositoryErr(err, "invalid username or password")
		}
		return entity.User{}, helper.RepositoryErr(err, "error during login")
	}

	return loggedInUser, nil
}

func (r *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	SQL := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err := tx.QueryRowContext(ctx, SQL, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return entity.User{}, helper.RepositoryErr(err, "error during registration")
	}

	user.Id = id
	return user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	SQL := "UPDATE users SET username = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5 RETURNING id, username, email, created_at, updated_at"

	err := tx.QueryRowContext(ctx, SQL, user.Username, user.Email, user.Password, user.UpdatedAt, user.Id).Scan(
		&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return entity.User{}, helper.RepositoryErr(err, "error updating user")
	}

	return user, nil
}
