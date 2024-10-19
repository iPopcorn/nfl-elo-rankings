package types

type Game struct {
	Team1             string  `json:"team_1"`
	Team2             string  `json:"team_2"`
	Team1Score        int     `json:"team_1_score"`
	Team2Score        int     `json:"team_2_score"`
	Team1RatingChange float64 `json:"team_1_rating_change"`
	Team2RatingChange float64 `json:"team_2_rating_change"`
}
