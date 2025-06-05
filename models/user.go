package models

import (
	"errors"

	"github.com/AhmedOthman94/REST-API-Gin/db"
	"github.com/AhmedOthman94/REST-API-Gin/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password) VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("could not prepare statement")
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := `
	SELECT email, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Email, &retrievedPassword)
	if err != nil {
		return errors.New("credintials are not valid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credintials are not valid")
	}

	return nil
}
