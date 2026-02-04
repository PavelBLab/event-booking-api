package models

import (
	"errors"

	"github.com/PavelBLab/event-booking-api/configurations/postgres"
	"github.com/PavelBLab/event-booking-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {

	hashedPassword, err := utils.HashPasswordConverter(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`

	// Use QueryRow instead of Exec to get the returned ID
	err = postgres.DB.QueryRow(query, u.Email, hashedPassword).Scan(&u.ID)

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = $1`
	row := postgres.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.ComparePasswords(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("invalid password")
	}

	return nil
}
