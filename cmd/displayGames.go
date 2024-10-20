package cmd

import (
	"fmt"

	"github.com/iPopcorn/nfl-elo-rankings/data"
	"github.com/iPopcorn/nfl-elo-rankings/types"
	"github.com/spf13/cobra"
)

var displayGames = &cobra.Command{
	Use:   "games",
	Short: "display all games a team played",
	Long: `display all games that a team has played.
Takes 1 arg, the name of the team whose games you want to see.`,
	RunE: displayGamesHandler,
}

func displayGamesHandler(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("received %d arguments but expected 1", len(args))
	}

	teamName := args[0]

	// Get data
	repo := data.RepositoryFactory("")

	state, err := repo.GetData()

	if err != nil {
		fmt.Println("Failed to get data")
		return err
	}

	var foundTeam types.Team

	var teamFound bool

	for _, team := range state.Teams {
		if team.Name == teamName {
			teamFound = true
			foundTeam = team
			break
		}
	}

	if !teamFound {
		fmt.Printf("could not find team: %q\n", teamName)
	}

	foundTeam.DisplayGames()

	return nil
}

func init() {
	rootCmd.AddCommand(displayGames)
}
