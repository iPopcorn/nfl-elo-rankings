package cmd

import (
	"fmt"

	"github.com/iPopcorn/nfl-elo-rankings/data"
	"github.com/iPopcorn/nfl-elo-rankings/types"
	"github.com/spf13/cobra"
)

var findTeam = &cobra.Command{
	Use:   "find",
	Short: "find teams by their names",
	Long: `find multiple teams by name.
Case sensitive
Takes many args, the name of the teams to find`,
	RunE: findTeamHandler,
}

func findTeamHandler(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("received no arguments but expected at least 1")
	}

	// Get data
	repo := data.RepositoryFactory("")

	state, err := repo.GetData()

	if err != nil {
		fmt.Println("Failed to get data")
		return err
	}

	foundTeams := []types.Team{}

	for _, teamName := range args {
		var teamFound bool

		for _, team := range state.Teams {
			if team.Name == teamName {
				teamFound = true
				foundTeams = append(foundTeams, team)
				break
			}
		}

		if !teamFound {
			fmt.Printf("could not find team: %q\n", teamName)
		}
	}

	for _, team := range foundTeams {
		fmt.Printf("%d) %q: %.0f\n", team.Rank, team.Name, team.Rating)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(findTeam)
}
