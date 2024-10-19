package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/iPopcorn/nfl-elo-rankings/types"
	"github.com/iPopcorn/nfl-elo-rankings/util"
)

type Repository struct {
	filename string
}

func RepositoryFactory(filename string) *Repository {
	defaultName := "data.json"

	if filename != "" {
		return &Repository{
			filename: filename,
		}
	}

	return &Repository{
		filename: defaultName,
	}
}

func (r *Repository) GetData() (*types.State, error) {
	location := "Repository.GetData()\n"
	filepath, err := util.GetPathToFile("", r.filename)

	if err != nil {
		fmt.Printf(location+"Failed to get path to file\n%v\n", err)
		return nil, err
	}

	data, err := os.ReadFile(filepath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No save data found, creating...")
			newState := r.InitState()

			return newState, nil
		} else {
			fmt.Printf(location+"Failed to read file\n%v\n", err)
			return nil, err
		}
	}

	var state types.State

	err = json.Unmarshal(data, &state)

	if err != nil {
		fmt.Printf(location+"Failed to de-serialize state.\nGiven: %s\n%v\n", string(data), err)

		return nil, err
	}

	return &state, nil
}

func (r *Repository) Save(newState types.State) error {
	location := "Repository.Save()\n"
	filepath, err := util.GetPathToFile("", r.filename)

	if err != nil {
		fmt.Printf(location+"Failed to get path to file\n%v\n", err)
		return err
	}
	fmt.Printf("filepath: %q\n", filepath)

	_, err = os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(filepath)

		if err != nil {
			fmt.Printf(location+"Failed to create file\n%v\n", err)
			return err
		}

		f.Close()
	}

	fmt.Printf(location + "marshalling data...\n")
	data, err := json.Marshal(newState)

	if err != nil {
		fmt.Printf(location + "Failed to marshal state into []byte")
		return err
	}

	return os.WriteFile(filepath, data, 0666)
}

func (r *Repository) InitState() *types.State {
	initialState := &types.State{
		Teams: []types.Team{},
	}

	return initialState
}
