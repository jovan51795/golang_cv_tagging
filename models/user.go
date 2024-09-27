package models

import (
	"errors"

	"77gsi_mynt.com/cv_tagging/db"
	"77gsi_mynt.com/cv_tagging/util"
)

type User struct {
	Id        int64
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func (u *User) Signup() error {
	query := `INSERT INTO user_tbl(first_name, last_name, email, password) VALUES($1, $2, $3, $4)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(u.Password)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(&u.Firstname, &u.Lastname, &u.Email, hashedPassword)

	if err != nil {
		return err
	}
	return err
}

func (u *User) Login() error {
	query := `SELECT id, password FROM user_tbl WHERE email = $1`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return err
	}

	isUserValid := util.ValidatePassword(u.Password, retrievedPassword)

	if !isUserValid {
		return errors.New("invalid credentials")
	}

	return nil
}
