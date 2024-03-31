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
