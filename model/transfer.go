package model

import "time"

type Transfer struct {
	ID              uint      `json:"id" gorm:"primarykey"`
	PlayerID        uint      `json:"player_id" gorm:"unique"`
	MarketValue     string    `json:"market_value" gorm:"unique"`
	AskedPrice      string    `json:"asked_price" gorm:"unique"`
	TransferredFrom Team      `json:"-"`
	TransferredTo   Team      `json:"-"`
	Transferred     bool      `json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
	Player          Player
}

const (
	UniqueConstraintPlayerID = "transfers_player_id_key"
)
