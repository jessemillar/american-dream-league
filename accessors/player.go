package accessors

type Hitter struct {
	ID             int
	Name           int
	OnBasePercent  int
	BattingAverage int
	Strikeouts     int
	SlugPercent    int
	PlayerID       int
	Hit            int
	AtBats         int
}

// GetHitters returns all hitters from the database
func (accessorGroup *AccessorGroup) GetHitters() (Hitter, error) {
	hitter := &Hitter{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Hitters").Scan(&hitter.ID, &hitter.Name)

	if err != nil {
		return Hitter{}, err
	}

	return *hitter, nil
}

// GetPitchers returns all hitters from the database
func (accessorGroup *AccessorGroup) GetPitchers() (Hitter, error) {
	hitter := &Hitter{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Pitchers").Scan(&hitter.ID, &hitter.Name)

	if err != nil {
		return Hitter{}, err
	}

	return *hitter, nil
}
