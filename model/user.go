package model

import (
	"errors"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	TeamID    uint   `json:"team_id"`
	Team      Team
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserAuth struct {
	ID     uint
	TeamID uint
	Email  string
}

const (
	UniqueConstraintEmail = "users_email_key"
)

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (user *User) Validate() error {
	if user.FirstName == "" {
		return errors.New("first_name is required")
	}
	if user.LastName == "" {
		return errors.New("last_name is required")
	}
	if !validEmail(user.Email) {
		return errors.New("email is invalid")
	}
	if user.Password == "" || len(user.Password) < 8 {
		return errors.New("password length must be at least 8 characters")
	}
	return nil
}

func (user *User) CreateTeam() {
	team := Team{}
	team.GenerateTeam()
	user.Team = team
}
