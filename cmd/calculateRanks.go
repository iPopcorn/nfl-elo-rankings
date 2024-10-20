package cmd

import (
	"errors"
	"fmt"
	"sort"

	"github.com/iPopcorn/nfl-elo-rankings/data"
	"github.com/spf13/cobra"
)

var calculateRanksCmd = &cobra.Command{
	Use:   "rank",
	Short: "Update the ranks of the teams",
	Long:  `Updates the ranks of the teams then lists them`,
	RunE:  calculateRanksHandler,
}

func calculateRanksHandler(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.New("received arguments but did not expect any")
	}

	// Get data
	repo := data.RepositoryFactory("")

	state, err := repo.GetData()

	if err != nil {
		fmt.Println("Failed to get data")
		return err
	}

	sort.Slice(state.Teams, func(i, j int) bool {
		return state.Teams[i].Rating > state.Teams[j].Rating
	})

	for i := range state.Teams {
		rank := i + 1
		state.Teams[i].Rank = rank
		state.Teams[i].SetRecord()
		state.Teams[i].Print()
	}

	err = repo.Save(*state)

	if err != nil {
		fmt.Printf("Failed to save data after updating ranks")

		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(calculateRanksCmd)
}
