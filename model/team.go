package model

import (
	"time"

	constants "github.com/nvzard/soccer-manager/helpers"
)

type Team struct {
	ID            uint      `json:"id" gorm:"primarykey"`
	Name          string    `json:"name"`
	Country       string    `json:"country"`
	AvailableCash int       `json:"available_cash" gorm:"default:5000000"`
	UserID        uint      `json:"user_id"`
	Players       []Player  `json:"players"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type TeamPatch struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (team *Team) GenerateTeam() {
	team.Name = constants.GetRandomTeamName()
	team.Country = constants.GetRandomCountry()
	// Goalkeepers: 3, Defenders: 6, Midfielders: 6, Attackers: 5
	team.Players = GeneratePlayers(3, 6, 6, 5)
}

func (team *Team) CalculateTeamValue() int {
	var teamValue int
	for _, player := range team.Players {
		teamValue += player.MarketValue
	}
	return teamValue
}
