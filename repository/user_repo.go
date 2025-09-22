package repository

import (
	"database/sql"
	"gotesting/config"
)

type User struct {
	ID   int
	Name string
}

// GetUserByID returns a user from DB
func GetUserByID(id int) (*User, error) {
	row := config.DB.QueryRow("SELECT id, name FROM users WHERE id=?", id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

// AddUser inserts a new user into DB
func AddUser(user User) error {
	_, err := config.DB.Exec("INSERT INTO users(id, name) VALUES(?, ?)", user.ID, user.Name)
	return err
}
