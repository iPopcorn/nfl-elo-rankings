package types

import (
	"fmt"
)

type Team struct {
	LastUpdated string  `json:"last_updated"`
	Name        string  `json:"name"`
	Tag         string  `json:"tag"`
	Rating      float64 `json:"rating"`
	Rank        int     `json:"rank"`
	Wins        int     `json:"wins"`
	Losses      int     `json:"losses"`
	Games       []Game  `json:"games"`
}

func (t *Team) SetRecord() {
	var wins, losses int

	for _, game := range t.Games {
		if game.Team1 == t.Name {
			if game.Team1Score > game.Team2Score {
				wins++
			} else {
				losses++
			}
		} else {
			if game.Team1Score > game.Team2Score {
				losses++
			} else {
				wins++
			}
		}
	}

	t.Wins = wins
	t.Losses = losses
}

func (t *Team) Print() {
	fmt.Printf("%d) %s (%d - %d) %.0f\n", t.Rank, t.Name, t.Wins, t.Losses, t.Rating)
}

func (t *Team) DisplayGames() {
	for _, game := range t.Games {
		fmt.Printf("%s %d - %d %s\n", game.Team1, game.Team1Score, game.Team2Score, game.Team2)
	}
}
