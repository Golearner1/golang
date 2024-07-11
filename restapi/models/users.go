package models

import (
	"errors"
	"fmt"

	"example.com/first-app/database"
	"example.com/first-app/secure"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(Email,Password)VALUES(?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer stmt.Close()
	encrpytpass, err := secure.Lockpassword(u.Password)
	if err != nil {
		fmt.Print(err)
		return err
	}
	result, err := stmt.Exec(u.Email, encrpytpass)
	if err != nil {
		fmt.Print(err)
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}
func (u *User) Validateuser() error {
	query := `SELECT id,Password FROM users WHERE Email=? `
	rows := database.DB.QueryRow(query, u.Email)
	var passwordhttp string
	err := rows.Scan(&u.ID, &passwordhttp)
	if err != nil {
		return err
	}
	passwordinvalid := secure.Checkpassword(u.Password, passwordhttp)

	if !passwordinvalid {
		return errors.New("invalid password")
	}
	passwordvalid := secure.Checkpassword(u.Password, passwordhttp)

	if !passwordvalid {
		return errors.New("invalid details")
	}
	return nil
}
