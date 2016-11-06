package accessors

type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetLeagueByName returns a league from the database by name
func (accessorGroup *AccessorGroup) GetLeagueByName(name string) (League, error) {
	league := &League{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM League WHERE name=?", name).Scan(&league.ID, &league.Name)

	if err != nil {
		return League{}, err
	}

	return *league, nil
}

// GetLeagueById returns a league from the database by leagueID
func (accessorGroup *AccessorGroup) GetLeagueById(name string) (League, error) {
	league := &League{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM League WHERE leagueID=?", name).Scan(&league.ID, &league.Name)

	if err != nil {
		return League{}, err
	}

	return *league, nil
}

// GetLeagueID returns a leagueID from the database from a name
func (accessorGroup *AccessorGroup) GetLeagueID(name string) (int, error) {
	league, err := accessorGroup.GetLeagueByName(name)
	if err != nil {
		return 0, err
	}

	return league.ID, nil
}

// MakeLeague adds a league to the database
func (accessorGroup *AccessorGroup) MakeLeague(name string) (League, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO League (name) VALUES (?)", name)
	if err != nil {
		return League{}, err
	}

	league, err := accessorGroup.GetLeagueByName(name)
	if err != nil {
		return League{}, err
	}

	return league, nil
}
