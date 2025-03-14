package repository

import (
	"appmusic"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user appmusic.User) (int, error) {
	var userId int
	// 1. Создаём пользователя
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, authors_id) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	err := row.Scan(&userId)
	if err != nil {
		return 0, err
	}

	// 2. Создаём автора
	var authorId int
	authorQuery := fmt.Sprintf("INSERT INTO authors (name, description) VALUES ($1, 'Автоматически созданный автор') RETURNING id")
	row = r.db.QueryRow(authorQuery, user.Name)
	err = row.Scan(&authorId)
	if err != nil {
		return 0, err
	}

	// 3. Привязываем `author_id` к пользователю
	updateQuery := fmt.Sprintf("UPDATE %s SET author_id = $1 WHERE id = $2", usersTable)
	_, err = r.db.Exec(updateQuery, authorId, userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

/*func (r *AuthPostgres) CreateUser(user appmusic.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}*/

func (r *AuthPostgres) GetUser(username, password string) (appmusic.User, error) {
	var user appmusic.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
