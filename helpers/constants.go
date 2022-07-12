package helpers

import (
	"math/rand"
	"time"
)

var FirstNames = []string{"Kyle", "Keenan", "Douglas", "Grant", "Jonathon", "Braydon", "Rolando", "Liam", "Ryland", "Nathaniel", "Bentley", "Layne"}
var LastNames = []string{"Bean", "Lynch", "Shannon", "Mejia", "Jarvis", "Zhang", "Jimenez", "Klein", "Holden", "Cohen", "Grimes", "Sellers"}
var Countries = []string{"Greenland", "Venezuela", "Germany", "Norway", "Brazil", "Spain", "Qatar", "Antarctica", "France", "Belarus", "Georgia", "USA"}
var TeamNames = []string{"Careless Wolverines", "Fearless Manticores", "Hungry Blazers", "Mad Gibbons", "Hungry Leopards", "Regal Stingers", "Thunder", "Extraordinary", "Iron Squids", "Mallards", "Crickets", "Sharks"}

const (
	ATTACKER   = "ATT"
	GOALKEEPER = "GK"
	DEFENDER   = "DEF"
	MIDFIELDER = "MID"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomFirstName() string {
	return FirstNames[rand.Intn(len(FirstNames))]
}

func GetRandomLastName() string {
	return LastNames[rand.Intn(len(LastNames))]
}

func GetRandomCountry() string {
	return Countries[rand.Intn(len(Countries))]
}

func GetRandomTeamName() string {
	return TeamNames[rand.Intn(len(TeamNames))]
}

func GetRandomAge() uint8 {
	minAge := 18
	maxAge := 40
	return uint8(minAge + rand.Intn(maxAge-minAge))
}
