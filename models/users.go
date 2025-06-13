package models

import (
	"errors"

	"example.com/events-api/db"
	"example.com/events-api/helpers"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email,password) VALUES(?,?)`

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err

}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}
	passwordIsValid := helpers.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
