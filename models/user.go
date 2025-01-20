package models

import (
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
