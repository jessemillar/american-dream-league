package helpers

import "math/rand"

type Hitter struct {
	Hits           int     `json:"hits"`
	Runs           int     `json:"runs"`
	Homeruns       int     `json:"homeruns"`
	RunsBattedIn   int     `json:"runsBattedIn"`
	BattingAverage float64 `json:"battingAverage"`
}

type Pitcher struct {
	InningsPitched   int `json:"InningsPitched"`
	EarnedRuns       int `json:"earnedRuns"`
	Strikeouts       int `json:"strikeouts"`
	Wins             int `json:"wins"`
	EarnedRunAverage int `json:"earnedRunAverage"`
}

func GenerateHitter() (Hitter, error) {
	hitter := Hitter{}

	hitter.BattingAverage = rand.Float64() * 0.3

	return hitter, nil
}
