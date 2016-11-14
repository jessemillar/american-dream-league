package accessors

type Hitter struct {
	HitterID       int
	OnBasePercent  int
	BattingAverage int
	Strikeouts     int
	SlugPercent    int
	PlayerID       int
	Hit            int
	AtBats         int
}

// GetLeagueByName returns a league from the database by name
func (accessorGroup *AccessorGroup) GetLeagueByName(name string) (League, error) {
	league := &League{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Hitters", name).Scan(&league.ID, &league.Name)

	if err != nil {
		return League{}, err
	}

	return *league, nil
}
