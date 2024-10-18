package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"math"

	"github.com/spf13/cobra"
	"github.com/iPopcorn/nfl-elo-rankings/data"
	"github.com/iPopcorn/nfl-elo-rankings/types"
)

var addGameCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the outcome of a game to update rankings",
	Long: `Add the outcome of a game in the following format:
team1 team1_score team2 team2_score`,
	RunE: addGameHandler,
}

var defaultRating = 1200

func addGameHandler(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Expected an argument but did not receive one")
	}

	if len(args) != 4 {
		return fmt.Errorf("Expected 4 arg, received %d args", len(args))
	}

	// parse args
	teamName1 := args[0]
	teamName2 := args[2]

	teamScore1, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf("Failed to convert arg to int\ngiven: %q", args[1])
		return err
	}

	teamScore2, err := strconv.Atoi(args[3])

	if err != nil {
		fmt.Printf("Failed to convert arg to int\ngiven: %q", args[3])
		return err
	}
	
	// Get data
	repo := data.RepositoryFactory("")

	state, err := repo.GetData()

	if err != nil {
		fmt.Println("Failed to get data")
		return err
	}

	// Check if either team is in data
	var team1, team2 *types.Team

	for _, team := range state.Teams {
		if teamName1 == team.Name {
			team1 = &team
		}

		if teamName2 == team.Name {
			team2 = &team
		}
	}

	// create team if missing
	if team1 == nil {
		var games []types.Game

		team1 = &types.Team{
			Name: teamName1,
			Rating: defaultRating,
			Games: games,
		}
	}

	if team2 == nil {
		var games []types.Game

		team2 = &types.Team{
			Name: teamName2,
			Rating: defaultRating,
			Games: games,
		}
	}

	// create game and calculate rating
	k := 30
	probability1 := (1.0/(1.0 + math.Pow10((team1.Rating - team2.Rating)) / 400))
	probability2 := (1.0/(1.0 + math.Pow10((team2.Rating - team1.Rating)) / 400))

	team1RatingChange := 

	game := &types.Game{
		Team1: team1.Name,
		Team2: team2.Name,
		Team1Score: teamScore1,
		Team2Score: teamScore2,
	}
	// add game to both teams
	return nil
}

func init() {
	rootCmd.AddCommand(addGameCmd)
}
