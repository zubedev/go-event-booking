package models

import (
	"go-event-booking/db"
	"go-event-booking/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() (*User, error) {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	exec, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	id, err := exec.LastInsertId()
	user.ID = id
	return user, err
}

func (user *User) ValidateCredentials() (*User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	storedUser := User{}
	err := row.Scan(&storedUser.ID, &storedUser.Email, &storedUser.Password)
	if err != nil {
		return nil, err
	}

	err = utils.VerifyPassword(storedUser.Password, user.Password)
	if err != nil {
		return nil, err
	}

	return &storedUser, nil
}
