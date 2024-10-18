package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var addGameCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the outcome of a game to update rankings",
	Long: `Add the outcome of a game in the following format:
team1 team1_score team2 team2_score`,
	RunE: addGameHandler,
}

func addGameHandler(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Expected an argument but did not receive one")
	}

	if len(args) != 4 {
		return fmt.Errorf("Expected 4 arg, received %d args", len(args))
	}

	fmt.Println("TODO: Implement")
	return nil
}

func init() {
	rootCmd.AddCommand(addGameCmd)
}
