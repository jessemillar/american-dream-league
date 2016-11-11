package helpers

import (
	"math/rand"
	"time"
)

type Hitter struct {
	Name           string  `json:"name"`
	Hits           int     `json:"hits"`
	Runs           int     `json:"runs"`
	Homeruns       int     `json:"homeruns"`
	RunsBattedIn   int     `json:"runsBattedIn"`
	BattingAverage float64 `json:"battingAverage"`
}

type Pitcher struct {
	Name             string `json:"name"`
	InningsPitched   int    `json:"inningsPitched"`
	EarnedRuns       int    `json:"earnedRuns"`
	Strikeouts       int    `json:"strikeouts"`
	Wins             int    `json:"wins"`
	EarnedRunAverage int    `json:"earnedRunAverage"`
}

func GenerateHitter() (Hitter, error) {
	rando := rand.New(rand.NewSource(time.Now().UnixNano())) // Instantiate the random number generator with a seed

	hitter := Hitter{}

	name, err := GetName()
	if err != nil {
		return Hitter{}, err
	}

	hitter.Name = name
	hitter.Hits = rando.Intn(100)
	hitter.Runs = rando.Intn(100)
	hitter.RunsBattedIn = rando.Intn(100)
	hitter.Homeruns = rando.Intn(100)
	hitter.BattingAverage = rando.Float64() * 0.4

	return hitter, nil
}
