package model

import (
	constants "github.com/nvzard/soccer-manager/helpers"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name          string `json:"name"`
	Country       string `json:"country"`
	AvailableCash int    `json:"availableCash" gorm:"default:5000000"`
	UserID        uint   `json:"user_id"`
	Players []Player `json:"players"`
}

func (team *Team) GenerateTeam() {
	team.Name = constants.GetRandomTeamName()
	team.Country = constants.GetRandomCountry()
	// Goalkeepers: 3, Defenders: 6, Midfielders: 6, Attackers: 5
	team.Players = GeneratePlayers(3, 6, 6, 5)
}
