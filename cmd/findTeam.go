package cmd

import (
	"fmt"

	"github.com/iPopcorn/nfl-elo-rankings/data"
	"github.com/iPopcorn/nfl-elo-rankings/types"
	"github.com/spf13/cobra"
)

var findTeam = &cobra.Command{
	Use:   "find",
	Short: "find a team by name",
	Long: `find a team by name.
Case sensitive
Takes 1 arg, the name of the team to find`,
	RunE: findTeamHandler,
}

func findTeamHandler(cmd *cobra.Command, args []string) error {
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

	var teamFound bool
	var foundTeam types.Team

	for _, team := range state.Teams {
		if team.Name == teamName {
			teamFound = true
			foundTeam = team
			break
		}
	}

	if !teamFound {
		fmt.Printf("could not find team: %q\n", teamName)

		return nil
	}

	fmt.Printf("%q: %f\n", foundTeam.Name, foundTeam.Rating)
	return nil
}

func init() {
	rootCmd.AddCommand(findTeam)
}
