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

func GeneratePitcher() (Pitcher, error) {
	rando := rand.New(rand.NewSource(time.Now().UnixNano())) // Instantiate the random number generator with a seed

	pitcher := Pitcher{}

	name, err := GetName()
	if err != nil {
		return Pitcher{}, err
	}

	pitcher.Name = name
	pitcher.InningsPitched = rando.Intn(100)
	pitcher.EarnedRuns = rando.Intn(100)
	pitcher.Strikeouts = rando.Intn(100)
	pitcher.Wins = rando.Intn(100)
	pitcher.EarnedRunAverage = rando.Intn(100)

	return pitcher, nil
}
