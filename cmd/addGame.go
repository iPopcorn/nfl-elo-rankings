package cmd

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/iPopcorn/nfl-elo-rankings/data"
	"github.com/iPopcorn/nfl-elo-rankings/types"
	"github.com/spf13/cobra"
)

var addGameCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the outcome of a game to update rankings",
	Long: `Add the outcome of a game in the following format:
team1 team1_score team2 team2_score`,
	RunE: addGameHandler,
}

var defaultRating = 1200.0

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
		fmt.Printf("Failed to convert arg to int\ngiven: %q\n", args[1])
		return err
	}

	teamScore2, err := strconv.Atoi(args[3])

	if err != nil {
		fmt.Printf("Failed to convert arg to int\ngiven: %q\n", args[3])
		return err
	}

	// Get data
	fmt.Println("Getting data...")
	repo := data.RepositoryFactory("")

	state, err := repo.GetData()

	if err != nil {
		fmt.Println("Failed to get data")
		return err
	}

	// Check if either team is in data
	var team1, team2 *types.Team
	var team1Index, team2Index int = -1, -1

	for i, team := range state.Teams {
		if teamName1 == team.Name {
			team1 = &team
			team1Index = i
		}

		if teamName2 == team.Name {
			team2 = &team
			team2Index = i
		}
	}

	// create team if missing
	if team1 == nil {
		fmt.Printf("Could not find %q in state, creating...\n", teamName1)
		var games []types.Game

		team1 = &types.Team{
			Name:   teamName1,
			Rating: defaultRating,
			Games:  games,
		}
	}

	if team2 == nil {
		fmt.Printf("Could not find %q in state, creating...\n", teamName2)
		var games []types.Game

		team2 = &types.Team{
			Name:   teamName2,
			Rating: defaultRating,
			Games:  games,
		}
	}

	// create game and calculate rating
	fmt.Println("Calculating rating based on game")
	k := 30.0
	win := 1.0
	lose := 0.0
	var newRating1, newRating2, team1RatingChange, team2RatingChange float64

	fmt.Printf("team2.Rating - team1.Rating = %f\n", team2.Rating-team1.Rating)

	probability1 := (1.0 / (1.0 + math.Pow(10, (team2.Rating-team1.Rating)/400)))
	probability2 := (1.0 / (1.0 + math.Pow(10, (team1.Rating-team2.Rating)/400)))

	if teamScore1 > teamScore2 {
		fmt.Printf("%q wins against %q\n", teamName1, teamName2)

		fmt.Printf("probability1: %f\n", probability1)
		fmt.Printf("probability2: %f\n", probability2)

		fmt.Printf("win - probability1 = %f\n", (win - probability1))

		team1RatingChange = k * (win - probability1)
		team2RatingChange = k * (lose - probability2)

		fmt.Printf("team1RatingChange = %f\n", team1RatingChange)
		fmt.Printf("team2RatingChange = %f\n", team2RatingChange)

		newRating1 = team1.Rating + team1RatingChange
		newRating2 = team2.Rating + team2RatingChange
	} else {
		fmt.Printf("%q wins against %q\n", teamName2, teamName1)
		newRating1 = team1.Rating + (k * (lose - probability1))
		newRating2 = team2.Rating + (k * (win - probability2))
	}

	fmt.Printf("New rating for %q is %f\n", teamName1, newRating1)
	fmt.Printf("New rating for %q is %f\n", teamName2, newRating2)

	// team1RatingChange := newRating1 - team1.Rating
	// team2RatingChange := newRating2 - team2.Rating

	game := &types.Game{
		Team1:             team1.Name,
		Team2:             team2.Name,
		Team1Score:        teamScore1,
		Team2Score:        teamScore2,
		Team1RatingChange: team1RatingChange,
		Team2RatingChange: team2RatingChange,
	}

	// add game to both teams
	team1.Games = append(team1.Games, *game)
	team2.Games = append(team2.Games, *game)
	team1.Rating = newRating1
	team2.Rating = newRating2

	// update state
	if team1Index > -1 {
		state.Teams[team1Index] = *team1
	} else {
		state.Teams = append(state.Teams, *team1)
	}

	if team2Index > -1 {
		state.Teams[team2Index] = *team2
	} else {
		state.Teams = append(state.Teams, *team2)
	}

	err = repo.Save(*state)

	if err != nil {
		fmt.Println("Failed to save state after updating")

		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(addGameCmd)
}
