package types

type Team struct {
	LastUpdated string  `json:"last_updated"`
	Name        string  `json:"name"`
	Tag         string  `json:"tag"`
	Rating      float64 `json:"rating"`
	Rank        int     `json:"rank"`
	Games       []Game  `json:"games"`
}
