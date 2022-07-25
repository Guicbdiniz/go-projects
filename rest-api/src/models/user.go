package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (user User) SaveUser(db *sql.DB) error {
	q := `INSERT INTO users (username, password) VALUES ($1, $2)`

	_, err := db.Exec(
		q,
		user.Username,
		user.Password,
	)

	return err
}

func ReadUserByUsername(username string, db *sql.DB) (User, error) {
	q := `SELECT id, username FROM users WHERE username=$1`

	rows, err := db.Query(q, username)

	if err != nil {
		return User{}, err
	}

	isThereResult := rows.Next()

	if !isThereResult {
		return User{}, fmt.Errorf("no user with username %s was found", username)
	}

	var selectedUsername string
	var id int
	err = rows.Scan(&id, &selectedUsername)

	if err != nil {
		return User{}, err
	}

	user := User{ID: id, Username: selectedUsername, Password: ""}

	return user, nil
}

func ReadAllUsers(db *sql.DB) ([]User, error) {
	users := make([]User, 0)

	rows, err := db.Query("SELECT id, username FROM users;")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var username string
		var id int
		if err := rows.Scan(&id, &username); err != nil {
			return users, err
		}
		users = append(users, User{ID: id, Username: username, Password: ""})
	}

	if err := rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

// func (user User) UpdateUser(db *sql.DB) error {}

// func DeleteUserByUsername(username string, db *sql.DB) error {}
