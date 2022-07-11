package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

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

// type Player struct {
// 	ID          int    `json:"id"`
// 	FirstName   string `json:"firstName"`
// 	LastName    string `json:"lastName"`
// 	Age         uint8  `json:"age"`
// 	Country     string `json:"country"`
// 	Position    string `json:"position"`
// 	MarketValue int    `json:"marketValue"`
// 	TeamID      *int   `json:"-"`
// }

// type Team struct {
// 	ID          int      `json:"id"`
// 	Name        string   `json:"name"`
// 	Country     string   `json:"country"`
// 	Players     []Player `json:"players"`
// 	BankBalance int      `json:"bankBalance"`
// 	AccountID   int      `json:"-"`
// }

// type Account struct {
// 	ID        int    `json:"id"`
// 	Username  string `json:"email"`
// 	Password  string `json:"-"`
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Team      *Team  `json:"team"`
// }

// type User struct {
// 	AccountID int
// 	Username  string
// }

// type Transfer struct {
// 	ID              int    `json:"id"`
// 	Player          Player `json:"player"`
// 	MarketValue     int    `json:"marketValue"`
// 	AskedPrice      int    `json:"askedPrice"`
// 	TransferredFrom *Team  `json:"-"`
// 	TransferredTo   *Team  `json:"-"`
// 	Transferred     bool   `json:"-"`
// }

// const (
// 	ATTACKER   = "AT"
// 	GOALKEEPER = "GK"
// 	DEFENDER   = "DF"
// 	MIDFIELDER = "MF"
// )
