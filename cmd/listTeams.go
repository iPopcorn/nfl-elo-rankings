package cmd

import (
	"errors"
	"fmt"
	"sort"

	"github.com/iPopcorn/nfl-elo-rankings/data"
	"github.com/spf13/cobra"
)

var listTeamsCmd = &cobra.Command{
	Use:   "list",
	Short: "List the teams",
	Long:  `List the teams by rating in descending order`,
	RunE:  listTeamHandler,
}

func listTeamHandler(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.New("Received arguments but did not expect any")
	}

	// Get data
	fmt.Println("Getting data...")
	repo := data.RepositoryFactory("")

	state, err := repo.GetData()

	if err != nil {
		fmt.Println("Failed to get data")
		return err
	}

	sort.Slice(state.Teams, func(i, j int) bool {
		return state.Teams[i].Rating > state.Teams[j].Rating
	})

	for i, team := range state.Teams {
		fmt.Printf("%d) %q %f\n", i+1, team.Name, team.Rating)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listTeamsCmd)
}
