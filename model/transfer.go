package model

import (
	"errors"
	"time"
)

type Transfer struct {
	ID          uint      `json:"id" gorm:"primarykey,index"`
	PlayerID    uint      `json:"player_id" gorm:"index"`
	MarketValue int64     `json:"market_value"`
	AskedPrice  int64     `json:"asked_price"`
	Transferred bool      `json:"-" gorm:"default:false"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Player      Player    `json:"player"`
}

type TransferRequest struct {
	PlayerID   uint  `json:"player_id"`
	AskedPrice int64 `json:"asked_price"`
}

func (transferRequest *TransferRequest) Validate() error {
	if transferRequest.AskedPrice <= 0 {
		return errors.New("asked_price must be greater than 0")
	}
	if transferRequest.PlayerID <= 0 {
		return errors.New("invalid player_id")
	}
	return nil
}
