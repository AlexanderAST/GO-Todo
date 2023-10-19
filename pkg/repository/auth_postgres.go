package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo"
)

type AuthPostgres struct {
	db *sqlx.DB
}
type AuthService struct {
	repo Authorization
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", USERS_TABLE)
	row := r.db.QueryRow(query, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}
func (r *AuthPostgres) AuthUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", USERS_TABLE)
	row := r.db.QueryRow(query, username)
	erras := row.Scan(&user.Id, &user.Username, &user.Password)
	if erras != nil {
		return user, errors.New("invalid mail or password")
	} else {
		query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", USERS_TABLE)
		row := r.db.QueryRow(query, username, password)
		erras = row.Scan(&user.Id, &user.Username, &user.Password)
		return user, erras
	}
}
