package model

import (
	"math/rand"
	"time"

	constants "github.com/nvzard/soccer-manager/helpers"
)

type Player struct {
	ID          uint      `json:"id" gorm:"primarykey,index"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Country     string    `json:"country"`
	Age         uint8     `json:"age"`
	Position    string    `json:"position"`
	MarketValue int64     `json:"market_value" gorm:"default:1000000"`
	TeamID      uint      `json:"team_id" gorm:"index"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type PlayerPatch struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
}

func (player *Player) GeneratePlayer(position string) {
	player.FirstName = constants.GetRandomFirstName()
	player.LastName = constants.GetRandomLastName()
	player.Country = constants.GetRandomCountry()
	player.Age = constants.GetRandomAge()
	player.Position = position
}

func (player *Player) Transfer(teamID uint) {
	player.TeamID = teamID
	player.increaseMarketValue()
}

// Randomly increase player market value b/w 10 and 100 percent
func (player *Player) increaseMarketValue() {
	rand.Seed(time.Now().UnixNano())
	randomPercent := (10 + rand.Intn(90))
	amountToIncrease := ((float64(player.MarketValue) * float64(randomPercent)) / float64(100))
	increasedAmount := player.MarketValue + int64(amountToIncrease)
	player.MarketValue = increasedAmount
}

func GeneratePlayers(goalkeeperCount, defendersCount, midfieldersCount, attackersCount int) []Player {
	totalPlayers := goalkeeperCount + defendersCount + midfieldersCount + attackersCount
	players := make([]Player, totalPlayers)
	i := 0

	for goalkeeperCount > 0 {
		players[i].GeneratePlayer(constants.GOALKEEPER)
		goalkeeperCount--
		i++
	}

	for defendersCount > 0 {
		players[i].GeneratePlayer(constants.DEFENDER)
		defendersCount--
		i++
	}

	for midfieldersCount > 0 {
		players[i].GeneratePlayer(constants.MIDFIELDER)
		midfieldersCount--
		i++
	}

	for attackersCount > 0 {
		players[i].GeneratePlayer(constants.ATTACKER)
		attackersCount--
		i++
	}

	return players
}
