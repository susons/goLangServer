package models

import (
	"errors"
	"fmt"

	"example.com/project_api/db"
	"example.com/project_api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	defer statement.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	result, err := statement.Exec(u.Email, hashPassword)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	userId, err := result.LastInsertId()

	u.Id = userId

	return err
}

func (u *User) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var passwordHash string

	err := row.Scan(&u.Id, &passwordHash)

	if err != nil {
		fmt.Println("Error: ", err)
		return errors.New("credentials invalid")
	}

	passwordValid := utils.ValidatePassword(u.Password, passwordHash)

	if !passwordValid {
		return errors.New("credentials invalid")
	}

	return nil
}
