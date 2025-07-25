package models

import (
	"errors"

	"github.com/jhunderdog/go_api/db"
	"github.com/jhunderdog/go_api/utils"
)

type User struct {
	ID int64 
	Email string `biding:"required"`
	Password string `biding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (? , ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
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

func (u *User) ValidateCrentials() error{
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID,&retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}